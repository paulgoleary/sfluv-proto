package ratio

import (
	"github.com/gin-gonic/gin/render"
	swagger "github.com/paulgoleary/local-luv-proto/ratio/go-client-generated"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestAuthResultRender(t *testing.T) {

	u := swagger.User{
		Id:                    "someUserId",
		CreateTime:            "someTime",
		UpdateTime:            "someOtherTime",
		FirstName:             "someName",
		LastName:              "someOtherName",
		Email:                 "someEmail",
		Country:               "someCountry",
		Phone:                 "somePhone",
		Kyc:                   nil,
		ConnectedBankAccounts: nil,
		Flags:                 nil,
	}
	testRes := fullAuthUserResult{
		Jwt:    "someJwt",
		UserId: "someUserId",
		User:   &u,
	}

	r := render.JSON{Data: testRes}
	rc := httptest.NewRecorder()
	err := r.Render(rc)
	require.NoError(t, err)
	require.True(t, len(rc.Body.Bytes()) > 0)
}
