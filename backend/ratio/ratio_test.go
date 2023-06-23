package ratio

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/paulgoleary/sfluv-proto/crypto"
	swagger "github.com/paulgoleary/sfluv-proto/ratio/go-client-generated"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"math/big"
	"os"
	"testing"
)

const sandboxUrl = "https://api.sandbox.ratio.me"

//	export const hashPersonalMessage = function (message: Buffer): Buffer {
//	 assertIsBuffer(message)
//	 const prefix = Buffer.from(`\u0019Ethereum Signed Message:\n${message.length}`, 'utf-8')
//	 return keccak(Buffer.concat([prefix, message]))
//	}
func ethSign(msg string, privateKey *ecdsa.PrivateKey) (sig []byte, err error) {
	message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	hash := ethgo.Keccak256([]byte(message))

	if sig, err = crypto.Sign(privateKey, hash); err != nil {
		return
	}
	sig[64] += 27
	return
}

func TestRatioAPI(t *testing.T) {

	ratioClientId := os.Getenv("RATIO_CLIENT_ID")
	ratioClientSecret := os.Getenv("RATIO_CLIENT_SECRET")

	cfg := swagger.NewConfiguration()
	cfg.BasePath = sandboxUrl

	c := swagger.NewAPIClient(cfg)

	sk, err := crypto.SKFromInt(big.NewInt(1))
	require.NoError(t, err)

	addr := crypto.PubKeyToAddress(&sk.PublicKey)

	// TODO: 	WalletNetwork string `json:"walletNetwork"` - OpenAPI is old?
	b := swagger.AuthenticateCryptoWalletStartRequest{
		WalletAddress: addr.String(),
		WalletNetwork: "POLYGON",
	}
	println(addr.String())
	resp, httpResp, err := c.AuthApi.V1AuthCryptoWalletstartPost(context.Background(), b, ratioClientId, ratioClientSecret)
	require.NoError(t, err)
	_ = httpResp

	println(resp.Challenge)
	sig, err := ethSign(resp.Challenge, sk)
	require.NoError(t, err)

	ba := swagger.AuthenticateCryptoWalletRequest{
		WalletAddress: addr.String(),
		WalletNetwork: "POLYGON",
		Signature:     hex.EncodeToString(sig),
	}
	authResp, httpResp, err := c.AuthApi.V1AuthCryptoWalletauthenticatePost(context.Background(), ba, ratioClientId, ratioClientSecret)
	require.NoError(t, err)
	println(authResp.SessionJwt)
}
