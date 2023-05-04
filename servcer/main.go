package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


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
	Date   []string    `json:"date"`
	Num 	[]int `json:"num"`

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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Next()
	})
//初始化动态折线图
	engine.GET("/data", func(c *gin.Context) {


		items := Item{
			Date: []string{},
			Num: []int{},
		}
		now := time.Now()

		for i := 9; i >= 0; i-- {

			ite := now.Format("15:04:05")
			items.Date= append(items.Date, ite)
			items.Num=append(items.Num,i+1)
			now = now.Add(-2 * time.Second)
		}
		reverseArray(items.Num)
		fmt.Println(items)
		c.JSON(http.StatusOK, items)
	})

	engine.Run(":9001")

}
