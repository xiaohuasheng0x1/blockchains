package util

import (
	"math/big"
	"strconv"
)

func ConvertToUint64(v string) uint64 {
	i, _ := strconv.ParseUint(v, 10, 64)
	return i
}

func ConvertToBigInt(v string) *big.Int {
	i, _ := new(big.Int).SetString(v, 10)
	return i
}
