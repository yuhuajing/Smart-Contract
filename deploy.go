package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"main/store"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func deploy() {
	chainID, _ := client.ChainID(context.Background())
	_privateKey, _ := crypto.HexToECDSA(privateKey)
	pubkey := _privateKey.Public()
	publicEcdsa, _ := pubkey.(*ecdsa.PublicKey)
	fromaddress := crypto.PubkeyToAddress(*publicEcdsa)
	nonce, _ := client.PendingNonceAt(context.Background(), fromaddress)
	balance, _ := client.PendingBalanceAt(context.Background(), fromaddress)
	fmt.Println(balance.String())
	gasprice, _ := client.SuggestGasPrice(context.Background())
	fmt.Println(gasprice)
	auth, _ := bind.NewKeyedTransactorWithChainID(_privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(600000) // in units
	fmt.Println(uint64(gasprice.Int64()) * auth.GasLimit)
	auth.GasPrice = gasprice
	version := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, version)
	if err == nil {
		fmt.Println(address.Hex())   //0x9f07b5fE8D3b0Fd020144Bbd70E3481223fD547c
		fmt.Println(tx.Hash().Hex()) //0x20eb99563e0f40fb015b8baca3a101a0993af436f3f0c067720523914d41273b
	} else {
		fmt.Println(err)
	}
	_ = instance
	nbalance, _ := client.PendingBalanceAt(context.Background(), fromaddress)
	fmt.Println(nbalance.String())
}
