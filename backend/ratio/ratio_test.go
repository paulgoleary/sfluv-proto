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

	c := getDefaultClient("")

	sk, err := crypto.SKFromInt(big.NewInt(6))
	require.NoError(t, err)

	addr := crypto.PubKeyToAddress(&sk.PublicKey)

	// TODO: 	WalletNetwork string `json:"walletNetwork"` - OpenAPI is old?
	b := swagger.AuthenticateCryptoWalletStartRequest{
		WalletAddress: addr.String(),
		WalletNetwork: "POLYGON",
	}
	println(addr.String())
	resp, httpResp, err := c.c.AuthApi.V1AuthCryptoWalletstartPost(context.Background(), b, c.ratioClientId, c.ratioClientSecret)
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
	authResp, httpResp, err := c.c.AuthApi.V1AuthCryptoWalletauthenticatePost(context.Background(), ba, c.ratioClientId, c.ratioClientSecret)
	require.NoError(t, err)
	println(authResp.SessionJwt)

	// need client with jwt auth now ...
	c = getDefaultClient(authResp.SessionJwt)
	// dc := getDefaultClient("")

	bs := swagger.SendSmsOtpRequest{
		PhoneNumber: "+15612023748",
	}
	sendOtpResp, _, err := c.c.AuthApi.V1AuthOtpSmssendPost(context.Background(), bs, c.ratioClientId, c.ratioClientSecret)
	require.NoError(t, err)
	println(sendOtpResp.PhoneId)
}

func TestSmsOtp(t *testing.T) {

	c := getDefaultClient("eyJhbGciOiJSUzI1NiIsImtpZCI6Imp3ay10ZXN0LWI5NWU0NTM5LTNlZGMtNGJiNS04MjdmLTViMjAyY2NmNDJlNiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsicHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyJdLCJleHAiOjE2ODg1MDE1OTgsImh0dHBzOi8vc3R5dGNoLmNvbS9zZXNzaW9uIjp7ImlkIjoic2Vzc2lvbi10ZXN0LTM4NzIyZTEzLTgyYzctNDJjMC05MmJjLTEwMDljNzI2NmQ4YSIsInN0YXJ0ZWRfYXQiOiIyMDIzLTA3LTA0VDIwOjA4OjE4WiIsImxhc3RfYWNjZXNzZWRfYXQiOiIyMDIzLTA3LTA0VDIwOjA4OjE4WiIsImV4cGlyZXNfYXQiOiIyMDIzLTA3LTA0VDIwOjE4OjE4WiIsImF0dHJpYnV0ZXMiOnsidXNlcl9hZ2VudCI6IiIsImlwX2FkZHJlc3MiOiIifSwiYXV0aGVudGljYXRpb25fZmFjdG9ycyI6W3sidHlwZSI6ImNyeXB0byIsImRlbGl2ZXJ5X21ldGhvZCI6ImNyeXB0b193YWxsZXQiLCJsYXN0X2F1dGhlbnRpY2F0ZWRfYXQiOiIyMDIzLTA3LTA0VDIwOjA4OjE4WiIsImNyeXB0b193YWxsZXRfZmFjdG9yIjp7ImNyeXB0b193YWxsZXRfaWQiOiJjcnlwdG8td2FsbGV0LXRlc3QtZTNiMTFiYWItOTEyNi00MGMxLWIyOGQtMDdkMTM5NTI1OWY5IiwiY3J5cHRvX3dhbGxldF9hZGRyZXNzIjoiMHhFNTdiRkU5RjQ0YjgxOTg5OEY0N0JGMzdFNUFGNzJhMDc4M2UxMTQxIiwiY3J5cHRvX3dhbGxldF90eXBlIjoiZXRoZXJldW0ifX1dfSwiaWF0IjoxNjg4NTAxMjk4LCJpc3MiOiJzdHl0Y2guY29tL3Byb2plY3QtdGVzdC1mOTk5ZDMzZS00OWRlLTQ4ODktYmVjNi04ZWIyODA0NGQwNDciLCJtb25pdG9yaW5nX3Nlc3Npb25faWQiOiIzMmQwMDU3My1hZDJlLTRjZWYtODIxZS1mYTMwYmUxY2ZmZmMiLCJuYmYiOjE2ODg1MDEyOTgsInN1YiI6InVzZXItdGVzdC0xMTdiNjBmNS05ZDc4LTQzNTQtODRjOS1mMjZkOGExZWQ5YmUifQ.hAlI3WWxEaIV9zPDZ__ZcpLZ7qWiqRLOhg4oWmuJyKNmFOaAMHD79VRzJ5m_pphvINwopC9fBqPK6_m6qkcN_5wteXk2AcpAOtOg1NPapnhP7kCnjNMvLOnrybU8gHmJzkiNQb__p1riGiRDxEAQagW0XlqNd4TMVTZdqOuVvcTSapUdps4eyfvdh9O3heky6r71O3B0Z9nFEKfnvOqRY8N22moptW9JJX4j3rtD2YImFNE1_YLeJvcC8DJ1D0XArvWd6TSuFbgXrXsMHQQrq4tb5on-bQKyOkMy34P5kcd3uhp_PB8DFdIZ_DV3UWdPOF82S3JFZiT7oewEzEtFoQ")
	// c := getDefaultClient("")

	bsa := swagger.AuthenticateSmsOtpRequest{
		Otp:     "586386",
		PhoneId: "phone-number-test-10807dd7-c1d9-42b6-9c19-bbca053a787f",
	}
	authOtpResp, _, err := c.c.AuthApi.V1AuthOtpSmsauthenticatePost(context.Background(), bsa, c.ratioClientId, c.ratioClientSecret)
	require.NoError(t, err)
	println(authOtpResp.SessionJwt)

}

func TestCrackJwt(t *testing.T) {
	testJwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6Imp3ay10ZXN0LWI5NWU0NTM5LTNlZGMtNGJiNS04MjdmLTViMjAyY2NmNDJlNiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsicHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyJdLCJleHAiOjE2ODg1MDE2MzUsImh0dHBzOi8vc3R5dGNoLmNvbS9zZXNzaW9uIjp7ImlkIjoic2Vzc2lvbi10ZXN0LTc4OTJlYjQyLTFiOGMtNDc2NC05NDc0LTFjNTA3MDZiNTMwYyIsInN0YXJ0ZWRfYXQiOiIyMDIzLTA3LTA0VDIwOjA4OjU1WiIsImxhc3RfYWNjZXNzZWRfYXQiOiIyMDIzLTA3LTA0VDIwOjA4OjU1WiIsImV4cGlyZXNfYXQiOiIyMDIzLTA3LTA0VDIwOjE4OjU1WiIsImF0dHJpYnV0ZXMiOnsidXNlcl9hZ2VudCI6IiIsImlwX2FkZHJlc3MiOiIifSwiYXV0aGVudGljYXRpb25fZmFjdG9ycyI6W3sidHlwZSI6ImNyeXB0byIsImRlbGl2ZXJ5X21ldGhvZCI6ImNyeXB0b193YWxsZXQiLCJsYXN0X2F1dGhlbnRpY2F0ZWRfYXQiOiIyMDIzLTA3LTA0VDIwOjA4OjE4WiIsImNyeXB0b193YWxsZXRfZmFjdG9yIjp7ImNyeXB0b193YWxsZXRfaWQiOiJjcnlwdG8td2FsbGV0LXRlc3QtZTNiMTFiYWItOTEyNi00MGMxLWIyOGQtMDdkMTM5NTI1OWY5IiwiY3J5cHRvX3dhbGxldF9hZGRyZXNzIjoiMHhFNTdiRkU5RjQ0YjgxOTg5OEY0N0JGMzdFNUFGNzJhMDc4M2UxMTQxIiwiY3J5cHRvX3dhbGxldF90eXBlIjoiZXRoZXJldW0ifX0seyJ0eXBlIjoib3RwIiwiZGVsaXZlcnlfbWV0aG9kIjoic21zIiwibGFzdF9hdXRoZW50aWNhdGVkX2F0IjoiMjAyMy0wNy0wNFQyMDowODo1NVoiLCJwaG9uZV9udW1iZXJfZmFjdG9yIjp7InBob25lX2lkIjoicGhvbmUtbnVtYmVyLXRlc3QtMTA4MDdkZDctYzFkOS00MmI2LTljMTktYmJjYTA1M2E3ODdmIiwicGhvbmVfbnVtYmVyIjoiKzE1NjEyMDIzNzQ4In19XX0sImlhdCI6MTY4ODUwMTMzNSwiaXNzIjoic3R5dGNoLmNvbS9wcm9qZWN0LXRlc3QtZjk5OWQzM2UtNDlkZS00ODg5LWJlYzYtOGViMjgwNDRkMDQ3IiwibW9uaXRvcmluZ19zZXNzaW9uX2lkIjoiMzJkMDA1NzMtYWQyZS00Y2VmLTgyMWUtZmEzMGJlMWNmZmZjIiwibmJmIjoxNjg4NTAxMzM1LCJzdWIiOiJ1c2VyLXRlc3QtMTE3YjYwZjUtOWQ3OC00MzU0LTg0YzktZjI2ZDhhMWVkOWJlIn0.cMizfE3iqyhul-4vkHwUjcm-JmbsJyFyLx-uLQEdKSqPv3uzASQa-aLdP23olFY0w5v2qWkq1gk-N2ZmTxLgpwMIOrgjln2Yy6E3WkbjCDWBewWq1CjyeJ8wJW-hxXbYPW4S9kIycFTxxyZHEYjyW5Lib-UBNG-gNYcdJw_tw4oIDz4Ykj_8aDa5MGTsI2ZNK8PN4CzlxoY-OuzqTUXEiOBaNQJykreoA4nUPsRCVoeBFn_D_MBBtNJpr-tqEc4CS-6w_A1ene56XCc_ZdWX4LBMzkHK2T2wie7aiylZJAoJp4594B8L-6DAWXAnq_NjDgeuCnc36_dnQYPvqm3_IA"

	jwtSplits := strings.Split(testJwt, ".")

	jwtBody, err := base64.RawURLEncoding.DecodeString(jwtSplits[1])
	require.NoError(t, err)
	println(string(jwtBody))
}

func TestCreateUser(t *testing.T) {

	c := getDefaultClient("eyJhbGciOiJSUzI1NiIsImtpZCI6Imp3ay10ZXN0LWI5NWU0NTM5LTNlZGMtNGJiNS04MjdmLTViMjAyY2NmNDJlNiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsicHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyJdLCJleHAiOjE2ODg1MDA2NTYsImh0dHBzOi8vc3R5dGNoLmNvbS9zZXNzaW9uIjp7ImlkIjoic2Vzc2lvbi10ZXN0LTIzNWY1NzdlLWNjYzItNDE2NC1iNmU4LWUzYjE1YjA2NTAxMiIsInN0YXJ0ZWRfYXQiOiIyMDIzLTA3LTA0VDE5OjUyOjM2WiIsImxhc3RfYWNjZXNzZWRfYXQiOiIyMDIzLTA3LTA0VDE5OjUyOjM2WiIsImV4cGlyZXNfYXQiOiIyMDIzLTA3LTA0VDIwOjAyOjM2WiIsImF0dHJpYnV0ZXMiOnsidXNlcl9hZ2VudCI6IiIsImlwX2FkZHJlc3MiOiIifSwiYXV0aGVudGljYXRpb25fZmFjdG9ycyI6W3sidHlwZSI6ImNyeXB0byIsImRlbGl2ZXJ5X21ldGhvZCI6ImNyeXB0b193YWxsZXQiLCJsYXN0X2F1dGhlbnRpY2F0ZWRfYXQiOiIyMDIzLTA3LTA0VDE5OjUyOjA1WiIsImNyeXB0b193YWxsZXRfZmFjdG9yIjp7ImNyeXB0b193YWxsZXRfaWQiOiJjcnlwdG8td2FsbGV0LXRlc3QtZTNiMTFiYWItOTEyNi00MGMxLWIyOGQtMDdkMTM5NTI1OWY5IiwiY3J5cHRvX3dhbGxldF9hZGRyZXNzIjoiMHhFNTdiRkU5RjQ0YjgxOTg5OEY0N0JGMzdFNUFGNzJhMDc4M2UxMTQxIiwiY3J5cHRvX3dhbGxldF90eXBlIjoiZXRoZXJldW0ifX0seyJ0eXBlIjoib3RwIiwiZGVsaXZlcnlfbWV0aG9kIjoic21zIiwibGFzdF9hdXRoZW50aWNhdGVkX2F0IjoiMjAyMy0wNy0wNFQxOTo1MjozNloiLCJwaG9uZV9udW1iZXJfZmFjdG9yIjp7InBob25lX2lkIjoicGhvbmUtbnVtYmVyLXRlc3QtMTA4MDdkZDctYzFkOS00MmI2LTljMTktYmJjYTA1M2E3ODdmIiwicGhvbmVfbnVtYmVyIjoiKzE1NjEyMDIzNzQ4In19XX0sImlhdCI6MTY4ODUwMDM1NiwiaXNzIjoic3R5dGNoLmNvbS9wcm9qZWN0LXRlc3QtZjk5OWQzM2UtNDlkZS00ODg5LWJlYzYtOGViMjgwNDRkMDQ3IiwibW9uaXRvcmluZ19zZXNzaW9uX2lkIjoiNjQxOTY0MzQtOGUwYS00NTc1LWJhYWEtYTg3MTc1ZDZjYThiIiwibmJmIjoxNjg4NTAwMzU2LCJzdWIiOiJ1c2VyLXRlc3QtMTE3YjYwZjUtOWQ3OC00MzU0LTg0YzktZjI2ZDhhMWVkOWJlIn0.t7SXVKId_qvhI4t366aDreiCSHRSeMDESvVWI8cYKBL3ymLJS0YDJqUAyeBItWGmPuSW7-n9jPO7YmqsSOYu3zoT-Ns-ZIAbGfkJrz-Zbs5juK22l-K5ZquMeJKcN_pbVaYMBAhSE_ELFNRUxjkm48wLeBAHavyUYZopGI_aoGAAq4S2PElOQtr1Ojx3SNdATsQgRPPOkO9HiH9DzY4P_4oNcOOorV7v_Z8qH-mdkGkm75p3-rv9UNKwOB4kBXMrjp-ZlguI9zpWOGUrkpMIFt769TFRFQUonF9DELx4GS9S4PXQ381l5FWepdJvStgBukUgLKuKIdbjzyfROUXXqw")

	maybeUser, _, err := c.c.UserApi.V1UsersUserIdGet(context.Background(), "b6fb3c76-dac7-490f-9a60-0ece43c681bb", c.ratioClientId, c.ratioClientSecret)
	if swagErr, ok := err.(swagger.GenericSwaggerError); ok {
		println(string(swagErr.Body()))
	}
	if err == nil {
		return
	}
	_ = maybeUser

	cu := swagger.CreateUserRequest{
		FirstName:     "Paul",
		MiddleName:    "G",
		LastName:      "OReally",
		Email:         "paul+15612023748@oleary.com",
		Country:       "US",
		Phone:         "+15612023748",
		AcceptedTerms: true,
	}

	createUserResp, _, err := c.c.UserApi.V1UsersPost(context.Background(), cu, c.ratioClientId, c.ratioClientSecret)
	if swagErr, ok := err.(swagger.GenericSwaggerError); ok {
		println(string(swagErr.Body()))
	}
	require.NoError(t, err)
	println(createUserResp.Id)
}
