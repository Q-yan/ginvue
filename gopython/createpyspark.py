from pyspark.sql import SparkSession
# import pymysql

spark = SparkSession.builder.appName("weblog").config("spark.driver.extraClassPath", "/Users/qingyanchen/Downloads/mysql-connector-java-8.0.25.jar").getOrCreate()

#
# # 设置 JDBC 配置信息
# jdbc_url = "jdbc:mysql://localhost:3306/testdata"
# jdbc_driver = "com.mysql.jdbc.Driver"
# jdbc_username = "root"
# jdbc_password = "123456"
#
# # 创建 DataFrame
# df = spark.read.jdbc(url=jdbc_url, table="mytable", properties={"driver": jdbc_driver, "user": jdbc_username, "password": jdbc_password})
# df.show()

spark.streams.awaitAnyTermination()

