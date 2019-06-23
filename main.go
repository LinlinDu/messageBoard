package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/middleware/logger"
	"message-board/models"
    "message-board/pkg/setting"
    "message-board/routers"
    "net/http"
	"os"
	"os/signal"
	"time"
	"context"
)

func init()  {
	setting.Setup()
	logger.Setup()
	models.Setup()
}

func main() {
	var listenAddr int
	flag.IntVar(&listenAddr, "port", setting.ServerSetting.HTTPPort, "server listen port")
	flag.Parse()

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	gin.SetMode(setting.RunMode)
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", listenAddr),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		<-quit
		log.Println("Server is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		s.SetKeepAlivesEnabled(false)
		if err := s.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()
	log.Println("Server is listening port: ", listenAddr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %d: %v\n", listenAddr, err)
	}

	<-done
	log.Println("Server stopped")
	}
