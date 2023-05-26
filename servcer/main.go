package main

import (
	"fmt"
	"ginvue/app"
	"ginvue/db"
	"ginvue/moudle"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
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
	engine.POST("/data", func(c *gin.Context) {

		appG := app.Gin{C: c}

		var data []*moudle.Difang_sum
		var requestData *moudle.RequestData
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
		//select load_type,sum(`count(1)`) as sum from difang_sum where timestamp like '2016-%' group by load_type
		err = db.MasterDB.Table("difang_sum").Select("load_type as name,sum(`count(1)`) as value").
			Where("timestamp  between ? and ?", start_date, end_date).GroupBy("load_type").Find(&data)
		if err != nil {
			return

		}

		appG.ResponseSucMsg(data)
	})

	engine.POST("/liuliang", func(c *gin.Context) {

		appG := app.Gin{C: c}

		//var requestData []string
		var requestData *moudle.RequestData1

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

		var yaowens []moudle.TypeCount
		var dnagzhengs []moudle.TypeCount
		var difangs []moudle.TypeCount
		var guangdians []moudle.TypeCount

		err := db.MasterDB.Table("type_count").Select("timestamp, SUM(count) AS count").Where("load_type LIKE ?", "党政").
			GroupBy("timestamp").OrderBy("timestamp DESC").Limit(7).Find(&yaowens)
		if err != nil {
			return
		}

		for _, yaowen := range yaowens {
			fenleis.Yaowen = append(fenleis.Yaowen, yaowen.Count)
		}
		err = db.MasterDB.Table("type_count").Select("timestamp, SUM(count) AS count").Where("load_type LIKE ?", "要闻").
			GroupBy("timestamp").OrderBy("timestamp DESC").Limit(7).Find(&dnagzhengs)
		if err != nil {
			return
		}

		for _, item := range dnagzhengs {
			fenleis.Dangzheng = append(fenleis.Dangzheng, item.Count)
		}
		err = db.MasterDB.Table("type_count").Select("timestamp, SUM(count) AS count").Where("load_type LIKE ?", "地方").
			GroupBy("timestamp").OrderBy("timestamp DESC").Limit(7).Find(&difangs)
		if err != nil {
			return
		}

		for _, item := range difangs {
			fenleis.Difang = append(fenleis.Difang, item.Count)
		}

		err = db.MasterDB.Table("type_count").Select("timestamp, SUM(count) AS count").Where("load_type LIKE ?", "观点").
			GroupBy("timestamp").OrderBy("timestamp DESC").Limit(7).Find(&guangdians)
		if err != nil {
			return
		}
		for _, item := range guangdians {
			fenleis.Guandian = append(fenleis.Guandian, item.Count)
			fenleis.Qita = append(fenleis.Qita, rand.Intn(100))
			fenleis.Date = append(fenleis.Date, item.Timestamp)
		}

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

		data := db.Selecttb(start_date, end_date)

		appG.ResponseSucMsg(data)
	})

	engine.Run(":9001")

}
