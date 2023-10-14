package erc4337

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"github.com/paulgoleary/local-luv-proto/chain"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"github.com/umbracle/ethgo/jsonrpc"
	"math/big"
)

// event Transfer(address indexed from, address indexed to, uint256 value);
var transferEvent = abi.MustNewEvent(`event Transfer(
	address indexed from,
	address indexed to,
	uint256 value
)`)

func StartConcierge(withKey ethgo.Key, chainUrl, dataDir string) (context.CancelFunc, error) {

	w, err := chain.NewWatcher(chain.SFLUVPolygonMainnetV1_1, transferEvent, chainUrl, dataDir)
	if err != nil {
		return nil, err
	}

	ec, err := jsonrpc.NewClient(chainUrl)
	if err != nil {
		return nil, err
	}

	abiBytes, err := abiIEP.ReadFile("abi/SFLUVv1.json")
	if err != nil {
		return nil, err
	}

	ep, err := chain.LoadReadContractAbi(ec, abiBytes, DefaultEntryPoint, withKey)
	if err != nil {
		return nil, err
	}

	handleAdded := func(logEvt *ethgo.Log) error {
		vals, err := transferEvent.ParseLog(logEvt)
		if err != nil {
			return err
		}

		toAddr, ok := vals["to"].(ethgo.Address)
		if !ok {
			return fmt.Errorf("should not happen - no 'to' field in transfer event")
		}

		if toAddr == chain.SFLUV_V1_Concierge {
			fromAddr, okFrom := vals["from"].(ethgo.Address)
			value, okVal := vals["value"].(*big.Int)
			if !okFrom || !okVal {
				return fmt.Errorf("should not happen - no 'from' or 'value' field in transfer event")
			}
			log.Infof("got concierge tx: from '%v', value '%v'", fromAddr.String(), value.String())

			if retBalance, err := ep.Call("balanceOf", ethgo.Latest, toAddr); err != nil {
				return err
			} else if currentBalance, ok := retBalance["0"].(*big.Int); !ok || currentBalance.Cmp(value) < 0 {
				gotStr := "[UNKNOWN]"
				if currentBalance != nil {
					gotStr = currentBalance.String()
				}
				log.Warnf("should not happen - concierge has insufficient balance: want %v, got %v, to addr '%v'",
					value.String(), gotStr, fromAddr.String())
				return fmt.Errorf("should not happen - concierge has insufficient balance")
			}

			// TODO: this is *very* naive atm. at minimum, should check that balance exists (DONE) and likely manage state to be sure
			//  balances are not unwrapped more than once in case of re-orgs, bugs, etc...
			if err = chain.TxnDoWait(ep.Txn("withdrawTo", fromAddr, value)); err != nil {
				return err
			}
		}

		return nil
	}

	return w.GetCancelFunc(), w.Start(handleAdded)
}
