

load_type=web_logs.select("timestamp","load_type")
#生成对类型切割的函数
split_type=functions.split(load_type["load_type"],'/')
#对类型列进行切割
load_type2=load_type.withColumn('load_type',split_type.getItem(0))
load_type2.show()
#生成对时间戳切割的函数
load_type3 = functions.split(load_type2['timestamp'], ' ')
#对时间戳列进行切割
load_type4=load_type2.withColumn('timestamp', load_type3.getItem(0))
load_type4.show()

load_type4=load_type4.groupBy("timestamp","load_type").agg(count("*").alias("count"))
load_type4.write.mode("append").jdbc(url=url, table="type_count", properties=properties)
load_type4.show()