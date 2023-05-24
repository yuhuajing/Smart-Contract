# Smart-Contract

1.为了使用高级语言和合约代码交互，需要先使用solc工具将合约编译成ABI（application binary Interface)，并且将ABI文件format成golang可使用的格式
```json
solc --abi store.sol -o storeVersion.abi
```

2.将合约编译成字节码
```json
solc --bin store.sol  -o storageVersion.bin
```

3. 将 ABI 接口和 bin字节码作为输入，生成交互的golang代码：

First generate the abigen tool in Go-ethereum

```json
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
cd go-ethereum/cmd/abige
go build
```

```json
./abigen --bin=storageVersion.bin/store.bin --abi=storeVersion.abi/store.abi --pkg=store --out=Store.go
```

4. Deploy the contract

Connect to testNet、construct the tx、deploy the contract

