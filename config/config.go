package config

import (
	"fmt"
	"net"
	"strings"

	"gorm.io/gorm"
)

var SavePath string = "http://" + GetHostIp() + ":8080/static/"

var MysqlUsername string = "root"
var MysqlPassword string = "12345" //123456
var MysqlUrl string = "tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"

var Db *gorm.DB

func GetHostIp() string {
	// 打个洞
	conn, err := net.Dial("udp", "baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
