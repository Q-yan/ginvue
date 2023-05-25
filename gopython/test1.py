import socket

# 创建一个TCP/IP服务器
server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# 绑定服务器地址和端口
server_address = ('localhost', 9999)
server_socket.bind(server_address)

# 开始监听连接请求
server_socket.listen(1)
print('服务器正在运行...')

# 等待客户端连接
while True:
    client_socket, client_address = server_socket.accept()
    print(f'客户端 {client_address} 已连接')

    # 接收并处理客户端发送的数据
    while True:
        data = client_socket.recv(1024)
        if not data:
            break
        print(data.decode('utf-8'))

    # 关闭客户端连接
    client_socket.close()
    print(f'客户端 {client_address} 已断开')