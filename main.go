package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func InitializeEngine() *gin.Engine {
	s := gin.Default()

	s.GET("/hello", GetHelloPage)
	return s
}

func GetHelloPage(c *gin.Context) {
	// Implementation Details
	// Returns a Page
	logger, _ := zap.NewProduction()
	resp := "Hello go"
	logger.Info("successfully got request")
	c.JSON(http.StatusOK, &resp)
}

func main() {
	s := InitializeEngine()
	go s.Run(":8080")
	fmt.Println("hello-go")
}
