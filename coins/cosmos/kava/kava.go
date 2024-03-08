package kava

import (
	"github.com/xiaohuasheng0x1/blockchains/coins/cosmos"
)

const (
	HRP = "kava"
)

func NewAddress(privateKeyHex string) (string, error) {
	return cosmos.NewAddress(privateKeyHex, HRP, false)
}

func ValidateAddress(address string) bool {
	return cosmos.ValidateAddress(address, HRP)
}
