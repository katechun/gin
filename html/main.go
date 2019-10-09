package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router :=gin.Default()
	router.LoadHTMLGlob("github.com/katechun/gin/html/templates/**/*")

	router.GET("/index",Index)
	router.GET("/test",Test)

	router.Run(":8080")
}

func Index(c *gin.Context){
	c.HTML(http.StatusOK,"index.tmpl",gin.H{
		"title":"Main website",
	})

}

func Test(c *gin.Context){
	c.HTML(http.StatusOK,"test.tmpl",gin.H{
		"title":"Test website",
	})

}
