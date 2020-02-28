package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main()  {
	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"index": "a",
			"index-text" : 33,
		})
	})
	
	r.POST("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"index": "a",
			"index-text" : 33,
		})
	})

	r.Run(":9090")
}