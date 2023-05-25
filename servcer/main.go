package main

import (
	"fmt"
	"ginvue/app"
	"ginvue/db"
	"ginvue/moudle"
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

func reverseArray(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
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

		items := moudle.Item{
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

	engine.GET("/data2", func(c *gin.Context) {
		//fmt.Println(time.Now().Format("15:04:05"))
		ite1 := moudle.Ites{
			Date: time.Now().Format("15:04:05"),
			Bar:  rand.Intn(900),
			Line: rand.Intn(30),
		}
		fmt.Println(ite1)
		c.JSON(http.StatusOK, ite1)
	})

	engine.POST("/liuliang", func(c *gin.Context) {

		appG := app.Gin{C: c}

		var requestData *moudle.RequestData
		c.ShouldBindJSON(&requestData)

		lls := []moudle.Liuliang{}

		data := new(struct {
			Times   []string
			Lengths []int
		})

		start_date := "2022-01-01 00:00:00"
		end_date := "2023-01-01 00:00:00"
		if requestData != nil {
			start_date = requestData.Start
			end_date = requestData.End
		}
		err := db.MasterDB.Table("time_sum").Where("timestamp between ? and ?",
			start_date, end_date).OrderBy("timestamp").Find(&lls)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, ll := range lls {
			data.Times = append(data.Times, ll.Timestamp)
			data.Lengths = append(data.Lengths, ll.ResponseContentLength)
		}

		appG.ResponseSucMsg(data)

	})

	engine.GET("/fenlei", func(c *gin.Context) {

		fenleis := moudle.Fenlei_data{}

		for i := 0; i < 7; i++ {

			fenleis.Yaowen = append(fenleis.Yaowen, rand.Intn(90))
			fenleis.Dangzheng = append(fenleis.Dangzheng, rand.Intn(80))
			fenleis.Difang = append(fenleis.Difang, rand.Intn(500))
			fenleis.Guandian = append(fenleis.Guandian, rand.Intn(100))
			fenleis.Qita = append(fenleis.Qita, rand.Intn(100))

		}

		//fmt.Println(fenleis)

		c.JSON(http.StatusOK, fenleis)

	})

	engine.POST("/relitu", func(c *gin.Context) {
		appG := app.Gin{C: c}

		var requestData *moudle.RequestData
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
		//fmt.Println(end_date)
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
		var arrs [][]float64
		countMap := make(map[string]int)
		for _, subArr := range jingweis {
			// 将二维数组转换为字符串表示，作为 map 的键
			key := fmt.Sprintf("%v", subArr)
			// 统计相同数组的数量
			countMap[key]++
			subArr = append(subArr, float64(countMap[key]))
			arrs = append(arrs, subArr)
		}

		appG.ResponseSucMsg(arrs)
	})
	//ip2eo()

	engine.POST("/duibi", func(c *gin.Context) {
		appG := app.Gin{C: c}
		var requestData *moudle.RequestData

		//var setime *startEnd
		err := c.ShouldBindJSON(&requestData)
		if err != nil {
			fmt.Println(err)
			return
		}
		start_date := "2022"
		end_date := "2023"
		if requestData.Start != "" {
			start_date = requestData.Start[:4]
			end_date = requestData.End[:4]
		}
		//var data1 = new(struct {
		//	Data   []string
		//	Values []int
		//})
		data := db.Selecttb(start_date, end_date)
		//
		//err = db.MasterDB.Table("time_count").GroupBy("year").Count(&starts).Where("timestamp like '?%'", start_date).Find(&starts)
		//err = db.MasterDB.Table("time_count").Cols("year").Where("timestamp like '?%'", end_date).Find(&ends)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}

		appG.ResponseSucMsg(data)
	})

	engine.Run(":9001")

}
