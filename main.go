package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

//GETDate
//gdg
func GetDate() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"date": time.Now().Format("2006/01/02"),
		})
	}
}

//
func GetTime() func(c *gin.Context) {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time": time.Now().Format(time.Kitchen),
		})
	}
}
func main() {

	r := gin.Default()

	r.GET("/time", GetTime())
	r.POST("/date", GetDate())

	r.Run()
}
