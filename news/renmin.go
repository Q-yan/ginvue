package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func main() {
	db, err := geoip2.Open("news/GeoLite2-City.mmdb") // 替换为您的 IP 地理位置数据库文件路径
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP("118.144.72.252") // 替换为您要转换的 IP 地址

	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	province := record.Subdivisions[0].Names["zh-CN"] // 根据具体的语言选择对应的字段
	fmt.Println("省份:", province)
}
