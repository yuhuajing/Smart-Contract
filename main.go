package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	privateKey = "af2bc4375adfc4491c35136f11bb65690b10390151a0ec9dc0b44c27bba8bc54"
	client, _  = ethclient.Dial("https://goerli.infura.io/v3/d7b27eea18a54fb389c2562ba19f8e36")
)

//Get goerli ETH
//https://testnet.help/en/ethfaucet/goerli#log

func main() {
	//deploy()
	query()
}
