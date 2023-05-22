from pyspark.sql.functions import col
from pyspark.sql import functions
from pyspark.sql import SparkSession
# import pymysql

spark = SparkSession.builder.appName("JDBC Connection").config("spark.driver.extraClassPath", "/Users/qingyanchen/Downloads/mysql-connector-java-8.0.25.jar").getOrCreate()

web_logs = spark.read.json('web_log1.json')
web_logs2=web_logs.select("timestamp","response_content_length")
split_col = functions.split(web_logs2['timestamp'], ' ')
web_logs2=web_logs2.withColumn('year', functions.split(web_logs2['timestamp'], '-').getItem(0))
web_logs2=web_logs2.withColumn('month', functions.split(web_logs2['timestamp'], '-').getItem(1))
web_logs2=web_logs2.withColumn('day', functions.split(web_logs2['timestamp'], '-').getItem(1))

web_logs2=web_logs2.withColumn('time', functions.split(web_logs2['timestamp'], ' ').getItem(1))
web_logs2=web_logs2.withColumn('response_content_length',functions.split(web_logs2['response_content_length'],'bytes').getItem(0)).orderBy("timestamp")
web_logs2.show()


web_logs3=web_logs2.withColumn('timestamp',functions.concat_ws('-',web_logs2['year'],web_logs2['month'])).select("timestamp","response_content_length")
web_logs3=web_logs3.withColumn('response_content_length',col('response_content_length').cast("int"))
web_logs3=web_logs3.groupby("timestamp").count().alias("count").orderBy("timestamp")
# web_logs3.write.mode("append").jdbc(url=url, table="time_sum", properties=properties)
# web_logs3=web_logs3.select(web_logs3["timestamp"],web_logs3["sum(response_content_length)"].alias("response_content_length")).orderBy("timestamp")
web_logs3.show()
web_logs2.write.mode("append").jdbc(url=url, table="time_count", properties=properties)


web_logs2=web_logs2.withColumn('response_content_length',col('response_content_length').cast("int"))
web_logs2=web_logs2.groupby("timestamp").sum("response_content_length").alias("response_content_length")
web_logs2=web_logs2.select(web_logs2["timestamp"],web_logs2["sum(response_content_length)"].alias("response_content_length"))
# web_logs3.write.mode("append").jdbc(url=url, table="time_count", properties=properties)
web_logs2.show()