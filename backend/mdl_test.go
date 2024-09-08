package local_luv_proto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

func TestDecodeFailure(t *testing.T) {
	testDec := "eyJjcnYiOiJQLTI1NiIsImt0eSI6IkVDIiwieCI6IjN0dVF5YWsxYXhSTGN3Q1hzOXl0U1NDNGt4QzQtUGFlWmhwUXVzQWxJM28iLCJ4NWMiOlsiTUlJQ2VEQ0NBaCtnQXdJQkFnSVVOYjF2czlucklGTmt4Vy9BaDVSYnFTazVENDh3Q2dZSUtvWkl6ajBFQXdJd1VURUxNQWtHQTFVRUJoTUNWVk14RGpBTUJnTlZCQWdNQlZWVExVTkJNUTh3RFFZRFZRUUtEQVpEUVMxRVRWWXhJVEFmQmdOVkJBTU1HRU5oYkdsbWIzSnVhV0VnUkUxV0lFbEJRMEVnVW05dmREQWVGdzB5TkRBeE1qWXlNRE01TXpsYUZ3MHlOVEF4TWpVeU1ETTVNemxhTUZZeEN6QUpCZ05WQkFZVEFsVlRNUTR3REFZRFZRUUlEQVZWVXkxRFFURVBNQTBHQTFVRUNnd0dRMEV0UkUxV01TWXdKQVlEVlFRRERCMURZV3hwWm05eWJtbGhJRVJOVmlCSlFVTkJJRlpESUZOcFoyNWxjakJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUFCTjdia01tcE5Xc1VTM01BbDdQY3JVa2d1Sk1RdVBqMm5tWWFVTHJBSlNONm80NlRZaXYvaG4wM0N0a1ZEaElPTDNMb3Q4dFZZV0JxTlBtWmZpOTZQVktqZ2M4d2djd3dIUVlEVlIwT0JCWUVGUGp0bVgwQjYwbVFRQmxSVGErblVNSEFIYjdFTUI4R0ExVWRJd1FZTUJhQUZMdDlkV2VTZW0vUG4zSjd1QXIzTnkrY0RGQTJNQjBHQ1dDR1NBR0crRUlCRFFRUUZnNURZV3hwWm05eWJtbGhJRVJOVmpBT0JnTlZIUThCQWY4RUJBTUNCNEF3SVFZRFZSMFNCQm93R0lFV2FXRmpZUzF6YVdkdVpYSkFaRzEyTG1OaExtZHZkakE0QmdOVkhSOEVNVEF2TUMyZ0s2QXBoaWRvZEhSd2N6b3ZMMk55YkM1a2JYWXVZMkV1WjI5MkwybGhZMkV2YldSdll5MXphV2R1WlhJd0NnWUlLb1pJemowRUF3SURSd0F3UkFJZ01zWk10Mis2OTJkOG8zbmROU3lMV0djR2hLY05hRFZGZ3poNk5WWmV6OElDSUFwVUVJaUdJZmxrSUFBa3V4Z2VwaWJLQVlUOCtaNEI4dTIwUXVzbC9QSC8iXSwieSI6Im80NlRZaXZfaG4wM0N0a1ZEaElPTDNMb3Q4dFZZV0JxTlBtWmZpOTZQVkkifQ"
	testLen := len(testDec)
	_ = testLen

	_, err := base64.RawURLEncoding.DecodeString(testDec)
	require.NoError(t, err)
}

func extractJWKPublicKey(kid string) (pk *ecdsa.PublicKey, err error) {
	kid = strings.TrimPrefix(kid, "did:jwk:")
	kid = strings.TrimSuffix(kid, "#0")

	var b []byte
	if b, err = base64.RawURLEncoding.DecodeString(kid); err != nil {
		return
	}
	var ks jwk.Set
	if ks, err = jwk.Parse(b); err != nil {
		return
	}
	// TODO: more robust
	maybeKey, _ := ks.Key(0)
	maybeECKey := maybeKey.(jwk.ECDSAPublicKey)
	pX := new(big.Int)
	pX.SetBytes(maybeECKey.X())
	pY := new(big.Int)
	pY.SetBytes(maybeECKey.Y())
	pk = &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     pX,
		Y:     pY,
	}

	return
}

// TODO: more robust
func extractEncVC(mapVP map[string]any) string {
	vcs := mapVP["verifiableCredential"].([]any)
	if len(vcs) > 0 {
		return vcs[0].(string)
	}
	return ""
}

func TestVPToken(t *testing.T) {

	encVPToken := os.Getenv("EncVPToken")

	kf := func(tok *jwt.Token) (key interface{}, err error) {
		if _, ok := tok.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tok.Header["alg"])
		} else {
			if kid, ok := tok.Header["kid"].(string); !ok {
				return nil, fmt.Errorf("expected 'kid' string in header")
			} else {
				key, err = extractJWKPublicKey(kid)
			}
		}
		return
	}
	p := new(jwt.Parser)
	p.SkipClaimsValidation = true
	vpTok, err := p.Parse(encVPToken, kf)
	require.NoError(t, err)

	mc := vpTok.Claims.(jwt.MapClaims)
	maybeVp := mc["vp"].(map[string]any)

	encVC := extractEncVC(maybeVp)

	vcTok, err := p.Parse(encVC, kf)
	require.NoError(t, err)

	_ = vcTok

}
