package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
	"github.com/umbracle/ethgo"
	"math/big"
)

// S256 is the secp256k1 elliptic curve
var S256 = btcec.S256()

// MarshalPublicKey marshals a public key on the secp256k1 elliptic curve.
func MarshalPublicKey(pub *ecdsa.PublicKey) []byte {
	return elliptic.Marshal(S256, pub.X, pub.Y)
}

// PubKeyToAddress returns the Ethereum address of a public key
func PubKeyToAddress(pub *ecdsa.PublicKey) ethgo.Address {
	buf := ethgo.Keccak256(MarshalPublicKey(pub)[1:])[12:]

	return ethgo.BytesToAddress(buf)
}

// Sign produces a compact signature of the data in hash with the given
// private key on the secp256k1 curve.
func Sign(priv *ecdsa.PrivateKey, hash []byte) ([]byte, error) {
	sig, err := btcec.SignCompact(S256, (*btcec.PrivateKey)(priv), hash, false)
	if err != nil {
		return nil, err
	}

	term := byte(0)
	if sig[0] == 28 {
		term = 1
	}

	return append(sig, term)[1:], nil
}

func SKFromHex(skHex string) (ret *ecdsa.PrivateKey, err error) {
	var skBytes []byte
	if skBytes, err = hex.DecodeString(skHex); err != nil {
		return
	}

	sk, _ := btcec.PrivKeyFromBytes(S256, skBytes)
	ret = (*ecdsa.PrivateKey)(sk)
	return
}

func SKFromInt(skInt *big.Int) (ret *ecdsa.PrivateKey, err error) {
	sk, _ := btcec.PrivKeyFromBytes(S256, skInt.Bytes())
	ret = (*ecdsa.PrivateKey)(sk)
	return
}

func RandSK() (ret *ecdsa.PrivateKey, err error) {
	var sk *btcec.PrivateKey
	if sk, err = btcec.NewPrivateKey(S256); err != nil {
		return
	}
	return SKFromInt(sk.D)
}
