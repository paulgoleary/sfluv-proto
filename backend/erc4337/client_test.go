package erc4337

import (
	"context"
	"github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/paulgoleary/local-luv-proto/chain"
	"github.com/stackup-wallet/stackup-bundler/pkg/client"
	"github.com/stackup-wallet/stackup-bundler/pkg/gas"
	"github.com/stackup-wallet/stackup-bundler/pkg/mempool"
	"github.com/stackup-wallet/stackup-bundler/pkg/modules/checks"
	"github.com/stackup-wallet/stackup-bundler/pkg/modules/paymaster"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"math/big"
	"os"
	"testing"
)

func TestStackupClient(t *testing.T) {

	conf := struct {
		SupportedEntryPoints    []common.Address
		EthClientUrl            string
		MaxBatchGasLimit        *big.Int
		MaxVerificationGas      *big.Int
		MaxOpsForUnstakedSender int
	}{
		SupportedEntryPoints:    []common.Address{common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789")},
		EthClientUrl:            os.Getenv("NETWORK_URL"),
		MaxBatchGasLimit:        big.NewInt(10_000_000), // TODO: value?
		MaxVerificationGas:      big.NewInt(10_000_000), // TODO: value?
		MaxOpsForUnstakedSender: 1,
	}

	rpc, err := rpc.Dial(conf.EthClientUrl)
	require.NoError(t, err)

	eth := ethclient.NewClient(rpc)
	chainId, err := eth.ChainID(context.Background())
	require.NoError(t, err)

	bo := badger.DefaultOptions("")
	bo.InMemory = true
	bdb, err := badger.Open(bo)
	require.NoError(t, err)

	mp, err := mempool.New(bdb)
	require.NoError(t, err)

	ov := gas.NewDefaultOverhead()

	check := checks.New(
		bdb,
		rpc,
		ov,
		conf.MaxVerificationGas,
		conf.MaxBatchGasLimit,
		conf.MaxOpsForUnstakedSender,
	)

	paymaster := paymaster.New(bdb)

	// init client - same as StackUp private
	c := client.New(mp, ov, chainId, conf.SupportedEntryPoints)
	c.SetGetUserOpReceiptFunc(client.GetUserOpReceiptWithEthClient(eth))
	c.SetGetGasEstimateFunc(client.GetGasEstimateWithEthClient(rpc, ov, chainId, conf.MaxBatchGasLimit))
	c.SetGetUserOpByHashFunc(client.GetUserOpByHashWithEthClient(eth))
	// c.UseLogger(logr)
	c.UseModules(
		check.ValidateOpValues(),
		paymaster.CheckStatus(),
		check.SimulateOp(),
		paymaster.IncOpsSeen(),
	)

	op, err := UserOpMint(ethgo.HexToAddress("0x32A629dE3fb4549EB2B204d37eb9C8CFb0b9AdCf"),
		chain.MockMumbaiAddr,
		ethgo.HexToAddress("0x32A629dE3fb4549EB2B204d37eb9C8CFb0b9AdCf"),
		ethgo.Ether(100))

	require.NoError(t, err)

	ei := gas.EstimateInput{
		Rpc:         rpc,
		EntryPoint:  conf.SupportedEntryPoints[0],
		Op:          op,
		Ov:          ov,
		ChainID:     chainId,
		MaxGasLimit: big.NewInt(1_000_000),
	}
	vgas, cgas, err := gas.EstimateGas(&ei)
	require.NoError(t, err)
	require.True(t, vgas > 0)
	require.True(t, cgas > 0)

	//res, err := c.SendUserOperation(MockUserOpData, DefaultEntryPoint)
	//require.NoError(t, err)
	//_ = res
}
