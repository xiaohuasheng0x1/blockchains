# stacks-sdk
Stacks SDK is used to interact with the Stacks blockchain, it contains various functions can be used to web3 wallet.

## Installation

### go get

To obtain the latest version, simply require the project using :

```shell
go get -u github.com/xiaohuasheng0x1/blockchains/coins/stacks
```

## Usage
### New Address
```go
	priKey := "598d99970d04be67e8b41ddd5c5453487eeab5345ea1638c9a2849dee377f2a301"
	addr, err := NewAddress(priKey)
	if err != nil {
		// todo
	}
```

###  Transfer
```go
	result, err := Transfer("598d99970d04be67e8b41ddd5c5453487eeab5345ea1638c9a2849dee377f2a3", "SP2P58SJY1XH6GX4W3YGEPZ2058DD3JHBPJ8W843Q", "20", big.NewInt(3000), big.NewInt(8), big.NewInt(200))
	if err != nil {
		// todo
	}
```

## License
Most packages or folder are [MIT](<https://github.com/xiaohuasheng0x1/blockchains/blob/main/coins/stacks/LICENSE>) licensed, see package or folder for the respective license.
