// Copyright (c) nano Author. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/jmesyan/nano/nodes"
	"io"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	ListenAddr string
	logger     = log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
)

func SerializeOrRaw(v interface{}) ([]byte, error) {
	if data, ok := v.([]byte); ok {
		return data, nil
	}
	data, err := Serializer.Marshal(v)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GobEncode(args ...interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte(nil))
	if err := gob.NewEncoder(buf).Encode(args); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GobDecode(reply interface{}, data []byte) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(reply)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func Stack() string {
	buf := make([]byte, 10000)
	n := runtime.Stack(buf, false)
	buf = buf[:n]

	s := string(buf)

	// skip nano frames lines
	const skip = 7
	count := 0
	index := strings.IndexFunc(s, func(c rune) bool {
		if c != '\n' {
			return false
		}
		count++
		return count == skip
	})
	return s[index+1:]
}
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

func Md5(sign string) string {
	h := md5.New()
	io.WriteString(h, sign)
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum
}
