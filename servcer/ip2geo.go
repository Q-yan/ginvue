package main

import (
	"ginvue/moudle"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/oschwald/geoip2-golang"
)

type jingweidus struct {
	Latitude  float64 `json:"latitude"`  //经度
	Longitude float64 `json:"longitude"` //纬度

}

func generateData() moudle.Ites {
	ite1 := moudle.Ites{
		Date: time.Now().Format("15:04:05"),
		Bar:  rand.Intn(900),
		Line: rand.Intn(30),
	}
	//fmt.Println(ite1)
	//c.JSON(http.StatusOK, ite1)
	//data := make(map[string]interface{})
	//data["timestamp"] = time.Now().Unix()
	//data["value"] = generateRandomValue()
	return ite1
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	// 向前端定时发送数据
	ticker := time.NewTicker(2 * time.Second)
	//fmt.Println(ticker)
	defer ticker.Stop()

	for range ticker.C {
		// 生成数据
		data := generateData()

		// 发送数据给前端
		err := conn.WriteJSON(data)
		if err != nil {
			log.Println("Failed to send data to WebSocket client:", err)
			break
		}
	}
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
