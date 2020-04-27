# ufo-adapter

## 官方资料

### 官网

http://www.ufo.link/

### 接口文档

#### Wallet API

https://github.com/ufo-project/ufochain/wiki/UFO-wallet-protocol-API

### 浏览器

http://explorer.ufo.link/

### 配置文件

财务系统集成ufo-adapter，通过AssetsAdapter接口加载如下配置UFO.ini：

```ini

# ufo Wallet RPC API
walletapi = "http://127.0.0.1:20001/api/wallet"

# ufo explore API
explorerapi = "http://127.0.0.1:20005"

# Fix Transaction Fess
fixfees = "0.000001"

# log debug info
logdebug = false

# Log file path
logdir = "./logs/"

# single node
enablesingle = true


```

系统集成ufo-adapter/ufo包功能

```go

    //创建ufo钱包管理对象
	clientNode := ufo.NewWalletManager()
	
	//加载client.ini配置文件
	c, err := config.NewConfig("ini", "UFO.ini")
	if err != nil {
		return nil
	}
	clientNode.LoadAssetsConfig(c)
	
	//向远程服务，创建用户托管钱包的地址
	addrs, err := clientNode.CreateRemoteWalletAddress(100, 10)
	if err != nil {
        return
	}
	
	//获取钱包余额
	balanceLocal, err := clientNode.GetLocalWalletBalance()
    	
	//发起转账交易
    rawTx := &openwallet.RawTransaction{
        To: map[string]string{
            "3b769e29f6e2fc59fb7d1cd88fa03bd0777318b83d0e5111941992ad5efbe670d31": "0.0000001",
        },
        FeeRate: "",
    }

    txdecoder := clientNode.TxDecoder
    tx, err := txdecoder.SubmitRawTransaction(nil, rawTx)
    
    //启动区块链扫描器
    scanner := clientNode.GetBlockScanner()
	scanner.Run()
	
```

### 注意事项

`钱包数据备份`

由于ufo无法适配openwallet钱包体系，所以地址私钥等都托管在ufo钱包上。
钱包管理员在安装ufo钱包后，需要备份好助记词和密码，定时备份wallet.db。

