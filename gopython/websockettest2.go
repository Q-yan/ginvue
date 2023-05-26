package main

import (
	"encoding/json"
	"fmt"
	"ginvue/moudle"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有的源
		},
	}
	clients = make(map[*websocket.Conn]bool)
	Num     = 1
)

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("WebSocket server started")
	log.Fatal(http.ListenAndServe(":8765", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 将 HTTP 请求升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()
	clients[conn] = true
	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message from client:", err)
			break
		}

		// 处理接收到的消息
		//log.Println("Received message:", string(msg))

		// 可以在这里编写逻辑，对接收到的消息进行处理
		err = conn.WriteJSON(string(msg))
		if err != nil {
			log.Println("Failed to send data to WebSocket client:", err)
			break
		}

		// 回复客户端确认消息
		if err := conn.WriteMessage(websocket.TextMessage, []byte("Message received")); err != nil {
			log.Println("Failed to send message to client:", err)
			break
		}
		sendMessageToAll(string(msg))
	}

}

func sendMessageToAll(message string) {

	var data moudle.Ites
	err := json.Unmarshal([]byte(message), &data)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return
	}
	//
	//num, err := strconv.ParseFloat(message, 64)
	//if err != nil {
	//	fmt.Println("转换失败:", err)
	//	return
	//}
	//ite1 := moudle.Ites{
	//	Date: time.Now().Format("15:04:05"),
	//	Bar:  int(num) / 1000000,
	//	Line: rand.Intn(30),
	//}
	data.Date = time.Now().Format("15:04:05")
	Num = Num + 1
	data.Num = Num
	for conn := range clients {
		err := conn.WriteJSON(data)
		if err != nil {
			log.Println("Failed to send message to client:", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}
