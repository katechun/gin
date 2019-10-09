package main

import "github.com/gin-gonic/gin"

type LoginForm struct {
	User string ` form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context){
	var form LoginForm

	if c.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			c.JSON(200,gin.H{"status":"you are logged in"})
		}else {
			c.JSON(401,gin.H{"status":"unauthorized"})
		}
	}
}

// 测试表单  curl -v --form user=user --form password=password http://192.168.100.35:8080/login
func main() {
	router := gin.Default()
	router.POST("/login",Login)

	router.Run(":8080")
}
