package erc4337

import (
	"github.com/gin-gonic/gin"
	"github.com/paulgoleary/local-luv-proto/config"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicContext(t *testing.T) {
	mc, err := MakeContext(config.Config{ChainRpcUrl: "some_url"})
	require.NoError(t, err)
	require.NotNil(t, mc.EntryPoint)
}

func TestContextUserOp(t *testing.T) {

	mc, err := makeTestContext(map[string]string{
		"nonce":  "0",
		"sender": "deadbeef",
	})
	require.NoError(t, err)

	w := httptest.NewRecorder()
	engine := gin.New()
	testContext := gin.CreateTestContextOnly(w, engine)

	testContext.Request, _ = http.NewRequest(http.MethodGet, "/doesntmatter?owner=6D64a4aF99563a82B212124604f6d1759376F37F", nil)
	mc.HandleGetSender(testContext)

	println(w.Body.String())

}
