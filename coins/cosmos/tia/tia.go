package tia

import (
	"github.com/xiaohuasheng0x1/blockchains/coins/cosmos"
)

const (
	HRP = "celestia"
)

func NewAddress(privateKeyHex string) (string, error) {
	return cosmos.NewAddress(privateKeyHex, HRP, false)
}

func ValidateAddress(address string) bool {
	return cosmos.ValidateAddress(address, HRP)
}

func Transfer(param cosmos.TransferParam, privateKeyHex string) (string, error) {
	return cosmos.Transfer(param, privateKeyHex)
}

func IbcTransfer(param cosmos.IbcTransferParam, privateKeyHex string) (string, error) {
	return cosmos.IbcTransfer(param, privateKeyHex)
}
