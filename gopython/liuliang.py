from pyspark.sql.functions import col
from pyspark.sql import functions
from pyspark.sql import SparkSession
# import pymysql

spark = SparkSession.builder.appName("JDBC Connection").config("spark.driver.extraClassPath", "/Users/qingyanchen/Downloads/mysql-connector-java-8.0.25.jar").getOrCreate()

web_logs = spark.read.json('web_log1.json')
web_logs2=web_logs.select("timestamp","response_content_length")
split_col = functions.split(web_logs2['timestamp'], ' ')
web_logs2=web_logs2.withColumn('timestamp', split_col.getItem(0))

web_logs2=web_logs2.withColumn('response_content_length',functions.split(web_logs2['response_content_length'],'bytes').getItem(0))


web_logs2=web_logs2.withColumn('response_content_length',col('response_content_length').cast("int"))
web_logs2=web_logs2.groupby("timestamp").sum("response_content_length").alias("response_content_length")
web_logs2=web_logs2.select(web_logs2["timestamp"],web_logs2["sum(response_content_length)"].alias("response_content_length"))
url = "jdbc:mysql://localhost:3306/testdata"

# Define the properties for the JDBC driver
properties = {
    "user": "root",
    "password": "123456",
    "driver": "com.mysql.jdbc.Driver"
}
web_logs2.write.mode("append").jdbc(url=url, table="time_sum", properties=properties)