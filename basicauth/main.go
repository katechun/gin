package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = gin.H{
	"foo":gin.H{"email":"foo@bar.com","phone":"123433"},
	"austin":gin.H{"email":"aistom@bar.com","phone":"6666"},
	"lena":gin.H{"email":"lena@bar.com","phone":"523443"},
}


func main() {

	r := gin.Default()

	authorized := r.Group("/admin",gin.BasicAuth(gin.Accounts{
		"foo":"bar",
		"austin":"1234",
		"lena":"hello2",
		"manu":"4321",
	}))

	authorized.GET("/secrets",func (c *gin.Context){
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret,ok := secrets[user];ok{
			c.JSON(http.StatusOK,gin.H{
				"user":user,
				"secret":secret,
			})
		}else {
			c.JSON(http.StatusOK,gin.H{"user":user,"secret":"NO SECRET :("})
		}
	})

	r.Run(":8080")
}