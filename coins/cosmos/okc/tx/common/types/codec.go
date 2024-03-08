package types

import (
	"github.com/xiaohuasheng0x1/blockchains/coins/cosmos/okc/tx/amino"
)

// Register the sdk message type
func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterInterface((*Msg)(nil), nil)
	cdc.RegisterInterface((*Tx)(nil), nil)
}

func init() {
	RegisterCodec(amino.GCodec)
}
