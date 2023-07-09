package main

import (
	"fmt"
	"github.com/paulgoleary/local-luv-proto/ratio"
	"github.com/paulgoleary/local-luv-proto/util"
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
	ratioGroup.POST("sessions", ratio.HandleNewSession)
	ratioGroup.POST("wallet", ratio.HandleSessionWallet)

	ratioAuth := ratioGroup.Group("auth")
	ratioAuth.Use(ratio.JwtAuthMiddleware())
	ratioAuth.POST("sms", ratio.HandleAuthSMS)

	return r
}

func main() {
	r := setupRouter()
	localIp := util.GetOutboundIP()
	println(fmt.Sprintf("server starting at local IP %v", localIp.String())) // TODO: logging...
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
