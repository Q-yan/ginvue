import asyncio
import websockets

# 保存所有 WebSocket 连接对象的字典
connected_clients = {}

# 定义 WebSocket 服务器的处理逻辑
async def handle_websocket(websocket, path):
    # 将连接对象保存到字典中
    connected_clients[id(websocket)] = websocket
    print(connected_clients)

    try:
        while True:
            # 接收客户端的消息
            message = await websocket.recv()
            print(f"Received message: {message}")

            # 回复客户端的消息
            reply_message = f"Server received: {message}"
            await websocket.send(reply_message)

            # 向指定的客户端发送消息
            client_id = id(websocket)  # 替换为指定客户端的标识符
            client = connected_clients.get(client_id)
            if client:
                await client.send(message)
    finally:
        # 连接关闭时从字典中移除连接对象
        del connected_clients[id(websocket)]

# 向指定的客户端发送消息
async def send_to_client(client_id, message):
    client = connected_clients.get(client_id)
    if client:
        await client.send(message)

# 启动 WebSocket 服务器
start_server = websockets.serve(handle_websocket, 'localhost', 8765)

# 运行事件循环
asyncio.get_event_loop().run_until_complete(start_server)
asyncio.get_event_loop().run_forever()
