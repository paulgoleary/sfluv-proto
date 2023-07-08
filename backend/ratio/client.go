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

func (c *ratioClient) authWalletStart(b *swagger.AuthenticateCryptoWalletStartRequest) (challenge string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.to)
	defer cancel()
	if resp, _, apiErr := c.c.AuthApi.V1AuthCryptoWalletstartPost(ctx, *b, c.ratioClientId, c.ratioClientSecret); apiErr != nil {
		log.Error(fmt.Sprintf("V1AuthCryptoWalletstartPost failed with error %v", apiErr))
		err = apiErr
	} else {
		challenge = resp.Challenge
	}
	return
}