# oasis-sdk
Oasis SDK is used to interact with the Oasis blockchain, it contains various functions can be used to web3 wallet.

## Installation

### go get

To obtain the latest version, simply require the project using :

```shell
go get -u github.com/xiaohuasheng0x1/blockchains/coins/oasis
```

## Usage
### New Address
```go
	privateKeyHex := "a30a45ef8c019d22b7e8d18f11197677bff80ff4d2f23ab9ac14bdbac32c86e7baf40754ed3843e0464f814c3c605d8c36500cfb6892e2bd441839102f4200ed"
    address, err := NewAddress(privateKeyHex)
	if err != nil {
		// todo
	}
```

###  Transfer
```go
	pk := "a30a45ef8c019d22b7e8d18f11197677bff80ff4d2f23ab9ac14bdbac32c86e7baf40754ed3843e0464f814c3c605d8c36500cfb6892e2bd441839102f4200ed"
	chainId := "b11b369e0da5bb230b220127f5e7b242d385ef8c6f54906243f30af63c815535"
	toAddr := "oasis1qqx0wgxjwlw3jwatuwqj6582hdm9rjs4pcnvzz66"
	amount := big.NewInt(100000000)
	feeAmount := big.NewInt(0)
	gas := uint64(2000)
	nonce := uint64(8)
	tx := NewTransferTx(nonce, gas, feeAmount, toAddr, amount)
	signedTx := SignTransaction(pk, chainId, tx)
	signedTxBytes, err := cbor.Marshal(signedTx)
	if err != nil {
		// todo
	}
```

## License
Most packages or folder are [MIT](<https://github.com/xiaohuasheng0x1/blockchains/blob/main/coins/oasis/LICENSE>) licensed, see package or folder for the respective license.
