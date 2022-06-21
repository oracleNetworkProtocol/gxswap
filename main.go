package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/oracleNetworkProtocol/cosmoswap/contracts/factory"
	"github.com/oracleNetworkProtocol/cosmoswap/contracts/lp"
	"github.com/oracleNetworkProtocol/cosmoswap/contracts/router"
	"github.com/oracleNetworkProtocol/cosmoswap/contracts/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	//连接的8545服务
	blockchain, _ = ethclient.Dial("http://127.0.0.1:8545")

	//地址的私钥
	myKey, _ = crypto.HexToECDSA(strings.TrimPrefix("私钥", "0x"))
	auth, _  = bind.NewKeyedTransactorWithChainID(myKey, big.NewInt(520))

	//自己的地址
	myAddress = common.HexToAddress("EIP-55地址")

	//需要操作的合约
	myTokenAddress = common.HexToAddress("EIP-55合约地址")
	myToken, _     = token.NewToken(myTokenAddress, blockchain)

	//wPlug地址
	wPlugAddress = common.HexToAddress("0x6A9CdB0d8De5487F96C70b1f7857041e1f1CEd6d")

	//路由合约地址
	routerContractAddress = common.HexToAddress("0x373E80e79d92eEdC5159202e90c0693bA2Bbf0Be")
	routerContract, _     = router.NewCosmoswapRouter02(routerContractAddress, blockchain)

	//工厂合约地址
	factoryContractAddress = common.HexToAddress("0x3A3Ce66B80D05E7dDB8119E4ac0b27Bdee2a2dDC")
	factoryContract, _     = factory.NewICosmoswapFactory(factoryContractAddress, blockchain)
)

func main() {
	// sell()
	// addLiquidityPLUG()
	removeLiquidityPLUG()
	// transfer()
	// balanceOf()
	// compileMytoken()
	// buy()
}

//buy plug兑换token
func buy(uplugcnNumber int64, tokenNumber int64) {
	trans, err := routerContract.SwapExactPLUGForTokens(
		&bind.TransactOpts{
			From:     myAddress,
			Signer:   auth.Signer,
			Value:    big.NewInt(uplugcnNumber), //使用的uplugcn数量
			GasPrice: big.NewInt(7),
		},
		big.NewInt(tokenNumber), //最低要兑换的token数量
		[]common.Address{ //交易对合约地址(0:wplug,1:token)
			wPlugAddress,   //Wrapped Plugcn (WPLUG) 合约地址
			myTokenAddress, //mytoken 合约地址
		},
		myAddress, //token输出地址
		big.NewInt(time.Now().Add(30*time.Minute).Unix()), //超时时间
	)
	if err != nil {
		log.Fatalf("TransferFrom err: %v \n", err)
	}
	fmt.Println("tx sent: ", trans.Hash().Hex())
}

//sell token兑换plug
func sell() {
	result, err := myToken.Approve(
		&bind.TransactOpts{
			From:   myAddress,
			Signer: auth.Signer,
			Value:  nil,
		},
		routerContractAddress,
		big.NewInt(9500000000),
	)

	if err != nil {
		fmt.Println("Approve:", err)
		return
	}
	fmt.Println(result.Hash().Hex())

	//如果mytoken的transfer方法有额外的销毁和转账操作需要使用routerContract.SwapExactTokensForPLUGSupportingFeeOnTransferTokens()
	trans, err := routerContract.SwapExactTokensForPLUG(
		&bind.TransactOpts{
			From:     myAddress,
			Signer:   auth.Signer,
			Value:    nil,
			GasPrice: big.NewInt(7),
		},
		big.NewInt(9500000000), //使用的token数量
		big.NewInt(1000000),    //最低需要兑换的uplugcn数量
		[]common.Address{
			myTokenAddress, //mytoken 合约地址
			wPlugAddress,   //Wrapped Plugcn (WPLUG) 合约地址
		},
		myAddress, //plug输出地址
		big.NewInt(time.Now().Add(30*time.Minute).Unix()), //超时时间
	)
	if err != nil {
		fmt.Println("sell err:", err)
		return
	}
	fmt.Println(trans.Hash().Hex())
}

func getPair() common.Address {
	addr, err := factoryContract.GetPair(
		&bind.CallOpts{
			From: myAddress,
		},
		wPlugAddress,
		myTokenAddress,
	)
	if err != nil {
		log.Fatalf("TransferFrom err: %v \n", err)
	}
	return addr
}

//removeLiquidityPLUG 撤资
func removeLiquidityPLUG() {
	lpTokenAddress := getPair()
	lpToken, _ := lp.NewCosmoswapPair(lpTokenAddress, blockchain)
	//授权给lp合约操作权限
	lpToken.Approve(
		&bind.TransactOpts{
			From:   myAddress,
			Signer: auth.Signer,
			Value:  nil,
		},
		routerContractAddress, //获取lp地址
		big.NewInt(1000),      //授权数量
	)

	// lpBalances, err := lpToken.BalanceOf(
	// 	&bind.CallOpts{
	// 		From: myAddress,
	// 	},
	// 	myAddress,
	// )

	//如果mytoken的transfer方法有额外的销毁和转账操作需要使用routerContract.RemoveLiquidityPLUGSupportingFeeOnTransferTokens方法
	trans, err := routerContract.RemoveLiquidityPLUG(
		&bind.TransactOpts{
			From:     myAddress,
			Signer:   auth.Signer,
			Value:    nil,
			GasPrice: big.NewInt(7),
		},
		myTokenAddress,   // mytoken 合约地址
		big.NewInt(1000), //lp数量
		big.NewInt(0),    //最低撤资token数量
		big.NewInt(0),    //最低撤资plug数量
		myAddress,        //输出地址
		big.NewInt(time.Now().Add(30*time.Minute).Unix()), //超时时间
	)
	if err != nil {
		fmt.Println("remove liquidity err:", err)
		return
	}
	fmt.Println(trans.Hash().Hex())
}

func addLiquidityPLUG() {
	//进行代币授权
	result, err := myToken.Approve(
		&bind.TransactOpts{
			From:   myAddress,
			Signer: auth.Signer,
			Value:  nil,
		},
		routerContractAddress,   //路由合约地址
		big.NewInt(60000000000), //授权数量
	)

	if err != nil {
		fmt.Println("Approve:", err)
		return
	}
	fmt.Println(result.Hash().Hex())

	//添加代币兑换uplugcn池
	trans, err := routerContract.AddLiquidityPLUG(
		&bind.TransactOpts{
			From:     myAddress,
			Signer:   auth.Signer,
			Value:    big.NewInt(60000000), //注入uplugcn数量
			GasPrice: big.NewInt(7),
		},
		myTokenAddress,          // mytoken 合约地址
		big.NewInt(60000000000), //注入代币数量
		big.NewInt(0),
		big.NewInt(0),
		myAddress, //获得lp地址
		big.NewInt(time.Now().Add(30*time.Minute).Unix()), //超时时间
	)
	if err != nil {
		fmt.Println("add liquidity err:", err)
		return
	}
	fmt.Println(trans.Hash().Hex())
}

func compileMytoken() {
	addr, tx, token, err := token.DeployToken(auth, blockchain, big.NewInt(1000000000000000), "dog coin", "dog")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("addr:", addr)
	fmt.Println("tx:", tx)
	fmt.Println("token:", token)
}

func balanceOf() {
	result, err := myToken.BalanceOf(nil, myAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("余额:", result)
}

func transfer() {
	//转账
	var val, _ = big.NewInt(1).SetString("10000000000000000000", 10)
	tx, err := myToken.Transfer(&bind.TransactOpts{
		From:   myAddress,
		Signer: auth.Signer,
		Value:  nil,
		// GasLimit: 22300,
		// GasPrice: big.NewInt(7),
	}, common.HexToAddress("0x2CB81fa59873FfD57595D2753F9511ba76E83abf"), val)
	if err != nil {
		log.Fatalf("TransferFrom err: %v \n", err)
	}
	fmt.Printf("tx sent: %s \n", tx.Hash().Hex())
}
