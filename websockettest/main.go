package main

import (
	"ginvue/moudle"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("WebSocket server started on :8765")
	log.Fatal(http.ListenAndServe(":8765", nil))
}

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

// 生成数据
func generateData() moudle.Ites {
	ite1 := moudle.Ites{
		Date: time.Now().Format("15:04:05"),
		Bar:  rand.Intn(900),
		Line: rand.Intn(30),
	}
	return ite1
}

// 生成随机数
func generateRandomValue() int {

	return 1
}
