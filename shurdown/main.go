package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/",func(c *gin.Context){
		time.Sleep(5*time.Second)
		c.String(http.StatusOK,"welcome Gin Server")
	})

	srv := &http.Server{
		Addr: ":8080",
		Handler:router,

	}

	go func (){
		if err := srv.ListenAndServe();err != nil && err != http.ErrServerClosed{
			log.Fatalf("listen: %s",err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx);err != nil {
		log.Fatal("Server Shutdown:",err)
	}

	log.Println("Server exiting")


}
