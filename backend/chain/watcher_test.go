package chain

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"github.com/umbracle/ethgo/blocktracker"
	"github.com/umbracle/ethgo/jsonrpc"
	"github.com/umbracle/ethgo/tracker"
	boltdbStore "github.com/umbracle/ethgo/tracker/store/boltdb"
	"math/big"
	"os"
	"testing"
)

// event Transfer(address indexed from, address indexed to, uint256 value);
var transferEvent = abi.MustNewEvent(`event Transfer(
	address indexed from,
	address indexed to,
	uint256 value
)`)

func noTestWatcher(t *testing.T) {
	ec, err := jsonrpc.NewClient(os.Getenv("CHAIN_URL"))
	require.NoError(t, err)

	store, err := boltdbStore.New("transfer.db")
	require.NoError(t, err)

	bt := blocktracker.NewBlockTracker(ec.Eth(), blocktracker.WithBlockMaxBacklog(100))
	tt, err := tracker.NewTracker(ec.Eth(),
		tracker.WithBatchSize(20000),
		tracker.WithStore(store),
		tracker.WithBlockTracker(bt),
		// tracker.WithEtherscan(os.Getenv("ETHERSCAN_APIKEY")),
		tracker.WithFilter(&tracker.FilterConfig{
			Async: true,
			Address: []ethgo.Address{
				SFLUVPolygonMainnetV1_1,
			},
			Start: SFLUV_V1_1_Deploy_Block,
		}),
	)
	require.NoError(t, err)

	lastBlock, err := tt.GetLastBlock()
	require.NoError(t, err)
	if lastBlock != nil {
		fmt.Printf("Last block processed: %d\n", lastBlock.Number)
	}

	go func() {
		if err := tt.Sync(context.Background()); err != nil {
			fmt.Printf("[ERR]: %v", err)
		}
	}()

	for {
		select {
		case evnt := <-tt.EventCh:
			for _, log := range evnt.Added {
				if transferEvent.Match(log) {
					vals, err := transferEvent.ParseLog(log)
					if err != nil {
						panic(err)
					}
					_ = vals

					toAddr, ok := vals["to"].(ethgo.Address)
					require.True(t, ok)

					if toAddr == SFLUV_V1_Concierge {
						fromAddr, ok := vals["from"].(ethgo.Address)
						require.True(t, ok)
						value, ok := vals["value"].(*big.Int)
						require.True(t, ok)
						fmt.Printf("got concierge tx: from '%v', value '%v'", fromAddr.String(), value.String())
					}

					//index := binary.LittleEndian.Uint64(vals["index"].([]byte))
					//amount := binary.LittleEndian.Uint64(vals["amount"].([]byte))
					//fmt.Printf("Deposit: Block %d Index %d Amount %d\n", log.BlockNumber, index, amount)
				}
			}
		case <-tt.DoneCh:
			fmt.Println("historical sync done")
		}
	}

	// ctx, cancelFn := context.WithCancel(context.Background())
}
