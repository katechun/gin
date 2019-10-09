package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func Jsonp(c *gin.Context){
	data := map[string]interface{}{
		"foo":"bar",
	}

	c.JSONP(http.StatusOK,data)
}

func main() {
	r := gin.Default()

	r.GET("/JSONP?callback=x",Jsonp)

	r.Run(":8080")
}
