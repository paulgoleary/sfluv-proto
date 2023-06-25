package main

import (
	"github.com/paulgoleary/local-luv-proto/ratio"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// health test
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	ratioGroup := r.Group("ratio")
	ratioGroup.POST("client/sessions", ratio.HandleNewSession)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
