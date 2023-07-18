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

func ethSign(msg string, privateKey *ecdsa.PrivateKey) (sig []byte, err error) {
	message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	hash := ethgo.Keccak256([]byte(message))

	if sig, err = crypto.Sign(privateKey, hash); err != nil {
		return
	}
	sig[64] += 27
	return
}

type testProfile struct {
	sk          *ecdsa.PrivateKey
	phone       string
	phoneId     string
	maybeUserId string
}

var profileWithUser testProfile
var profileNoUser testProfile

var tp *testProfile

const phoneSK6 = "+15612023748"

func init() {
	sk6, _ := crypto.SKFromInt(big.NewInt(6))
	profileWithUser = testProfile{
		sk:          sk6,
		phone:       phoneSK6,
		phoneId:     "phone-number-test-10807dd7-c1d9-42b6-9c19-bbca053a787f",
		maybeUserId: "b6fb3c76-dac7-490f-9a60-0ece43c681bb",
	}

	ski := new(big.Int)
	ski, _ = ski.SetString("93776917413213983275849089862893824682941322802105134120772412085212056855608", 10)
	sk, _ := crypto.SKFromInt(ski)
	profileNoUser = testProfile{
		sk:    sk,
		phone: "+15799777923",
	}

	tp = &profileWithUser
}

func TestRatioAPI(t *testing.T) {

	c := getDefaultClient("")

	addr := crypto.PubKeyToAddress(&tp.sk.PublicKey)

	// TODO: 	WalletNetwork string `json:"walletNetwork"` - OpenAPI is old?
	b := swagger.AuthenticateCryptoWalletStartRequest{
		WalletAddress: addr.String(),
		WalletNetwork: "POLYGON",
	}
	println(addr.String())
	resp, httpResp, err := c.c.AuthApi.V1AuthCryptoWalletstartPost(context.Background(), b, c.ratioClientId, c.ratioClientSecret)
	if swagErr, ok := err.(swagger.GenericSwaggerError); ok {
		println(string(swagErr.Body()))
	}
	require.NoError(t, err)
	_ = httpResp

	println(resp.Challenge)
	sig, err := ethSign(resp.Challenge, tp.sk)
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

	if authResp.UserMask != nil {
		user, _, err := c.c.UserApi.V1UsersUserIdGet(context.Background(), authResp.UserMask.Id, c.ratioClientId, c.ratioClientSecret)
		require.NoError(t, err)
		require.Equal(t, tp.phone, user.Phone)
	}

	bs := swagger.SendSmsOtpRequest{
		PhoneNumber: tp.phone,
	}
	sendOtpResp, _, err := c.c.AuthApi.V1AuthOtpSmssendPost(context.Background(), bs, c.ratioClientId, c.ratioClientSecret)
	if swagErr, ok := err.(swagger.GenericSwaggerError); ok {
		println(string(swagErr.Body()))
	}
	require.NoError(t, err)
	println(sendOtpResp.PhoneId)
}

func TestSmsOtp(t *testing.T) {

	// assumes wallet auth flow has been run - i.e. current jwt
	// c := getDefaultClient(walletAuthJwt)
	c := getDefaultClient("eyJhbGciOiJSUzI1NiIsImtpZCI6Imp3ay10ZXN0LWI5NWU0NTM5LTNlZGMtNGJiNS04MjdmLTViMjAyY2NmNDJlNiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsicHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyJdLCJleHAiOjE2ODk2NDIwMDIsImh0dHBzOi8vc3R5dGNoLmNvbS9zZXNzaW9uIjp7ImlkIjoic2Vzc2lvbi10ZXN0LTc2YmJlODg1LTE3OWMtNDZhNS1iZDNjLTM0NDJiZTI1NjUyNCIsInN0YXJ0ZWRfYXQiOiIyMDIzLTA3LTE4VDAwOjU1OjAyWiIsImxhc3RfYWNjZXNzZWRfYXQiOiIyMDIzLTA3LTE4VDAwOjU1OjAyWiIsImV4cGlyZXNfYXQiOiIyMDIzLTA3LTE4VDAxOjA1OjAyWiIsImF0dHJpYnV0ZXMiOnsidXNlcl9hZ2VudCI6IiIsImlwX2FkZHJlc3MiOiIifSwiYXV0aGVudGljYXRpb25fZmFjdG9ycyI6W3sidHlwZSI6ImNyeXB0byIsImRlbGl2ZXJ5X21ldGhvZCI6ImNyeXB0b193YWxsZXQiLCJsYXN0X2F1dGhlbnRpY2F0ZWRfYXQiOiIyMDIzLTA3LTE4VDAwOjU1OjAyWiIsImNyeXB0b193YWxsZXRfZmFjdG9yIjp7ImNyeXB0b193YWxsZXRfaWQiOiJjcnlwdG8td2FsbGV0LXRlc3QtZTNiMTFiYWItOTEyNi00MGMxLWIyOGQtMDdkMTM5NTI1OWY5IiwiY3J5cHRvX3dhbGxldF9hZGRyZXNzIjoiMHhFNTdiRkU5RjQ0YjgxOTg5OEY0N0JGMzdFNUFGNzJhMDc4M2UxMTQxIiwiY3J5cHRvX3dhbGxldF90eXBlIjoiZXRoZXJldW0ifX1dfSwiaWF0IjoxNjg5NjQxNzAyLCJpc3MiOiJzdHl0Y2guY29tL3Byb2plY3QtdGVzdC1mOTk5ZDMzZS00OWRlLTQ4ODktYmVjNi04ZWIyODA0NGQwNDciLCJtb25pdG9yaW5nX3Nlc3Npb25faWQiOiI5MzE1ZTg3ZS00OTM3LTRhMWUtYjk3MC0yOTZjNjU4NjJmYjYiLCJuYmYiOjE2ODk2NDE3MDIsInN1YiI6InVzZXItdGVzdC0xMTdiNjBmNS05ZDc4LTQzNTQtODRjOS1mMjZkOGExZWQ5YmUifQ.S4TFblRPNU7LAhWOWI6yoefVmBYhn_ht0ofpXvcEuB9EVQewjtcetfBZ-1U6giN85SxQEzFoUwEoyN4khxOXoRIyprKFMayeVrV1iuIJrH7B3TgTtln1iev4wKpBcWo6FBqr9CT4IomY6md_KvV9d1ZCURmULeC48GLlenjLABRufwKCWdfp5vpYkXbFz7E9CQZIQNTJ5_KGd__D2JtrtiTBolrECBbDMlwEAPgZ737ir33bYh_FFsWkC74yfL_jdBPQ7uaElX3r8CDrvqOxx-o3xdOqFNsr5ScnL8AzRwuRPLfZmbBbY6053pgTqaf7Nqu2vxzkFgt2S5SfG_snlg")

	bsa := swagger.AuthenticateSmsOtpRequest{
		Otp:     "446771",
		PhoneId: tp.phoneId,
	}
	authOtpResp, _, err := c.c.AuthApi.V1AuthOtpSmsauthenticatePost(context.Background(), bsa, c.ratioClientId, c.ratioClientSecret)
	if swagErr, ok := err.(swagger.GenericSwaggerError); ok {
		println(string(swagErr.Body()))
	}
	require.NoError(t, err)
	println(authOtpResp.SessionJwt)

}

func TestGetUser(t *testing.T) {
	c := getDefaultClient("")
	maybeUser, _, err := c.c.UserApi.V1UsersUserIdGet(context.Background(), tp.maybeUserId, c.ratioClientId, c.ratioClientSecret)
	if swagErr, ok := err.(swagger.GenericSwaggerError); ok {
		println(string(swagErr.Body()))
	}
	if err == nil {
		userWallets, _, err := c.c.WalletApi.V1UsersUserIdWalletsGet(context.Background(), maybeUser.Id, c.ratioClientId, c.ratioClientSecret)
		require.NoError(t, err)
		if len(userWallets.Items) == 0 {
			b := swagger.ConnectWalletRequest{
				Address: crypto.PubKeyToAddress(&tp.sk.PublicKey).String(),
				Type_:   "POLYGON",
				Name:    "SFLUV Default Wallet",
			}
			wallet, _, err := c.c.WalletApi.V1UsersUserIdWalletsPost(context.Background(), b, c.ratioClientId, c.ratioClientSecret, maybeUser.Id)
			require.NoError(t, err)
			_ = wallet
		}
	}
}

func TestCreateUser(t *testing.T) {

	c := getDefaultClient("eyJhbGciOiJSUzI1NiIsImtpZCI6Imp3ay10ZXN0LWI5NWU0NTM5LTNlZGMtNGJiNS04MjdmLTViMjAyY2NmNDJlNiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsicHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyJdLCJleHAiOjE2ODkyNzU0ODAsImh0dHBzOi8vc3R5dGNoLmNvbS9zZXNzaW9uIjp7ImlkIjoic2Vzc2lvbi10ZXN0LTQ2YjI1MzIyLWNlZmItNDY5NS1hY2U0LWE4ZTAwMWZmMmU0OSIsInN0YXJ0ZWRfYXQiOiIyMDIzLTA3LTEzVDE5OjA2OjIwWiIsImxhc3RfYWNjZXNzZWRfYXQiOiIyMDIzLTA3LTEzVDE5OjA2OjIwWiIsImV4cGlyZXNfYXQiOiIyMDIzLTA3LTEzVDE5OjE2OjIwWiIsImF0dHJpYnV0ZXMiOnsidXNlcl9hZ2VudCI6IiIsImlwX2FkZHJlc3MiOiIifSwiYXV0aGVudGljYXRpb25fZmFjdG9ycyI6W3sidHlwZSI6ImNyeXB0byIsImRlbGl2ZXJ5X21ldGhvZCI6ImNyeXB0b193YWxsZXQiLCJsYXN0X2F1dGhlbnRpY2F0ZWRfYXQiOiIyMDIzLTA3LTEzVDE5OjA1OjI1WiIsImNyeXB0b193YWxsZXRfZmFjdG9yIjp7ImNyeXB0b193YWxsZXRfaWQiOiJjcnlwdG8td2FsbGV0LXRlc3QtZTNiMTFiYWItOTEyNi00MGMxLWIyOGQtMDdkMTM5NTI1OWY5IiwiY3J5cHRvX3dhbGxldF9hZGRyZXNzIjoiMHhFNTdiRkU5RjQ0YjgxOTg5OEY0N0JGMzdFNUFGNzJhMDc4M2UxMTQxIiwiY3J5cHRvX3dhbGxldF90eXBlIjoiZXRoZXJldW0ifX0seyJ0eXBlIjoib3RwIiwiZGVsaXZlcnlfbWV0aG9kIjoic21zIiwibGFzdF9hdXRoZW50aWNhdGVkX2F0IjoiMjAyMy0wNy0xM1QxOTowNjoyMFoiLCJwaG9uZV9udW1iZXJfZmFjdG9yIjp7InBob25lX2lkIjoicGhvbmUtbnVtYmVyLXRlc3QtMTA4MDdkZDctYzFkOS00MmI2LTljMTktYmJjYTA1M2E3ODdmIiwicGhvbmVfbnVtYmVyIjoiKzE1NjEyMDIzNzQ4In19XX0sImlhdCI6MTY4OTI3NTE4MCwiaXNzIjoic3R5dGNoLmNvbS9wcm9qZWN0LXRlc3QtZjk5OWQzM2UtNDlkZS00ODg5LWJlYzYtOGViMjgwNDRkMDQ3IiwibW9uaXRvcmluZ19zZXNzaW9uX2lkIjoiNzVjMmEwMTYtYjMxMy00ZjIxLTk5OTMtOTliMmQ4YzYyMGZmIiwibmJmIjoxNjg5Mjc1MTgwLCJzdWIiOiJ1c2VyLXRlc3QtMTE3YjYwZjUtOWQ3OC00MzU0LTg0YzktZjI2ZDhhMWVkOWJlIn0.dG_P4VoWVVLgwrR_HQhSEzCYkxMTDRpB47spdvAyirx1K5EE7PGh0BZn0nIiDIgExFcOLdiedoP7sQoA08qRKPxvNGUj1SMGedOeOPnQxju6qa6zKuB4uNF-UIUn73-ZNuaNZRysPnk6Gp9e7aDm4w5bM2CISfvuDGby1s7ZkscYjARykUkB5fS36jH1OGC60hddID_xI_y8W4R_5wbcq3bObCpA8kKYQOIPY9LcJFx7AlRAzRIIYH8OGOIeyuJj2TLWRogct2LQSNu_OqLVqWjQ2jGsgIEAoE4I0vrJxFMc8t4F4uS1RHmXap5halQGXwB2-r6QmskbnA18v61r_A")

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

func TestCrackJwt(t *testing.T) {
	testJwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6Imp3ay10ZXN0LWI5NWU0NTM5LTNlZGMtNGJiNS04MjdmLTViMjAyY2NmNDJlNiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsicHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyJdLCJleHAiOjE2ODkyNzU0ODAsImh0dHBzOi8vc3R5dGNoLmNvbS9zZXNzaW9uIjp7ImlkIjoic2Vzc2lvbi10ZXN0LTQ2YjI1MzIyLWNlZmItNDY5NS1hY2U0LWE4ZTAwMWZmMmU0OSIsInN0YXJ0ZWRfYXQiOiIyMDIzLTA3LTEzVDE5OjA2OjIwWiIsImxhc3RfYWNjZXNzZWRfYXQiOiIyMDIzLTA3LTEzVDE5OjA2OjIwWiIsImV4cGlyZXNfYXQiOiIyMDIzLTA3LTEzVDE5OjE2OjIwWiIsImF0dHJpYnV0ZXMiOnsidXNlcl9hZ2VudCI6IiIsImlwX2FkZHJlc3MiOiIifSwiYXV0aGVudGljYXRpb25fZmFjdG9ycyI6W3sidHlwZSI6ImNyeXB0byIsImRlbGl2ZXJ5X21ldGhvZCI6ImNyeXB0b193YWxsZXQiLCJsYXN0X2F1dGhlbnRpY2F0ZWRfYXQiOiIyMDIzLTA3LTEzVDE5OjA1OjI1WiIsImNyeXB0b193YWxsZXRfZmFjdG9yIjp7ImNyeXB0b193YWxsZXRfaWQiOiJjcnlwdG8td2FsbGV0LXRlc3QtZTNiMTFiYWItOTEyNi00MGMxLWIyOGQtMDdkMTM5NTI1OWY5IiwiY3J5cHRvX3dhbGxldF9hZGRyZXNzIjoiMHhFNTdiRkU5RjQ0YjgxOTg5OEY0N0JGMzdFNUFGNzJhMDc4M2UxMTQxIiwiY3J5cHRvX3dhbGxldF90eXBlIjoiZXRoZXJldW0ifX0seyJ0eXBlIjoib3RwIiwiZGVsaXZlcnlfbWV0aG9kIjoic21zIiwibGFzdF9hdXRoZW50aWNhdGVkX2F0IjoiMjAyMy0wNy0xM1QxOTowNjoyMFoiLCJwaG9uZV9udW1iZXJfZmFjdG9yIjp7InBob25lX2lkIjoicGhvbmUtbnVtYmVyLXRlc3QtMTA4MDdkZDctYzFkOS00MmI2LTljMTktYmJjYTA1M2E3ODdmIiwicGhvbmVfbnVtYmVyIjoiKzE1NjEyMDIzNzQ4In19XX0sImlhdCI6MTY4OTI3NTE4MCwiaXNzIjoic3R5dGNoLmNvbS9wcm9qZWN0LXRlc3QtZjk5OWQzM2UtNDlkZS00ODg5LWJlYzYtOGViMjgwNDRkMDQ3IiwibW9uaXRvcmluZ19zZXNzaW9uX2lkIjoiNzVjMmEwMTYtYjMxMy00ZjIxLTk5OTMtOTliMmQ4YzYyMGZmIiwibmJmIjoxNjg5Mjc1MTgwLCJzdWIiOiJ1c2VyLXRlc3QtMTE3YjYwZjUtOWQ3OC00MzU0LTg0YzktZjI2ZDhhMWVkOWJlIn0.dG_P4VoWVVLgwrR_HQhSEzCYkxMTDRpB47spdvAyirx1K5EE7PGh0BZn0nIiDIgExFcOLdiedoP7sQoA08qRKPxvNGUj1SMGedOeOPnQxju6qa6zKuB4uNF-UIUn73-ZNuaNZRysPnk6Gp9e7aDm4w5bM2CISfvuDGby1s7ZkscYjARykUkB5fS36jH1OGC60hddID_xI_y8W4R_5wbcq3bObCpA8kKYQOIPY9LcJFx7AlRAzRIIYH8OGOIeyuJj2TLWRogct2LQSNu_OqLVqWjQ2jGsgIEAoE4I0vrJxFMc8t4F4uS1RHmXap5halQGXwB2-r6QmskbnA18v61r_A"

	jwtSplits := strings.Split(testJwt, ".")

	jwtBody, err := base64.RawURLEncoding.DecodeString(jwtSplits[1])
	require.NoError(t, err)
	println(string(jwtBody))
}

func TestUserKYC(t *testing.T) {
	c := getDefaultClient("eyJhbGciOiJSUzI1NiIsImtpZCI6Imp3ay10ZXN0LWI5NWU0NTM5LTNlZGMtNGJiNS04MjdmLTViMjAyY2NmNDJlNiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsicHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyJdLCJleHAiOjE2ODk2NDIwMzksImh0dHBzOi8vc3R5dGNoLmNvbS9zZXNzaW9uIjp7ImlkIjoic2Vzc2lvbi10ZXN0LWQ3ZTQzOGZhLTc5YjAtNDZjZi1hNmViLTc4ZjdhMTcwOThjNiIsInN0YXJ0ZWRfYXQiOiIyMDIzLTA3LTE4VDAwOjU1OjM5WiIsImxhc3RfYWNjZXNzZWRfYXQiOiIyMDIzLTA3LTE4VDAwOjU1OjM5WiIsImV4cGlyZXNfYXQiOiIyMDIzLTA3LTE4VDAxOjA1OjM5WiIsImF0dHJpYnV0ZXMiOnsidXNlcl9hZ2VudCI6IiIsImlwX2FkZHJlc3MiOiIifSwiYXV0aGVudGljYXRpb25fZmFjdG9ycyI6W3sidHlwZSI6ImNyeXB0byIsImRlbGl2ZXJ5X21ldGhvZCI6ImNyeXB0b193YWxsZXQiLCJsYXN0X2F1dGhlbnRpY2F0ZWRfYXQiOiIyMDIzLTA3LTE4VDAwOjU1OjAyWiIsImNyeXB0b193YWxsZXRfZmFjdG9yIjp7ImNyeXB0b193YWxsZXRfaWQiOiJjcnlwdG8td2FsbGV0LXRlc3QtZTNiMTFiYWItOTEyNi00MGMxLWIyOGQtMDdkMTM5NTI1OWY5IiwiY3J5cHRvX3dhbGxldF9hZGRyZXNzIjoiMHhFNTdiRkU5RjQ0YjgxOTg5OEY0N0JGMzdFNUFGNzJhMDc4M2UxMTQxIiwiY3J5cHRvX3dhbGxldF90eXBlIjoiZXRoZXJldW0ifX0seyJ0eXBlIjoib3RwIiwiZGVsaXZlcnlfbWV0aG9kIjoic21zIiwibGFzdF9hdXRoZW50aWNhdGVkX2F0IjoiMjAyMy0wNy0xOFQwMDo1NTozOVoiLCJwaG9uZV9udW1iZXJfZmFjdG9yIjp7InBob25lX2lkIjoicGhvbmUtbnVtYmVyLXRlc3QtMTA4MDdkZDctYzFkOS00MmI2LTljMTktYmJjYTA1M2E3ODdmIiwicGhvbmVfbnVtYmVyIjoiKzE1NjEyMDIzNzQ4In19XX0sImlhdCI6MTY4OTY0MTczOSwiaXNzIjoic3R5dGNoLmNvbS9wcm9qZWN0LXRlc3QtZjk5OWQzM2UtNDlkZS00ODg5LWJlYzYtOGViMjgwNDRkMDQ3IiwibW9uaXRvcmluZ19zZXNzaW9uX2lkIjoiOTMxNWU4N2UtNDkzNy00YTFlLWI5NzAtMjk2YzY1ODYyZmI2IiwibmJmIjoxNjg5NjQxNzM5LCJzdWIiOiJ1c2VyLXRlc3QtMTE3YjYwZjUtOWQ3OC00MzU0LTg0YzktZjI2ZDhhMWVkOWJlIn0.K8V8j919G7fnVQ3yaGa0IMp-lh78oa5xqcH8MdOy6Cemi-QqyTSdSsPU01BkgxCH6Fm9G8B_oyoMFI9uEyF4HqgvWa2ADbyiOtIwLcwfmoAXE9wHJMm7TK_9XrYDOd-9MMiOpODSt8VN50RqO4IeWk9JBcRJhc6RFcXW0O_RhvIV-pdsMAI3z3XbgIkdATKG7tAD5z_Ar4-pVDEST6sxFHb4TG2w2q6Vo_BmcC0ysBe0XMTPkv0XsvN040KjFh2j0626ETXMPPsZ5Zju3aVdygj4XpVLDd0A0uMawRfiX_13y8cADCh8dLBqjB_Th8MU1U3rB6m4RJlUO5COYjvWQQ")

	maybeUser, _, err := c.c.UserApi.V1UsersUserIdGet(context.Background(), tp.maybeUserId, c.ratioClientId, c.ratioClientSecret)
	if swagErr, ok := err.(swagger.GenericSwaggerError); ok {
		println(string(swagErr.Body()))
	}
	_ = maybeUser

	ba := swagger.SubmitKycRequest{
		DateOfBirth: "1968-01-01",
		IdType:      "SSN",
		IdNumber:    "587947847",
		Line1:       "1212 1st st.",
		Line2:       "",
		City:        "San Francisco",
		State:       "CA",
		PostalCode:  "94131",
	}
	user, _, err := c.c.UserApi.V1UsersUserIdKycPost(context.Background(), ba, c.ratioClientId, c.ratioClientSecret, tp.maybeUserId)
	if swagErr, ok := err.(swagger.GenericSwaggerError); ok {
		println(string(swagErr.Body()))
	}
	require.NoError(t, err)
	_ = user
}
