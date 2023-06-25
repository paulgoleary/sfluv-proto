package ratio

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swagger "github.com/paulgoleary/local-luv-proto/ratio/go-client-generated"
	"net/http"
)

func HandleNewSession(c *gin.Context) {

	b := swagger.AuthenticateCryptoWalletStartRequest{}
	if err := c.BindJSON(&b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := getDefaultClient()
	if challenge, err := client.authWalletStart(&b); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf(`{"challenge":"%v"}`, challenge))
	}
	return
}
