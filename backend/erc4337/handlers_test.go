package erc4337

import (
	"github.com/paulgoleary/local-luv-proto/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContext(t *testing.T) {
	mc, err := MakeContext(config.Config{ChainRpcUrl: "some_rule"})
	require.NoError(t, err)
	require.NotNil(t, mc.EntryPoint)
}
