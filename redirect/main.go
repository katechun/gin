package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//http重定向
	r.GET("/redirect",func(c *gin.Context){
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com/")
	})

	//路由重定向
	r.GET("/test",func(c *gin.Context){
		c.Request.URL.Path="/test2"
		r.HandleContext(c)
	})

	r.GET("/test2",func(c *gin.Context){
		c.JSON(200,gin.H{"helo":"world"})
	})


	r.Run(":8080")
}
