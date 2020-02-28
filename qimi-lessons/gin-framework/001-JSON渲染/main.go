package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// gin.H 是map[string]interface{}的缩写
	r.GET("/someJSON", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		var json struct {
			Name   string `json:"user"`
			Gender string
			Age    int
		}
		json.Name = "panda"
		json.Gender = "男"
		json.Age = 18

		// c.BindJSON(&json)
		c.JSON(http.StatusOK, json)
	})

	r.Run()
}
