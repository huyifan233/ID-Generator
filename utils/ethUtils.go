package utils

import (
	"../conf/smartconrtact"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"sync"
	"time"
)






type Message struct {
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to"`
	Value    string          `json:"value"`
	Data     string          `json:"data"`
	GasLimit string          `json:"gasLimit"`
	GasPrice string          `json:"gasPrice"`
}

type Transaction struct {
	From string `json:"from"`
	Hash string `json:"hash"`
	Input string `json:"input"`
}

type Receipt struct {
	TransactionHash  string          `json:"transactionHash"`
	TransactionIndex string          `json:"transactionIndex"`
	BlockHash        string          `json:"blockHash"`
	BlockNumber      string          `json:"blockNumber"`
	To               *common.Address `json:"to"`
	From             common.Address  `json:"from"`
}

type Block struct {
	Number string `json:"number"`
	Hash   string `json:"hash"`
	//ParentHash string `json:"parentHash"`
	//nonce string `json:"nonce"`
	//sha3Uncles string `json:"sha3Uncles"`
	//logsBloom string `json:"logsBloom"`
	//TransactionsRoot string `json:"transactionsRoot"`
	//stateRoot string `json:"stateRoot"`
	//miner string `json:"miner"`
	//difficulty string `json:"difficulty"`
	//totalDifficulty string  `json:"totalDifficulty"`
	//extraData string `json:"extraData"`
	//size string  `json:"size"`
	//gasLimit string `json:"gasLimit"`
	//gasUsed string `json:"gasUsed"`
	//timestamp string `json:timestamp`
	Transactions []Transaction `json:transactions`
	//uncles []string `json:uncles`
}
var clientConnect *ethclient.Client
var mutex sync.Mutex
var simulate *backends.SimulatedBackend
var instance *generateId.GenerateId
var contractAddress common.Address
var autheration *bind.TransactOpts
var count int64
var IdChannel chan string
//server
//const key = `{"address":"48f84bdb03d39a35b55a6c32cbef1ae960ab17eb","crypto":{"cipher":"aes-128-ctr","ciphertext":"98bfb59d728625da2e321bcd260f099bd0f34b6173d32121c611d85ab7c2c6d2","cipherparams":{"iv":"8e549798d0261c433141f6200e00f2a3"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"4e787078e7cf0504dbe3a86acdf546e779ed9ad903063cf37e32f479653c92bb"},"mac":"29f98ef4e8bb062391b5efaa89a5eb146fe3b63d9f02aef9cd2c2690d31a488e"},"id":"cfc0ec80-12b2-4e95-a883-98534e2acae3","version":3}`
//local
const key = `{"address":"4e0dc4ffa4609989c00183bc4a78b5ea70c33494","crypto":{"cipher":"aes-128-ctr","ciphertext":"ab3eaddba953c02861e539ed93f5c28945267c3435b63c2531bb669f2ebd96cc","cipherparams":{"iv":"d271f17fbd59f57d1de068702b3817e4"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"400aacf467e873f9232418231abcc535934bb027db758fd637ff0739ff6a84f6"},"mac":"00fb7135ce6dbfef42676b786f2459ce894040b686101a35ea726d5ee2780e76"},"id":"064efa43-f6d2-4673-8c27-c56d197695b2","version":3}`

func ConnectEthereum(){
	Log.Info("以太坊连接: ",config.Url.EthereumUrl)
	if clientConnect == nil {
		mutex.Lock()
		if clientConnect == nil {
			conn, err := ethclient.Dial(config.Url.EthereumUrl)
			if err != nil{
				Log.Error("获取以太坊连接失败: ",err)
			}else{
				clientConnect = conn

			}
		}
		mutex.Unlock()
	}
}


func DeployContract(){

	auth, err := bind.NewTransactor(strings.NewReader(config.Ec.EthereumAdminAccount),config.Ec.EthereumAdminPassword)
	if err != nil{
		Log.Error("解锁管理员账户失败",err)
	}
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1337000000000)}

	balance, err := clientConnect.BalanceAt(context.TODO(),auth.From,nil)
	fmt.Println("From",auth.From.String())
	fmt.Println("Balance", balance)
	sim := backends.NewSimulatedBackend(alloc, 100000000)
	address,_,token,err := generateId.DeployGenerateId(auth, clientConnect)
	if err != nil {
		Log.Error("部署智能合约失败",err)
	}
	simulate = sim
	contractAddress = address
	instance = token
	autheration = auth

}

func NewContract(){

	//contractAddress := "0x23d67b243d501d1c890627bcde2d89e2f2bf3686"
	auth, err := bind.NewTransactor(strings.NewReader(config.Ec.EthereumAdminAccount),config.Ec.EthereumAdminPassword)
	if err != nil{
		Log.Error("解锁管理员账户失败",err)
	}
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1337000000000)}

	balance, err := clientConnect.BalanceAt(context.TODO(),auth.From,nil)
	fmt.Println("From",auth.From.String())
	fmt.Println("Balance", balance)
	sim := backends.NewSimulatedBackend(alloc, 100000000)
	token,err :=  generateId.NewGenerateId(common.HexToAddress(config.Ec.EthereumContractAddress), clientConnect)
	if err != nil {
		Log.Error("获取智能合约失败对象: ",err)
	}
	simulate = sim
	instance = token
	autheration = auth


}


func WaitPendingTrasnactions(){

	for{
		count, err := clientConnect.PendingTransactionCount(context.TODO())
		if err != nil{
			Log.Error("获取等待中的交易失败: ",err)
		}
		if count == 0{

			break
		}
		time.Sleep(time.Duration(100)*time.Millisecond)
	}

}

func WaitUtilHavePendingTrasnactions(){
	for{
		count, err := clientConnect.PendingTransactionCount(context.TODO())
		if err != nil{
			Log.Error("获取等待中的交易失败: ",err)
		}
		if count != 0{

			break
		}
		time.Sleep(time.Duration(100)*time.Millisecond)
	}
}






func SetIdToEthereum(id [32]byte) {

	//count,_ := clientConnect.PendingNonceAt(context.TODO(),common.HexToAddress(config.Ec.EthereumAdminAccount))
	//count0 ,_ := clientConnect.PendingNonceAt(context.TODO(),common.HexToAddress(config.Ec.EthereumContractAddress))
	//Log.Info("当前的Nonce值： ",count0)
	_, err := instance.SetId(&bind.TransactOpts{
		From:     autheration.From,
		Signer:   autheration.Signer,
		GasLimit: 288162,
		Value:    big.NewInt(0),
		//Nonce:	  big.NewInt(count),
	}, id)
	if err != nil{
		Log.Error("将值设置到以太坊失败: ",err)
	}
	simulate.Commit()

}

func GetIdFromEthereum() string {

	info, err := instance.Id(&bind.CallOpts{Pending: false})
	if err != nil{
		Log.Error("从以太坊中获取id值失败:",err)
	}
	return string(info[:len(info)])
}