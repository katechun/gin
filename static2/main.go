package main

import (
	"github.com/alecthomas/template"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	r:= gin.New()

	t,err := loadTemplate()
	if err != nil {
		panic(err)
	}

	r.SetHTMLTemplate(t)

	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK,"/html/index.tmpl",nil)
	})

	r.Run(":8080")

}


func loadTemplate()(*template.Template,error) {
	t := template.New("")
	for name,file := range Assets.Files{
		if file.IsDir() || !strings.HasSuffix(name,".tmpl"){
			continue
		}
		h,err := ioutil.ReadAll(file)
		if err != nil {
			return  nil ,err
		}

		t,err := t.New(name).Parse(string(h))
		if err != nil {
			return  nil,err
		}
	}
	return t,nil
}