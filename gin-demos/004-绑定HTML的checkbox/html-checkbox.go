package main

import (
	"github.com/gin-gonic/gin"

)

type myForm struct {
    Colors []string `form:"colors[]"`
}


func formHandler(c *gin.Context) {
    var fakeForm myForm
    c.ShouldBind(&fakeForm)
    c.JSON(200, gin.H{"color": fakeForm.Colors})
}

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("./*")
    r.GET("/", indexHandler)
	r.POST("/", formHandler)
	
    r.Run() // listen and serve on 0.0.0.0:8080
}

func indexHandler(c *gin.Context) {
    c.HTML(200, "form.html", nil)
}