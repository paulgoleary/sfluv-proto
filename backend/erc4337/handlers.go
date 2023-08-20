package erc4337

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paulgoleary/local-luv-proto/chain"
	"github.com/paulgoleary/local-luv-proto/config"
	"github.com/stackup-wallet/stackup-bundler/pkg/userop"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
	"math/big"
	"net/http"
	"strings"
)

// GET? erc4337/userop/approve?target=XXXX&spender=YYYY&amount=10000&owner=ZZZZ
var addrHexLength = len(ethgo.ZeroAddress.String())

func handleRequiredAddress(addrHex string) (ret *ethgo.Address) {
	if !strings.HasPrefix(addrHex, "0x") {
		addrHex = "0x" + addrHex
	}
	if len(addrHex) != addrHexLength {
		return
	}
	if addr := ethgo.HexToAddress(addrHex); addr != ethgo.ZeroAddress {
		ret = &addr
	}
	return
}

type HandlerContext struct {
	ChainId    *big.Int
	EntryPoint *contract.Contract
}

func MakeContext(config config.Config) (*HandlerContext, error) {
	abiBytes, err := abiIEP.ReadFile("abi/IEntryPoint.json")
	if err != nil {
		return nil, err
	}

	rpc, err := jsonrpc.NewClient(config.ChainRpcUrl)
	if err != nil {
		return nil, err
	}

	chainId, err := rpc.Eth().ChainID()
	if err != nil {
		return nil, err
	}
	hc := &HandlerContext{ChainId: chainId}

	hc.EntryPoint, err = chain.LoadReadContractAbi(rpc, abiBytes, DefaultEntryPoint)
	if err != nil {
		return nil, err
	}

	return hc, nil
}

func (hc *HandlerContext) GetOwnerInfo(ownerAddr ethgo.Address) (nonce *big.Int, senderAddr ethgo.Address, err error) {

	var ownerInitCode []byte
	if ownerInitCode, err = MakeDefaultInitCode(ownerAddr); err != nil {
		return
	}

	_, err = hc.EntryPoint.Call("getSenderAddress", ethgo.Latest, ownerInitCode)
	// this method is expected to revert
	if senderAddr, err = getSenderAddressFromError(err); err != nil {
		return
	}

	var res map[string]interface{}
	if res, err = hc.EntryPoint.Call("getNonce", ethgo.Latest, senderAddr, big.NewInt(0)); err != nil {
		return
	}
	var ok bool
	if nonce, ok = res["nonce"].(*big.Int); !ok {
		err = fmt.Errorf("unexpected - expected *big.Int for nonce return value")
	}
	return
}

func (hc *HandlerContext) HandleUserOpApprove(c *gin.Context) {

	q := c.Request.URL.Query()

	targetAddr := handleRequiredAddress(q.Get("target"))
	spenderAddr := handleRequiredAddress(q.Get("spender"))
	ownerAddr := handleRequiredAddress(q.Get("owner"))

	amount, ok := new(big.Int).SetString(q.Get("amount"), 10)

	if targetAddr == nil || spenderAddr == nil || ownerAddr == nil || !ok {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid or missing parameter(s)"))
		return
	}

	nonce, senderAddr, err := hc.GetOwnerInfo(*ownerAddr)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if op, err := UserOpApprove(nonce, *ownerAddr, senderAddr, *targetAddr, *spenderAddr, amount); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else {
		opJson, _ := op.ToMap()
		c.JSON(http.StatusOK, opJson)
	}
}

func (hc *HandlerContext) HandleUserOpWithdrawTo(c *gin.Context) {

	q := c.Request.URL.Query()

	targetAddr := handleRequiredAddress(q.Get("target"))
	toAddr := handleRequiredAddress(q.Get("to"))
	ownerAddr := handleRequiredAddress(q.Get("owner"))

	amount, ok := new(big.Int).SetString(q.Get("amount"), 10)

	if targetAddr == nil || toAddr == nil || ownerAddr == nil || !ok {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid or missing parameter(s)"))
		return
	}

	nonce, senderAddr, err := hc.GetOwnerInfo(*ownerAddr)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if op, err := UserOpWithdrawTo(nonce, *ownerAddr, senderAddr, *targetAddr, *toAddr, amount); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else {
		opJson, _ := op.ToMap()
		c.JSON(http.StatusOK, opJson)
	}
}

type userOpSendRequest struct {
	EntryPointAddr string         `json:"entryPoint"`
	Op             map[string]any `json:"op"`
}

func (hc *HandlerContext) HandleUserOpSend(c *gin.Context) {
	req := userOpSendRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if userOp, err := userop.New(req.Op); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else {
		ownerAddr, err := UserOpEcrecover(userOp, hc.ChainId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("ecrecover failure: %v", err.Error()))
			return
		}

		_, senderAddr, err := hc.GetOwnerInfo(ownerAddr)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if senderAddr.String() != userOp.Sender.String() {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("op sender address does not match recovered sender address"))
			return
		}
	}
	c.JSON(http.StatusOK, fmt.Sprintf(`{"op hash":"%v"}`, "TODO!!!"))
	return
}
