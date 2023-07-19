package ratio

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swagger "github.com/paulgoleary/local-luv-proto/ratio/go-client-generated"
	"github.com/paulgoleary/local-luv-proto/util"
	"net/http"
	"strings"
)

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// TODO: not sure if i want to do full auth here given that Ratio does so. but maybe some lightweight auth to
//  validate that the token is not garbage or trivially bad - i.e. expired

const contextJWT = "jwt"

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	if tokenString == "" {
		return fmt.Errorf("invalid a/o missing auth token")
	}
	c.Set(contextJWT, tokenString)
	//_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	//	}
	//	return []byte(os.Getenv("API_SECRET")), nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := TokenValid(c); err != nil {
			c.JSON(http.StatusUnauthorized, `{"error":"Unauthorized"}`)
			c.Abort()
			return
		}
		c.Next()
	}
}

func HandleNewSession(c *gin.Context) {
	b := swagger.AuthenticateCryptoWalletStartRequest{}
	if err := c.BindJSON(&b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := getDefaultClient("")
	if challenge, err := client.authWalletStart(&b); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf(`{"challenge":"%v"}`, challenge))
	}
	return
}

func HandleSessionWallet(c *gin.Context) {
	b := swagger.AuthenticateCryptoWalletRequest{}
	if err := c.BindJSON(&b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := getDefaultClient("")
	if jwt, maybeUserId, err := client.authWalletSignature(&b); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		if maybeUserId != "" {
			client = getDefaultClient(jwt)
			var user swagger.User
			if user, err = client.getUser(maybeUserId); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			} else {
				c.JSON(http.StatusOK, fmt.Sprintf(`{"jwt":"%v", "userId":"%v", "phoneNumber":"%v"}`, jwt, maybeUserId, user.Phone))
			}
		} else {
			c.JSON(http.StatusOK, fmt.Sprintf(`{"jwt":"%v"}`, jwt))
		}
	}
	return
}

func HandleSMSSend(c *gin.Context) {
	jwt := c.GetString(contextJWT)
	b := swagger.SendSmsOtpRequest{}
	if err := c.BindJSON(&b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := getDefaultClient(jwt)
	if phoneId, err := client.authSmsOtpSend(&b); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf(`{"phoneId":"%v"}`, phoneId))
	}
}

type fullAuthUserResult struct {
	Jwt    string        `json:"jwt,omitempty"`
	UserId string        `json:"userId,omitempty"`
	User   *swagger.User `json:"user,omitempty"`
}

func HandleSMSAuth(c *gin.Context) {
	jwtIn := c.GetString(contextJWT)
	b := swagger.AuthenticateSmsOtpRequest{}
	if err := c.BindJSON(&b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := getDefaultClient(jwtIn)
	if jwtOut, maybeUser, err := client.authSmsOtpAuth(&b); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		if maybeUser == nil {
			c.JSON(http.StatusOK, fmt.Sprintf(`{"jwt":"%v"}`, jwtOut))
		} else {
			res := fullAuthUserResult{
				Jwt:    jwtOut,
				UserId: maybeUser.Id,
				User:   maybeUser,
			}
			c.JSON(http.StatusOK, res)
		}
	}
}

func HandleCreateUser(c *gin.Context) {
	jwtIn := c.GetString(contextJWT)
	b := swagger.CreateUserRequest{}
	if err := c.BindJSON(&b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := getDefaultClient(jwtIn)
	maybeAddr := util.JWTExtractRatioAuthAddress(jwtIn)
	if user, err := client.authCreateUser(&b, maybeAddr); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf(`{"userId":"%v"}`, user.Id))
	}
}

func HandleUserKyc(c *gin.Context) {
	jwtIn := c.GetString(contextJWT)
	b := swagger.SubmitKycRequest{}
	if err := c.BindJSON(&b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := getDefaultClient(jwtIn)
	userId, _ := c.Params.Get("userId")
	if user, err := client.authUserKyc(&b, userId); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf(`{"userId":"%v"}`, user.Id))
	}
}
