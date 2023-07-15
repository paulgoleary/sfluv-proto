package ratio

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/apex/log"
	swagger "github.com/paulgoleary/local-luv-proto/ratio/go-client-generated"
	"github.com/paulgoleary/local-luv-proto/util"
	"os"
	"time"
)

const sandboxUrl = "https://api.sandbox.ratio.me"

type ratioClient struct {
	ratioClientId     string
	ratioClientSecret string

	c *swagger.APIClient

	to time.Duration
}

func getDefaultClient(maybeJwt string) *ratioClient {

	c := &ratioClient{}

	c.ratioClientId = os.Getenv("RATIO_CLIENT_ID")
	c.ratioClientSecret = os.Getenv("RATIO_CLIENT_SECRET")
	if c.ratioClientId == "" || c.ratioClientSecret == "" {
		log.Error("client is misconfigured - missing client ID a/o secret") // TODO: return error?
	}

	cfg := swagger.NewConfiguration()
	cfg.BasePath = sandboxUrl

	localIp := util.GetOutboundIP()
	fingerPrintJson := fmt.Sprintf(`{"ip":"%v","userAgent":"%v"}`, localIp.String(), cfg.UserAgent)
	fingerPrintEnc := base64.StdEncoding.EncodeToString([]byte(fingerPrintJson))
	cfg.DefaultHeader["ratio-device-fingerprint"] = fingerPrintEnc

	if maybeJwt != "" {
		cfg.DefaultHeader["Authorization"] = "Bearer " + maybeJwt
	}

	c.c = swagger.NewAPIClient(cfg)

	c.to = 10 * time.Second

	return c
}

func handleApiError(apiName string, err error) {
	if swagErr, ok := err.(swagger.GenericSwaggerError); ok {
		log.Info(fmt.Sprintf("%v failed with error %v, %v", apiName, err, string(swagErr.Body())))
		println(string(swagErr.Body()))
	} else {
		log.Info(fmt.Sprintf("%v failed with error %v", apiName, err))
	}
}

func (c *ratioClient) authWalletStart(b *swagger.AuthenticateCryptoWalletStartRequest) (challenge string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.to)
	defer cancel()
	var resp swagger.AuthenticateCryptoWalletStartResponse
	if resp, _, err = c.c.AuthApi.V1AuthCryptoWalletstartPost(ctx, *b, c.ratioClientId, c.ratioClientSecret); err != nil {
		handleApiError("V1AuthCryptoWalletstartPost", err)
	} else {
		challenge = resp.Challenge
	}
	return
}

func (c *ratioClient) authWalletSignature(ba *swagger.AuthenticateCryptoWalletRequest) (jwt, maybeUserId string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.to)
	defer cancel()
	var authResp swagger.AuthResponse
	if authResp, _, err = c.c.AuthApi.V1AuthCryptoWalletauthenticatePost(ctx, *ba, c.ratioClientId, c.ratioClientSecret); err != nil {
		handleApiError("V1AuthCryptoWalletauthenticatePost", err)
	} else {
		jwt = authResp.SessionJwt
	}
	return
}

func (c *ratioClient) authSmsOtpSend(ba *swagger.SendSmsOtpRequest) (phoneId string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.to)
	defer cancel()
	var sendOtpResp swagger.SendSmsOtpResponse
	if sendOtpResp, _, err = c.c.AuthApi.V1AuthOtpSmssendPost(ctx, *ba, c.ratioClientId, c.ratioClientSecret); err != nil {
		handleApiError("V1AuthOtpSmssendPost", err)
	} else {
		phoneId = sendOtpResp.PhoneId
	}
	return
}

func (c *ratioClient) authSmsOtpAuth(ba *swagger.AuthenticateSmsOtpRequest) (jwt string, maybeUser *swagger.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.to)
	defer cancel()
	var authOtpResp swagger.AuthResponse
	if authOtpResp, _, err = c.c.AuthApi.V1AuthOtpSmsauthenticatePost(ctx, *ba, c.ratioClientId, c.ratioClientSecret); err != nil {
		handleApiError("V1AuthOtpSmsauthenticatePost", err)
	} else {
		jwt = authOtpResp.SessionJwt
		maybeUser = authOtpResp.User
	}
	return
}

func (c *ratioClient) authCreateUser(ba *swagger.CreateUserRequest, maybeAddr string) (user swagger.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.to)
	defer cancel()
	if user, _, err = c.c.UserApi.V1UsersPost(ctx, *ba, c.ratioClientId, c.ratioClientSecret); err != nil {
		handleApiError("V1UsersPost", err)
		return
	}
	// if default address is provided, attach to user
	if maybeAddr != "" {
		b := swagger.ConnectWalletRequest{
			Address: maybeAddr,
			Type_:   "POLYGON",
			Name:    "SFLUV Default Wallet",
		}
		if _, _, err = c.c.WalletApi.V1UsersUserIdWalletsPost(ctx, b, c.ratioClientId, c.ratioClientSecret, user.Id); err != nil {
			handleApiError("V1UsersUserIdWalletsPost", err)
			return
		}
	}
	return
}

func (c *ratioClient) getUser(userId string) (user swagger.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.to)
	defer cancel()
	if user, _, err = c.c.UserApi.V1UsersUserIdGet(ctx, userId, c.ratioClientId, c.ratioClientSecret); err != nil {
		handleApiError("V1UsersUserIdGet", err)
	}
	return
}
