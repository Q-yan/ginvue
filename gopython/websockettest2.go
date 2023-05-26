package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有的源
		},
	}
	clients = make(map[*websocket.Conn]bool)
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
		log.Println("Received message:", string(msg))

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
	for conn := range clients {
		err := conn.WriteJSON(message)
		if err != nil {
			log.Println("Failed to send message to client:", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}

func handleClientMessages(conn *websocket.Conn) {
	for {
		// 检查连接状态
		if _, _, err := conn.NextReader(); err != nil {
			log.Println("Connection closed:", err)
			break
		}

		// 从连接中接收消息
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message from client:", err)
			break
		}

		// 处理接收到的消息
		log.Println("Received message:", string(message))

		// 将消息发送给所有连接的前端
		sendMessageToAll(string(message))
	}
}
