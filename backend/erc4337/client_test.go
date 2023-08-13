package erc4337

import (
	"context"
	"github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/paulgoleary/local-luv-proto/chain"
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/stackup-wallet/stackup-bundler/pkg/client"
	"github.com/stackup-wallet/stackup-bundler/pkg/gas"
	"github.com/stackup-wallet/stackup-bundler/pkg/mempool"
	"github.com/stackup-wallet/stackup-bundler/pkg/modules/checks"
	"github.com/stackup-wallet/stackup-bundler/pkg/modules/paymaster"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
	"github.com/umbracle/ethgo/wallet"
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
		SupportedEntryPoints:    []common.Address{common.Address(DefaultEntryPoint)},
		EthClientUrl:            os.Getenv("CHAIN_URL"),
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

	op, err := UserOpMint(
		ethgo.HexToAddress("0x32A629dE3fb4549EB2B204d37eb9C8CFb0b9AdCf"),
		ethgo.HexToAddress("0x054dF6203225bB58d9243eBf9DAd55608a436042"),
		chain.MockMumbaiAddr,
		ethgo.HexToAddress("0x32A629dE3fb4549EB2B204d37eb9C8CFb0b9AdCf"),
		ethgo.Ether(100))
	require.NoError(t, err)

	sk, err := crypto.SKFromHex(os.Getenv("CHAIN_SK"))
	require.NoError(t, err)
	k := &chain.EcdsaKey{SK: sk}

	// don't *actually* need to sign this correctly afaik but wth ...
	opHash := op.GetUserOpHash(common.Address(DefaultEntryPoint), chainId)
	opEthHash := crypto.EthSignedMessageHash(opHash.Bytes())

	sig, err := crypto.Sign(k.SK, opEthHash)
	require.NoError(t, err)

	checkRecover, err := wallet.Ecrecover(opEthHash, sig)
	require.NoError(t, err)
	require.Equal(t, k.Address(), checkRecover)

	// ETH magic ...?
	sig[64] += 27
	op.Signature = sig

	opMap, _ := op.ToMap()

	//resEst, err := c.EstimateUserOperationGas(opMap, DefaultEntryPoint.String())
	//require.NoError(t, err)
	//_ = resEst

	ec, err := jsonrpc.NewClient(os.Getenv("CHAIN_URL"))
	require.NoError(t, err)

	ep, err := chain.LoadContract(ec, "IEntryPoint.sol/IEntryPoint", k, DefaultEntryPoint)
	require.NoError(t, err)
	userOpHashRes, err := ep.Call("getUserOpHash", ethgo.Latest, opMap)
	require.NoError(t, err)
	checkHash, ok := userOpHashRes["0"].([32]byte)
	require.True(t, ok)
	println(hexutil.Encode(checkHash[:]))
	println(opHash.String())

	resSend, err := c.SendUserOperation(opMap, DefaultEntryPoint.String())
	require.NoError(t, err)
	_ = resSend

}

func TestStakeManager(t *testing.T) {

	ec, err := jsonrpc.NewClient(os.Getenv("CHAIN_URL"))
	require.NoError(t, err)

	sk, err := crypto.SKFromHex(os.Getenv("CHAIN_SK"))
	require.NoError(t, err)
	k := &chain.EcdsaKey{SK: sk}

	ep, err := chain.LoadContract(ec, "IEntryPoint.sol/IEntryPoint", k, DefaultEntryPoint)
	require.NoError(t, err)

	tx, txErr := ep.Txn("depositTo", ethgo.HexToAddress("0x054dF6203225bB58d9243eBf9DAd55608a436042"))
	require.NoError(t, txErr)
	tx.WithOpts(&contract.TxnOpts{Value: ethgo.Ether(1)})

	err = chain.TxnDoWait(tx, nil)
	require.NoError(t, err)

	res, err := ep.Call("balanceOf", ethgo.Latest, ethgo.HexToAddress("0x054dF6203225bB58d9243eBf9DAd55608a436042"))
	require.NoError(t, err)
	checkBalance, ok := res["0"].(*big.Int)
	require.True(t, ok)
	require.True(t, checkBalance.Cmp(big.NewInt(0)) > 0)

}
