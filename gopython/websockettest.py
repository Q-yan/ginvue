import asyncio
import websockets

# 保存所有 WebSocket 连接对象的集合
connected_clients = set()

# 定义 WebSocket 服务器的处理逻辑
async def handle_websocket(websocket, path):
    # 将连接对象添加到集合中
    connected_clients.add(websocket)

    try:
        while True:
            # 接收客户端的消息
            message = await websocket.recv()
            print(f"Received message: {message}")

            # 向所有连接的前端客户端发送消息
            await send_to_all_clients(message)
    finally:
        # 连接关闭时从集合中移除连接对象
        connected_clients.remove(websocket)

# 向所有连接的前端客户端发送消息
async def send_to_all_clients(message):
    for client in connected_clients:
        await client.send(message)

# 启动 WebSocket 服务器
start_server = websockets.serve(handle_websocket, 'localhost', 8765)

# 运行事件循环
asyncio.get_event_loop().run_until_complete(start_server)
asyncio.get_event_loop().run_forever()
