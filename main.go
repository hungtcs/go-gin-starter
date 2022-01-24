package main

import (
	"fmt"
	"gin-examples/internal/config"
	"gin-examples/internal/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()

	config := config.New()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.SetTrustedProxies(config.Server.TrustedProxies)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	logger.Logger.Infof("Start server http://%s:%d", config.Server.Host, config.Server.Port)
	if error := router.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)); error != nil {
		logger.Logger.Fatalf("server start failed %w", error)
	}

	go handleShutdown()
}

func handleShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	fmt.Println("Do Something before exit")

	os.Exit(0)
}
