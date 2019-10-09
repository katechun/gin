package ascii

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/someJSON",SomeJson)

	r.Run(":8080")
}


func SomeJson( c *gin.Context){
	data := map[string]interface{}{
		"lang":"GO语言",
		"tag":"<br>",
	}

	c.AsciiJSON(http.StatusOK,data)

}
