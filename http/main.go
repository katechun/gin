package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router := gin.Default()

	router.GET("/someGet",getting)
	router.POST("/somePost",posting)
	router.PUT("/somePut",putting)
	router.DELETE("/someDelete",deleting)
	router.PATCH("/somePatch",patching)
	router.HEAD("/someHead",head)
	router.OPTIONS("someOptions",options)

	router.Run(":8080")
}
