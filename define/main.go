package main

import (
	"github.com/gin-gonic/gin"
     "html/template"
	)

func main() {
	router := gin.Default()
	 ff:= template.Must(template.ParseFiles("github.com/katechun/gin/define/template/file1","github.com/katechun/gin/define/template/file2"))
	router.SetHTMLTemplate(ff)
	router.Run(":8080")
}
