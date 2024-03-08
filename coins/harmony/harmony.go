package harmony

import (
	"encoding/hex"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/xiaohuasheng0x1/blockchains/coins/ethereum"
)

const HRP = "one"

func NewAddress(seedHex string, followETH bool) (string, error) {
	p, err := hex.DecodeString(seedHex)
	if err != nil {
		return "", err
	}
	prvKey, _ := btcec.PrivKeyFromBytes(p)
	addr := ethereum.GetNewAddress(prvKey.PubKey())
	if followETH {
		return addr, nil
	}
	bytes, err := hex.DecodeString(addr[2:])
	if err != nil {
		return "", err
	}
	bech32Address, err := bech32.EncodeFromBase256(HRP, bytes)
	if err != nil {
		return "", err
	}
	return bech32Address, nil
}

func Transfer(transaction *ethereum.EthTransaction, chainId *big.Int, prvKey *btcec.PrivateKey) (string, error) {
	return transaction.SignTransaction(chainId, prvKey)
}
