package ratio

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/paulgoleary/local-luv-proto/crypto"
	swagger "github.com/paulgoleary/local-luv-proto/ratio/go-client-generated"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"math/big"
	"os"
	"strings"
	"testing"
)

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

	sk, err := crypto.SKFromInt(big.NewInt(2))
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

	// need client with jwt auth now ...
	dc := getDefaultClient(authResp.SessionJwt)
	// dc := getDefaultClient("")

	bs := swagger.SendSmsOtpRequest{
		PhoneNumber: "+12059522124",
	}
	sendOtpResp, _, err := dc.c.AuthApi.V1AuthOtpSmssendPost(context.Background(), bs, ratioClientId, ratioClientSecret)
	require.NoError(t, err)
	println(sendOtpResp.PhoneId)
}

func TestSmsOtp(t *testing.T) {

	c := getDefaultClient("eyJhbGciOiJSUzI1NiIsImtpZCI6Imp3ay10ZXN0LWI5NWU0NTM5LTNlZGMtNGJiNS04MjdmLTViMjAyY2NmNDJlNiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsicHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyJdLCJleHAiOjE2ODc3MjA4NTksImh0dHBzOi8vc3R5dGNoLmNvbS9zZXNzaW9uIjp7ImlkIjoic2Vzc2lvbi10ZXN0LTdjZTYwODdjLTY2YWMtNGViNS1hYmViLWJjNjgxOWM0MTExNyIsInN0YXJ0ZWRfYXQiOiIyMDIzLTA2LTI1VDE5OjE1OjU5WiIsImxhc3RfYWNjZXNzZWRfYXQiOiIyMDIzLTA2LTI1VDE5OjE1OjU5WiIsImV4cGlyZXNfYXQiOiIyMDIzLTA2LTI1VDE5OjI1OjU5WiIsImF0dHJpYnV0ZXMiOnsidXNlcl9hZ2VudCI6IiIsImlwX2FkZHJlc3MiOiIifSwiYXV0aGVudGljYXRpb25fZmFjdG9ycyI6W3sidHlwZSI6ImNyeXB0byIsImRlbGl2ZXJ5X21ldGhvZCI6ImNyeXB0b193YWxsZXQiLCJsYXN0X2F1dGhlbnRpY2F0ZWRfYXQiOiIyMDIzLTA2LTI1VDE5OjE1OjU5WiIsImNyeXB0b193YWxsZXRfZmFjdG9yIjp7ImNyeXB0b193YWxsZXRfaWQiOiJjcnlwdG8td2FsbGV0LXRlc3QtOGNmODVkNGYtMDA3YS00NDk5LWI5YWEtMTdkYjViYTMyNDgyIiwiY3J5cHRvX3dhbGxldF9hZGRyZXNzIjoiMHgyQjVBRDVjNDc5NWMwMjY1MTRmODMxN2M3YTIxNUUyMThEY0NENmNGIiwiY3J5cHRvX3dhbGxldF90eXBlIjoiZXRoZXJldW0ifX1dfSwiaWF0IjoxNjg3NzIwNTU5LCJpc3MiOiJzdHl0Y2guY29tL3Byb2plY3QtdGVzdC1mOTk5ZDMzZS00OWRlLTQ4ODktYmVjNi04ZWIyODA0NGQwNDciLCJtb25pdG9yaW5nX3Nlc3Npb25faWQiOiIwZGNhNzc2Mi0xNGJkLTQxZWEtOTU0Yy04MDNlY2NmMWE1NTIiLCJuYmYiOjE2ODc3MjA1NTksInN1YiI6InVzZXItdGVzdC1kOTBhM2E0MS02NWFkLTRjMWItYTY2MC1mN2JhY2JkYTc2ZGIifQ.gDzkGc6sKpjud3RdvJ5dDKjp33QqQTg2f59Ri0NMK53cYKb5p76BeoD-tDDPxZ6HcDJswBRpcE58Byu22yAJL9se4m01atIWbeP78on9_XeHLmCeDrJCxBymmQkIP6x7mrZ95CnT-K5KBD88IVh4ELLGhHCJ45DR_Au_QA7wRXtEYK82yEPWUyO5TzyJmA3WfnJyDUcAdgwkQXLrt3XQQbZcVc_A90xubvSeEeiRRKV1uD4eQ-cU0I2xu6fXa9nwzYI8QI1rNTxg3xqSlBKrwh4pN96Z26HBCQgpiyRiA0GnIQ8XqbFhnhTEYoMY5bx050vVrKR2dtfdZxa5ssrz3g")
	// c := getDefaultClient("")

	bsa := swagger.AuthenticateSmsOtpRequest{
		Otp:     "903097",
		PhoneId: "phone-number-test-0f8a06d2-c184-45b8-a4a1-5d15c775ac3f",
	}
	authOtpResp, _, err := c.c.AuthApi.V1AuthOtpSmsauthenticatePost(context.Background(), bsa, c.ratioClientId, c.ratioClientSecret)
	require.NoError(t, err)
	println(authOtpResp.SessionJwt)

}

func TestCrackJwt(t *testing.T) {
	testJwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6Imp3ay10ZXN0LWI5NWU0NTM5LTNlZGMtNGJiNS04MjdmLTViMjAyY2NmNDJlNiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsicHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyJdLCJleHAiOjE2ODc3MjA5MTcsImh0dHBzOi8vc3R5dGNoLmNvbS9zZXNzaW9uIjp7ImlkIjoic2Vzc2lvbi10ZXN0LTE3OGQ1NWVkLWY1ODYtNDljOC04NDQ1LTU3MGIzZDUxZDQ4MyIsInN0YXJ0ZWRfYXQiOiIyMDIzLTA2LTI1VDE5OjE2OjU3WiIsImxhc3RfYWNjZXNzZWRfYXQiOiIyMDIzLTA2LTI1VDE5OjE2OjU3WiIsImV4cGlyZXNfYXQiOiIyMDIzLTA2LTI1VDE5OjI2OjU3WiIsImF0dHJpYnV0ZXMiOnsidXNlcl9hZ2VudCI6IiIsImlwX2FkZHJlc3MiOiIifSwiYXV0aGVudGljYXRpb25fZmFjdG9ycyI6W3sidHlwZSI6Im90cCIsImRlbGl2ZXJ5X21ldGhvZCI6InNtcyIsImxhc3RfYXV0aGVudGljYXRlZF9hdCI6IjIwMjMtMDYtMjVUMTk6MTY6NTdaIiwicGhvbmVfbnVtYmVyX2ZhY3RvciI6eyJwaG9uZV9pZCI6InBob25lLW51bWJlci10ZXN0LTBmOGEwNmQyLWMxODQtNDViOC1hNGExLTVkMTVjNzc1YWMzZiIsInBob25lX251bWJlciI6IisxMjA1OTUyMjEyNCJ9fV19LCJpYXQiOjE2ODc3MjA2MTcsImlzcyI6InN0eXRjaC5jb20vcHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyIsIm5iZiI6MTY4NzcyMDYxNywic3ViIjoidXNlci10ZXN0LTc0OGY3OGI3LWZjYTQtNDk1ZC04N2M1LWIyMDhlYmQ5MzQzNyJ9.oGmPlFIRJdtPaeEBUGSmqOCW-paU7G0kBabzIdJCWcB930y7cNhREZO34OcT4mIRZC7lQrjUFT7Q-aj2qOmsp7HPQHEddj2a9PFK3j6jqfuYDnP2ajoZkF1HCcullV19tEnTUXpRFPy4Co-1Mgx2xYjBxOJg8rEtbj3VH05-2gBtdKNwbjOzzpXTtbvV1T7TvmoJzwl_7TYCotnLcpDNf1CDuqTA2q_W24GHSGHk_Q4IVH9Q02TgIz8-Z0RlWvgZ2ruoKsfqhj52R2Y5SbotiwYRK_pGF6VtZbiWYlDJAAcTLaxKBQfvgKANex5k_Kb3hJi3Zdl2JbIbU_KuQao8hw"

	//claims := jwt.MapClaims{}
	//jwt.ParseWithClaims()

	jwtSplits := strings.Split(testJwt, ".")

	jwtBody, err := base64.RawURLEncoding.DecodeString(jwtSplits[1])
	require.NoError(t, err)
	println(string(jwtBody))
}

func TestCreateUser(t *testing.T) {

	c := getDefaultClient("eyJhbGciOiJSUzI1NiIsImtpZCI6Imp3ay10ZXN0LWI5NWU0NTM5LTNlZGMtNGJiNS04MjdmLTViMjAyY2NmNDJlNiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsicHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyJdLCJleHAiOjE2ODc3MjA5MTcsImh0dHBzOi8vc3R5dGNoLmNvbS9zZXNzaW9uIjp7ImlkIjoic2Vzc2lvbi10ZXN0LTE3OGQ1NWVkLWY1ODYtNDljOC04NDQ1LTU3MGIzZDUxZDQ4MyIsInN0YXJ0ZWRfYXQiOiIyMDIzLTA2LTI1VDE5OjE2OjU3WiIsImxhc3RfYWNjZXNzZWRfYXQiOiIyMDIzLTA2LTI1VDE5OjE2OjU3WiIsImV4cGlyZXNfYXQiOiIyMDIzLTA2LTI1VDE5OjI2OjU3WiIsImF0dHJpYnV0ZXMiOnsidXNlcl9hZ2VudCI6IiIsImlwX2FkZHJlc3MiOiIifSwiYXV0aGVudGljYXRpb25fZmFjdG9ycyI6W3sidHlwZSI6Im90cCIsImRlbGl2ZXJ5X21ldGhvZCI6InNtcyIsImxhc3RfYXV0aGVudGljYXRlZF9hdCI6IjIwMjMtMDYtMjVUMTk6MTY6NTdaIiwicGhvbmVfbnVtYmVyX2ZhY3RvciI6eyJwaG9uZV9pZCI6InBob25lLW51bWJlci10ZXN0LTBmOGEwNmQyLWMxODQtNDViOC1hNGExLTVkMTVjNzc1YWMzZiIsInBob25lX251bWJlciI6IisxMjA1OTUyMjEyNCJ9fV19LCJpYXQiOjE2ODc3MjA2MTcsImlzcyI6InN0eXRjaC5jb20vcHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyIsIm5iZiI6MTY4NzcyMDYxNywic3ViIjoidXNlci10ZXN0LTc0OGY3OGI3LWZjYTQtNDk1ZC04N2M1LWIyMDhlYmQ5MzQzNyJ9.oGmPlFIRJdtPaeEBUGSmqOCW-paU7G0kBabzIdJCWcB930y7cNhREZO34OcT4mIRZC7lQrjUFT7Q-aj2qOmsp7HPQHEddj2a9PFK3j6jqfuYDnP2ajoZkF1HCcullV19tEnTUXpRFPy4Co-1Mgx2xYjBxOJg8rEtbj3VH05-2gBtdKNwbjOzzpXTtbvV1T7TvmoJzwl_7TYCotnLcpDNf1CDuqTA2q_W24GHSGHk_Q4IVH9Q02TgIz8-Z0RlWvgZ2ruoKsfqhj52R2Y5SbotiwYRK_pGF6VtZbiWYlDJAAcTLaxKBQfvgKANex5k_Kb3hJi3Zdl2JbIbU_KuQao8hw")

	cu := swagger.CreateUserRequest{
		FirstName:     "Paul",
		MiddleName:    "G",
		LastName:      "O'Really",
		Email:         "paul+12059522124@oleary.com",
		Country:       "USA",
		Phone:         "+12059522124",
		AcceptedTerms: true,
	}

	createUserResp, _, err := c.c.UserApi.V1UsersPost(context.Background(), cu, c.ratioClientId, c.ratioClientSecret)
	require.NoError(t, err)
	_ = createUserResp
}
