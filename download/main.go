package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/download",Download)

	r.Run(":8080")
}


func Download(c *gin.Context){

	filename := "aaa.zip"
	c.Writer.Header().Add("Content-Disposition",fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("github.com/katechun/gin/download/tmp/aaa.zip")

}