package filecoin

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xiaohuasheng0x1/blockchains/util"
)

func TestNewTx(t *testing.T) {
	from := "f1bh3d2y6xxugpg3ygzxnjhcrs5ffxh5nvqmanbia"
	to := from
	nonce := 1
	value := big.NewInt(100)
	gasLimit := 608335
	gasFeeCap := big.NewInt(1643831112)
	gasPremium := big.NewInt(99707)
	method := 0
	message := NewTx(from, to, nonce, method, gasLimit, value, gasFeeCap, gasPremium)
	h := util.EncodeHexWith0x(message.Hash())
	assert.Equal(t, "0xe496300a898281703d377eaebe15b362571e813dc924d5a846815bda07c4b960", h)
}
