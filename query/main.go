package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//测试表单   curl -v --form name=manu --form message=this_is_great http://192.168.100.35:8080/post?id=1234&page=1
func main() {
	router := gin.Default()

	router.POST("/post",func(c *gin.Context){
		id := c.Query("id")
		page := c.DefaultQuery("page","0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s",id,page,name,message)
	})

	router.Run(":8080")
}
