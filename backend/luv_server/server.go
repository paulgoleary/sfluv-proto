package main

import (
	"fmt"
	"github.com/paulgoleary/local-luv-proto/config"
	"github.com/paulgoleary/local-luv-proto/erc4337"
	"github.com/paulgoleary/local-luv-proto/util"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

var DefaultPolygonRpcUrl = "https://polygon-rpc.com"

func setupRouter(hc *erc4337.HandlerContext) *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	// health test
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	erc4337Group := r.Group("erc4337")
	erc4337Group.GET("sender", hc.HandleGetSender)

	erc4337Group.GET("userop/approve", hc.HandleUserOpApprove)
	erc4337Group.GET("userop/withdrawto", hc.HandleUserOpWithdrawTo)

	erc4337Group.PUT("userop/send", hc.HandleUserOpSend)

	return r
}

func main() {

	maybeEnvUrl := os.Getenv("CHAIN_URL")
	if maybeEnvUrl == "" {
		maybeEnvUrl = DefaultPolygonRpcUrl
	}
	cfg := config.Config{ChainRpcUrl: maybeEnvUrl}
	hc, err := erc4337.MakeContext(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := setupRouter(hc)
	localIp := util.GetOutboundIP()
	println(fmt.Sprintf("server starting at local IP %v", localIp.String())) // TODO: logging...
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
