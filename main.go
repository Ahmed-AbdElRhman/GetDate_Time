package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func GetDate() string {
	return time.Now().Format("2006/01/02")
}

func GetTime() string {
	return time.Now().Format(time.Kitchen)
}

func GetApi() func(c *gin.Context) {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time": GetTime(),
			"date": GetDate(),
		})
	}
}
func main() {

	r := gin.Default()

	r.GET("/Api", GetApi())

	r.Run()
}
