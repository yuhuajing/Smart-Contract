package main

import (
	"fmt"
	"log"
	"main/store"

	"github.com/ethereum/go-ethereum/common"
)

func query() {
	address := common.HexToAddress("0x9f07b5fE8D3b0Fd020144Bbd70E3481223fD547c")
	instance, err := store.NewStore(address, client)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	version, err := instance.Version(nil)
	if err != nil {
		fmt.Println("error querying version")
		log.Fatal(err)
	}
	fmt.Println(version)
}
