package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"test/api"
	"test/config"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	r := gin.New()
	r.Use(gin.Recovery())

	// log
	r.Use(gin.Logger())
	log.SetOutput(gin.DefaultWriter)

	// cors
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"POST", "HEAD", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type",
			"Authorization", "X-Real-Ip",
			"X-Appengine-Remote-Addr", "Access-Control-Allow-Origin"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	// redirect
	r.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/version")
	})

	group := r.Group("/dingtalk")
	group.POST("send", api.DingDingTalk)

	// gracefully shutdown, work on go1.8+
	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", "0.0.0.0", config.Port),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen server err", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err.Error())
	}
	log.Println("Server exiting")
}
