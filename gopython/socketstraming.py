from pyspark.sql import SparkSession
from pyspark.streaming import StreamingContext

# 创建 SparkSession
spark = SparkSession.builder.appName("SocketStreamingExample").getOrCreate()

# 创建 StreamingContext，设置批处理间隔时间
ssc = StreamingContext(spark.sparkContext, batchDuration=1)
# 创建一个输入数据流，指定监听的主机和端口号
lines = ssc.socketTextStream("localhost", 9999)
# 对输入数据流进行转换和处理
# 例如，可以使用 map、flatMap、filter 等操作来处理每个批次的数据
words = lines.flatMap(lambda line: line.split(" "))
wordCounts = words.map(lambda word: (word, 1)).reduceByKey(lambda x, y: x + y)
# 输出处理结果
# 例如，可以打印每个批次的结果
wordCounts.pprint()
# 启动 StreamingContext
ssc.start()
# 等待 StreamingContext 结束
ssc.awaitTermination()
