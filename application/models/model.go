package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"hewolf/config"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	DB     *dbm
	dbr    *xorm.Engine
	logger *logrus.Entry
)

type dbm struct {
	db *xorm.Engine
}

func init() {
	logger = filesystemLogger("stargames", "app.Models")
	config := config.Connections[config.DefaultConnect]
	args := fmt.Sprintf("charset=%s", config.Charset)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		config.Username,
		config.Password, 
		config.Host,
		config.Port,
		config.Database,
		args,
	)
	if dbx, err := xorm.NewEngine(config.Driver, dsn); err != nil {
		logger.Fatal(err)
	} else {
		DB = &dbm{
			db: dbx,
		}
		dbr = DB.db
	}
	dbr.SetMaxIdleConns(20)
	dbr.SetMaxOpenConns(15)
	dbr.ShowSQL(true)
	// 定时ping数据库, 保持连接池连接
	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		for {
			select {
			case <-ticker.C:
				dbr.Ping()
			}
		}
	}()
}

func filesystemLogger(logFileName, logSource string) *logrus.Entry {
	baseLogPaht := config.STORAGEPATH + path.Join("/logs", logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+"_%Y%m%d.log",
		rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})
	logger := logrus.New()
	logger.AddHook(lfHook)
	loger := logger.WithField("source", logSource)
	return loger
}

func GetDb() *xorm.Engine {
	return DB.db
}

func (this *dbm) Select(sql string) []map[string]interface{} {
	vals, err := this.db.QueryInterface(sql)
	if err != nil {
		logger.Info(err.Error())
		return nil
	} else {
		return vals
	}
}

func (this *dbm) SelectOne(sql string) map[string]interface{} {
	vals, err := this.db.QueryInterface(sql)
	if err != nil {
		logger.Info(err.Error())
		return nil
	} else {
		if vals == nil {
			return nil
		}
		return vals[0]
	}
}

func (this *dbm) Insert(sql string) int64 {
	res, err := this.db.Exec(sql)
	if err != nil {
		logger.Info(err.Error())
		return 0
	} else {
		insertid, err := res.LastInsertId()
		if err != nil {
			logger.Info(err.Error())
			return 0
		} else {
			return insertid
		}
	}
}

func (this *dbm) Update(sql string) int64 {
	res, err := this.db.Exec(sql)
	if err != nil {
		logger.Info(err.Error())
		return 0
	} else {
		affected, err := res.RowsAffected()
		if err != nil {
			logger.Info(err.Error())
			return 0
		} else {
			return affected
		}
	}
}

func (this *dbm) Delete(sql string) bool {
	res, err := this.db.Exec(sql)
	if err != nil {
		logger.Info(err.Error())
		return false
	} else {
		affected, err := res.RowsAffected()
		if err != nil {
			logger.Info(err.Error())
			return false
		} else {
			if affected > 0 {
				return true
			}
			return false
		}
	}
}

func IfcToString(i interface{}) string {
	switch i.(type) {
	case []byte:
		return string(i.([]byte))
	case string:
		return i.(string)
	}
	return fmt.Sprintf("%v", i)
}
func StringToInt(valstr string) int {
	val, err := strconv.Atoi(valstr)
	if err != nil {
		val = 0
	}
	return val
}

func Int64ToString(valint int64) string {
	return strconv.FormatInt(valint, 10)
}

func Int64Toint(valint int64) int {
	valstr := Int64ToString(valint)
	return StringToInt(valstr)
}
func IfcToInt(i interface{}) int {
	switch i.(type) {
	case []byte:
		n := StringToInt(string(i.([]byte)))
		return n
	case int:
		return i.(int)
	case int64:
		return Int64Toint(i.(int64))
	}
	return 0
}

func IfcToInt64(i interface{}) int64 {
	switch i.(type) {
	case []byte:
		n, _ := strconv.ParseInt(string(i.([]byte)), 10, 64)
		return n
	case int:
		return int64(i.(int))
	case int64:
		return i.(int64)
	}
	return 0
}

func IfcToFloat64(i interface{}) float64 {
	switch i.(type) {
	case []byte:
		n, _ := strconv.ParseFloat(string(i.([]byte)), 64)
		return n
	case float64:
		return i.(float64)
	case float32:
		return float64(i.(float32))
	}
	return 0
}

func StrToTime(str string, tpl string) int64 {
	t, err := time.Parse(tpl, str)
	if err != nil {
		logger.Info(err.Error())
		return 0
	}
	return t.Unix()
}

func CheckEmail(email string) bool {
	ret, err := regexp.MatchString(`([\w-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([\w-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)`, email)
	if err != nil {
		logger.Info(err.Error())
		return false
	}
	return ret
}

func CheckIP(ip string) bool {
	ret, err := regexp.MatchString(`(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])`, ip)
	if err != nil {
		println(err.Error())
		return false
	}
	return ret
}

func CheckCN(nickname string) bool {
	ret, err := regexp.MatchString(`[x4e00-x9fa5]+`, nickname)
	if err != nil {
		println(err.Error())
		return false
	}
	return ret
}

type ModelValue struct {
	Intval int
	Strval string
}

func byteIsNumber(b []byte) bool {
	str := string(b)
	ret, err := regexp.MatchString(`^([0-9]|\+|\-)[0-9]*$`, str)
	if err != nil {
		logger.Info(err.Error())
		return false
	}
	return ret
}

func FaceToMV(i interface{}) ModelValue {
	val := ModelValue{Intval: 0, Strval: ""}
	switch i.(type) {
	case int:
		val.Intval = i.(int)
	case string:
		val.Strval = i.(string)
	case []byte:
		if byteIsNumber(i.([]byte)) {
			val.Intval = StringToInt(string(i.([]byte)))
		} else {
			val.Strval = string(i.([]byte))
		}
	}
	return val
}

func MapFaceToMV(m map[string]interface{}) map[string]ModelValue {
	val := make(map[string]ModelValue)
	for k, v := range m {
		val[k] = FaceToMV(v)
	}
	return val
}

func MapArrToMV(ms []map[string]interface{}) []map[string]ModelValue {
	len := len(ms)
	var vals []map[string]ModelValue
	for i := 0; i < len; i++ {
		vals = append(vals, MapFaceToMV(ms[i]))
	}
	return vals
}

func GetValType(val interface{}) string {
	switch val.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return ""
	}
}

func GetCountDate(years int, months int, days int) time.Time {
	return time.Unix(time.Now().AddDate(years, months, days).Unix(), 0)
}

//IncreUpdate 表累加更新
func IncreUpdate(table string, prikeys []string, data map[string]interface{}) bool {
	sqlf := "insert into %s(%s) value(%s) on duplicate key update %s"
	keys, values := []string{}, []string{}
	kvalues := []string{}
	prikmaps := make(map[string]bool)
	for _, v := range prikeys {
		prikmaps[v] = true
	}
	for k, v := range data {
		keys = append(keys, k)
		values = append(values, fmt.Sprintf("%v", v))
		if !prikmaps[k] {
			kvalues = append(kvalues, fmt.Sprintf("%s=%s+%v", k, k, v))
		}
	}
	sql := fmt.Sprintf(sqlf, table, strings.Join(keys, ","), strings.Join(values, ","), strings.Join(kvalues, ","))
	_, err := dbr.Exec(sql)
	if err != nil {
		logger.Error(err)
		return false
	}
	return true
}

//CoverUpdate 表覆盖更新
func CoverUpdate(table string, prikeys []string, data map[string]interface{}) bool {
	sqlf := "insert into %s(%s) value(%s) on duplicate key update %s"
	keys, values := []string{}, []string{}
	kvalues := []string{}
	prikmaps := make(map[string]bool)
	for _, v := range prikeys {
		prikmaps[v] = true
	}
	for k, v := range data {
		keys = append(keys, k)
		values = append(values, fmt.Sprintf("%v", v))
		if !prikmaps[k] {
			kvalues = append(kvalues, fmt.Sprintf("%s=%v", k, v))
		}
	}
	sql := fmt.Sprintf(sqlf, table, strings.Join(keys, ","), strings.Join(values, ","), strings.Join(kvalues, ","))
	_, err := dbr.Exec(sql)
	if err != nil {
		logger.Error(err)
		return false
	}
	return true
}
