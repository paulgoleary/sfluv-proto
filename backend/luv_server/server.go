package main

import (
	"fmt"
	"github.com/paulgoleary/local-luv-proto/erc4337"
	"github.com/paulgoleary/local-luv-proto/util"
	"net/http"

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

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	// health test
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	hc := erc4337.HandlerContext{}

	// GET? erc4337/userop/approve?target=XXXX&spender=YYYY&amount=10000&owner=ZZZZ
	// function approve(address spender, uint256 amount) external returns (bool)

	// GET? erc4337/userop/depositFor
	erc4337Group := r.Group("erc4337")
	erc4337Group.GET("userop/approve", hc.HandleUserOpApprove)

	return r
}

func main() {
	r := setupRouter()
	localIp := util.GetOutboundIP()
	println(fmt.Sprintf("server starting at local IP %v", localIp.String())) // TODO: logging...
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
