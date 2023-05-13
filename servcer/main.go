package main

import (
	"fmt"
	"ginvue/app"
	"ginvue/db"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	db.Setup()
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

type Item struct {
	Date []string `json:"date"`
	Num  []int    `json:"num"`
	Bar  []int    `json:"bar"`
	Line []int    `json:"line"`
}

type ites struct {
	Date string `json:"date"`
	Num  int    `json:"num"`
	Bar  int    `json:"bar"`
	Line int    `json:"line"`
}

type ll struct {
}

type liuliang struct {
	Date   []string `json:"date"`
	Length []int    `json:"lenthg"`
}

type fenlei_data struct {
	Yaowen    []int `json:"yaowen"`
	Dangzheng []int `json:"dangzheng"`
	Guandian  []int `json:"guandian"`
	Difang    []int `json:"difang"`
	Qita      []int `json:"qita"`
}

func reverseArray(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

type startEnd struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func main() {
	engine := gin.Default()
	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT,DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	//初始化动态折线图
	engine.GET("/data", func(c *gin.Context) {

		items := Item{
			Date: []string{},
			Num:  []int{},
		}
		now := time.Now()

		for i := 9; i >= 0; i-- {

			ite := now.Format("15:04:05")
			items.Date = append(items.Date, ite)
			items.Num = append(items.Num, i+1)
			items.Bar = append(items.Bar, i*100)
			items.Line = append(items.Line, i*5)
			now = now.Add(-2 * time.Second)
		}
		reverseArray(items.Num)
		fmt.Println(items)
		c.JSON(http.StatusOK, items)
	})

	//engine.GET("/data2", func(c *gin.Context) {
	//	//fmt.Println(time.Now().Format("15:04:05"))
	//	ite1 := ites{
	//		Date: time.Now().Format("15:04:05"),
	//		Bar:  rand.Intn(900),
	//		Line: rand.Intn(30),
	//	}
	//	fmt.Println(ite1)
	//	c.JSON(http.StatusOK, ite1)
	//})

	engine.GET("/liuliang", func(c *gin.Context) {
		lls := liuliang{}

		date := []string{}
		data := []int{}

		//startDate := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
		//endDate := time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC)

		for i := 0; i < 200; i++ {
			startDate := time.Date(2010, time.January, i+1, 0, 0, 0, 0, time.UTC)
			date = append(date, startDate.Format("2006/06/02"))
			data = append(data, rand.Intn(1000))
		}
		// 遍历日期区间
		//for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
		//	//fmt.Println(d.Format("2006/06/02"))
		//	lls.Date=append(lls.Date,d.Format("2006/06/02"))
		//	lls.Length=append(lls.Length,rand.Intn(300))

		lls.Date = date
		lls.Length = data
		//fmt.Println(lls)
		c.JSON(http.StatusOK, lls)

	})

	engine.GET("/fenlei", func(c *gin.Context) {

		fenleis := fenlei_data{}

		for i := 0; i < 7; i++ {

			fenleis.Yaowen = append(fenleis.Yaowen, rand.Intn(90))
			fenleis.Dangzheng = append(fenleis.Dangzheng, rand.Intn(80))
			fenleis.Difang = append(fenleis.Difang, rand.Intn(500))
			fenleis.Guandian = append(fenleis.Guandian, rand.Intn(100))
			fenleis.Qita = append(fenleis.Qita, rand.Intn(100))

		}

		fmt.Println(fenleis)

		c.JSON(http.StatusOK, fenleis)

	})

	type RequestData struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}

	engine.POST("/duibi", func(c *gin.Context) {
		//start :=c.PostForm("start")
		//end:=c.PostForm("end")
		//df = web_logs.withColumn("timestamp", to_timestamp(col("timestamp"), "yyyy-MM-dd HH:mm:ss"))
		//start_time = "2022-01-01 00:00:00"
		//end_time = "2024-01-31 23:59:59"
		//web_logs.agg(count("*").alias("count")).show()
		//web_logs.filter((col("timestamp") >= start_time) & (col("timestamp") <= end_time)).agg(count("*").alias("count")).show()

	})

	engine.POST("/relitu", func(c *gin.Context) {
		appG := app.Gin{C: c}

		var requestData *RequestData
		var ips []string

		//var setime *startEnd
		err := c.ShouldBindJSON(&requestData)
		if err != nil {
			fmt.Println(err)
			return
		}

		start_date := "2022-01-01 00:00:00"
		end_date := "2023-01-01 00:00:00"
		if requestData != nil {
			start_date = requestData.Start
			end_date = requestData.End
		}
		fmt.Println(end_date)
		err = db.MasterDB.Table("timeip").Cols("ip_address").Where("timestamp between ? and ?",
			start_date, end_date).Find(&ips)

		//var tcs []struct {
		//	IpAddress string `json:"ip_address"`
		//	Count     int    `json:"count"`
		//}
		//err = db.MasterDB.Table("timeip").Select("ip_address, count(*) as count").Where("timestamp between ? and ?",
		//	start_date, end_date).GroupBy("ip_address").OrderBy("count").Find(&tcs)

		if err != nil {
			fmt.Println(err)
			return
		}

		var jingweis [][]float64
		for _, ip := range ips {
			if ip != "" {
				jingweidu := ip2eo(ip)
				jingweis = append(jingweis, jingweidu)
			}

		}
		appG.ResponseSucMsg(jingweis)
	})
	//ip2eo()

	engine.Run(":9001")

}
