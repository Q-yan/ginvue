from pyspark.sql import SparkSession
from pyspark.streaming import StreamingContext
import socket
import json
import json
import websocket


# 创建 SparkSession
spark = SparkSession.builder.appName("SocketStreamingExample").getOrCreate()

# 创建 StreamingContext，设置批处理间隔时间
ssc = StreamingContext(spark.sparkContext, batchDuration=5)
# 创建一个输入数据流，指定监听的主机和端口号
lines = ssc.socketTextStream("localhost", 9919)


# 创建 WebSocket 连接
ws = websocket.WebSocket()
ws.connect("ws://localhost:8765/ws")  # 替换为实际的 WebSocket 地址


# 定义结束作业的函数
def stop_streaming():
    ssc.stop()
    print("Streaming job stopped.")


# 在另一个线程中监听用户输入
import threading
def read_input():
    while True:
        user_input = input()
        if user_input.lower() == "quit":
            stop_streaming()
            break

# 在每个批次中计算行数
def process_batch(batch):
    # 获取当前批次的行数
    count = batch.count()
    # 打印行数
    print("访问量:", count)
    # 将数据转换为 JSON 字符串
    #     json_data = json.dumps(rdd.collect())
    # 发送数据给 WebSocket
    ws.send(str(count).encode('utf-8'))



    # 处理数据的函数
def process_data(rdd):
    # 将数据转换为 JSON 字符串
    json_data = json.dumps(rdd.collect())

    # 发送数据给 WebSocket
    ws.send(json_data)

# 在 DStream 上应用 process_batch 函数
lines.foreachRDD(process_batch)
# 启动 Spark Streaming
ssc.start()
ssc.awaitTermination()