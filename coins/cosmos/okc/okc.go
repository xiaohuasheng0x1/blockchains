package okc

import (
	"github.com/xiaohuasheng0x1/blockchains/coins/cosmos"
)

const (
	HRP = "ex"
)

// NewAddress method of eth is used
func NewAddress(privateKey string) (string, error) {
	return cosmos.NewAddress(privateKey, HRP, true)
}

func ValidateAddress(address string) bool {
	return cosmos.ValidateAddress(address, HRP)
}
