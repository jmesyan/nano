package utils

import (
	"errors"
	"fmt"
	"github.com/jmesyan/nano/nodes"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	ListenAddr string
	logger     = log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
)

func getMacAddr() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		panic(err.Error())
	}
	mac, macerr := "", errors.New("无法获取到正确的MAC地址")
	for i := 0; i < len(netInterfaces); i++ {

		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				ipnet, ok := address.(*net.IPNet)

				if ok && ipnet.IP.IsGlobalUnicast() {

					mac = netInterfaces[i].HardwareAddr.String()
					return mac, nil
				}
			}
		}
	}
	return mac, macerr
}

func GenerateNodeId(ntype nodes.NodeType, gsid string) string {
	nkind := nodes.NodeTypesToKind[ntype]
	if ntype == nodes.NodeGameServer {
		return fmt.Sprintf("%s_%s", nkind, gsid)
	}
	address := GenerateLocalAddr()
	return fmt.Sprintf("%s_%s", nkind, address)
}

func GenerateLocalAddr() string {
	macaddr, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	if len(ListenAddr) == 0 {
		logger.Fatal("no listen addr")
		return ""
	}
	ls, lp := strings.Split(ListenAddr, ":"), "0"
	if len(ls) == 2 {
		lp = ls[1]
	} else {
		logger.Fatal("addr type is wrong")
		return ""
	}
	return fmt.Sprintf("%s:%s", macaddr, lp)
}

func GenerateTopic(s ...string) string {
	return strings.Join(s, ".")
}

func StringToInt(valstr string) int {
	val, err := strconv.Atoi(valstr)
	if err != nil {
		val = 0
	}
	return val
}

func StringToInt64(valstr string) int64 {
	val := StringToInt(valstr)
	return int64(val)
}

func Int64ToString(valint int64) string {
	return strconv.FormatInt(valint, 10)
}

func Int64Toint(valint int64) int {
	valstr := Int64ToString(valint)
	return StringToInt(valstr)
}

func IntToInt64(valint int) int64 {
	s := strconv.Itoa(valint)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		println(err.Error())
		return 0
	}
	return i
}
func IntToString(intval int) string {
	return strconv.Itoa(intval)
}

func Time() int {
	now := time.Now().Unix()
	return Int64Toint(now)
}
