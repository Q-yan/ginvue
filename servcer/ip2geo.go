package main

import (
	"net"

	"github.com/oschwald/geoip2-golang"
)

type jingweidus struct {
	Latitude  float64 `json:"latitude"`  //经度
	Longitude float64 `json:"longitude"` //纬度

}

func ip2eo(ipaddr string) (jingwei []float64) {
	db, err := geoip2.Open("news/GeoLite2-City.mmdb") // 替换为您的 IP 地理位置数据库文件路径
	if err != nil {
		return nil
	}
	defer db.Close()

	ip := net.ParseIP(ipaddr) // 替换为您要转换的 IP 地址

	record, err := db.City(ip)
	if err != nil {
		return nil
	}

	//jingweid = &jingweidus{
	//	Latitude:  0,
	//	Longitude: 0,
	//}

	jingwei = append(jingwei, record.Location.Longitude, record.Location.Latitude)
	//jingweid.Latitude =
	//jingweid.Longitude =
	//jingweidus{
	//	Latitude:  //经度
	//	Longitude: // 纬度
	//}

	return
}
