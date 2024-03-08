package aptos

import (
	"encoding/hex"

	"github.com/xiaohuasheng0x1/blockchains/crypto/ed25519"
	"github.com/xiaohuasheng0x1/blockchains/wallet"
)

const HexPrefix = "0x"

type AptosWallet struct {
	wallet.WalletBase
}

func (aw *AptosWallet) GetRandomPrivateKey() (string, error) {
	p, err := ed25519.GenerateKey()
	if err != nil {
		return "", err
	}
	return HexPrefix + hex.EncodeToString(p), nil
}
