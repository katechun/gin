package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
   <title>Htts Test </title>
   <script src="/assets/app.js"></script>
</head>
<body>
   <h1 style="color:red;">Welcome,Ginner!</h1>
</body>
</html>
`))

func Https(c *gin.Context){
	if pusher := c.Writer.Pusher();pusher != nil {
		if err := pusher.Push("/assets/app.js",nil);err!=nil{
			log.Printf("Failed to push:%v",err)
		}
	}
	c.HTML(200,"https",gin.H{
		"status":"success",
	})
}

func main() {
	r := gin.Default()
	r.Static("/assets","github.com/katechun/gin/http2/assets")
	r.SetHTMLTemplate(html)

	r.GET("/",Https)

	r.RunTLS(":8080","github.com/katechun/gin/http2/testdata/server.pem","github.com/katechun/gin/http2/testdata/server.key")
}


