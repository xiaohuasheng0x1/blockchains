package iris

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xiaohuasheng0x1/blockchains/coins/cosmos"
)

// https://lcd-iris.keplr.app/cosmos/auth/v1beta1/accounts/iaa1rvs5xph4l3px2efynqsthus8p6r4exyrt6k4ya
// curl -X POST -d '{"tx_bytes":"CrsBCrgBCikvaWJjLmFwcGxpY2F0aW9ucy50cmFuc2Zlci52MS5Nc2dUcmFuc2ZlchKKAQoIdHJhbnNmZXISCWNoYW5uZWwtMxoOCgV1aXJpcxIFMTAwMDAiKmlhYTFydnM1eHBoNGwzcHgyZWZ5bnFzdGh1czhwNnI0ZXh5cnQ2azR5YSorb3NtbzFydnM1eHBoNGwzcHgyZWZ5bnFzdGh1czhwNnI0ZXh5cmtyOTVzNzIAOICU1ansnMKCFxJoClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDnCTAPeltdUKTRGnHRm5QS12GCMOZ3du0gUsZYfj3OlMSBAoCCAEYAxIUCg4KBXVpcmlzEgUyNDAwMBCgjQYaQBz9RpSzDcFmuye06mbliAL/ieZL6MYxOk4g9+kxdxAuQfsHpFmyNvsUQZ6ybpkUN5zxt+/yUEiiw0VkZUZ9R1k=","mode":"BROADCAST_MODE_SYNC"}' https://lcd-iris.keplr.app/cosmos/tx/v1beta1/txs
func TestTransfer(t *testing.T) {
	privateKeyHex := "1790962db820729606cd7b255ace1ac5ebb129ac8e9b2d8534d022194ab25b37"
	address, err := NewAddress(privateKeyHex)
	require.Nil(t, err)
	expected := "iaa145q0tcdur4tcx2ya5cphqx96e54yflfy4sruf9"
	require.Equal(t, expected, address)

	param := cosmos.TransferParam{}
	param.FromAddress = address
	param.ToAddress = address
	param.Demon = "uiris"
	param.Amount = "10000"
	param.CommonParam.ChainId = "irishub-1"
	param.CommonParam.Sequence = 0
	param.CommonParam.AccountNumber = 46542
	param.CommonParam.FeeDemon = "uiris"
	param.CommonParam.FeeAmount = "24000"
	param.CommonParam.GasLimit = 100000
	param.CommonParam.Memo = ""
	param.CommonParam.TimeoutHeight = 0
	signedTx, err := cosmos.Transfer(param, privateKeyHex)
	require.Nil(t, err)
	t.Log("signedTx : ", signedTx)
}

func TestIbcTransfer(t *testing.T) {
	privateKeyHex := "1790962db820729606cd7b255ace1ac5ebb129ac8e9b2d8534d022194ab25b37"
	p := cosmos.IbcTransferParam{}
	p.CommonParam.ChainId = "irishub-1"
	p.CommonParam.Sequence = 3
	p.CommonParam.AccountNumber = 46542
	p.CommonParam.FeeDemon = "uiris"
	p.CommonParam.FeeAmount = "24000"
	p.CommonParam.GasLimit = 100000
	p.CommonParam.Memo = ""
	p.CommonParam.TimeoutHeight = 0
	p.FromAddress = "iaa1rvs5xph4l3px2efynqsthus8p6r4exyrt6k4ya"
	p.ToAddress = "osmo1rvs5xph4l3px2efynqsthus8p6r4exyrkr95s7"
	p.Demon = "uiris"
	p.Amount = "10000"
	p.SourcePort = "transfer"
	p.SourceChannel = "channel-3"
	p.TimeOutInSeconds = uint64(time.Now().UnixMilli()/1000) + 300
	signedIBCTx, err := cosmos.IbcTransfer(p, privateKeyHex)
	require.Nil(t, err)
	t.Log("signedIBCTx : ", signedIBCTx)
}
