package harmony

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/stretchr/testify/require"
	"github.com/xiaohuasheng0x1/blockchains/coins/ethereum"
)

func TestNewAddress(t *testing.T) {
	seedHex := "1790962db820729606cd7b255ace1ac5ebb129ac8e9b2d8534d022194ab25b37"
	addr, err := NewAddress(seedHex, true)
	require.Nil(t, err)
	expected := "0x97e2728c08bd0bfba631929e10bceaec8fc5c961"
	require.Equal(t, expected, addr)

	bytes, err := hex.DecodeString(addr[2:])
	require.Nil(t, err)
	bech32Address, err := bech32.EncodeFromBase256(HRP, bytes)
	require.Nil(t, err)
	expected = "one1jl389rqgh59lhf33j20pp082aj8utjtpuass6r"
	require.Equal(t, expected, bech32Address)
	hrp, addrBytes, err := bech32.DecodeToBase256(bech32Address)
	require.Nil(t, err)
	require.Equal(t, addr, "0x"+hex.EncodeToString(addrBytes))
	require.Equal(t, HRP, hrp)
}

func TestTransfer(t *testing.T) {
	p, _ := hex.DecodeString("1790962db820729606cd7b255ace1ac5ebb129ac8e9b2d8534d022194ab25b37")
	prvKey, _ := btcec.PrivKeyFromBytes(p)
	transaction := ethereum.NewEthTransaction(
		big.NewInt(int64(00)),
		big.NewInt(int64(420000)),
		big.NewInt(int64(200000000000)),
		big.NewInt(int64(100000000000)),
		"97e2728c08bd0bfba631929e10bceaec8fc5c961", "",
	)
	signedTx, err := Transfer(transaction, big.NewInt(int64(1666700000)), prvKey)
	require.Nil(t, err)
	expected := "0xf86e80852e90edd000830668a09497e2728c08bd0bfba631929e10bceaec8fc5c96185174876e8008084c6afa5e4a00eff8095fbcdd29d7d3e26e211f345c6b7f26c66bba360dff936608c056ab786a00c86c1ebc0792fa2c0d326d3652a756813674f514db93278ebcd084663ba5422"
	require.Equal(t, expected, signedTx)
}
