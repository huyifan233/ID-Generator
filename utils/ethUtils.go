package utils

import (
	"../conf/smartconrtact"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
	"sync"
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
//server
//const key = `{"address":"48f84bdb03d39a35b55a6c32cbef1ae960ab17eb","crypto":{"cipher":"aes-128-ctr","ciphertext":"98bfb59d728625da2e321bcd260f099bd0f34b6173d32121c611d85ab7c2c6d2","cipherparams":{"iv":"8e549798d0261c433141f6200e00f2a3"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"4e787078e7cf0504dbe3a86acdf546e779ed9ad903063cf37e32f479653c92bb"},"mac":"29f98ef4e8bb062391b5efaa89a5eb146fe3b63d9f02aef9cd2c2690d31a488e"},"id":"cfc0ec80-12b2-4e95-a883-98534e2acae3","version":3}`
//local
const key = `{"address":"12769c3419a7f491cf4e576e2e983e009d579076","crypto":{"cipher":"aes-128-ctr","ciphertext":"215430a18ab1132c6eaecdf966bc0d878a3be06cff5dce173d801afec5002db5","cipherparams":{"iv":"d41d87954da3dfca1f38e14111169fb8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d5268e70fbf8666435bf82ee53850f14486810f1944110de2aede933ae97fff1"},"mac":"a80e95fc657473f543bb989da9e8e2cde73e2b8bde52e6b05355b44a40c165ec"},"id":"5e034459-6e81-4208-9e26-c91641d20f5d","version":3}`
func ConnectEthereum(){

	if clientConnect == nil {
		mutex.Lock()
		if clientConnect == nil {
			conn, err := ethclient.Dial(config.Url.EthereumUrl)
			if err != nil{
				log.Fatal(err)
			}else{
				clientConnect = conn

			}
		}
		mutex.Unlock()
	}
}


func DeployContract(){

	auth, err := bind.NewTransactor(strings.NewReader(key),"abc")
	if err != nil{
		log.Fatal("解锁管理员账户失败",err)
	}
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(13370000000000)}
	fmt.Println("From",auth.From.String())
	fmt.Println("gasLimit",auth.GasLimit)
	fmt.Println("gasPrice",auth.GasPrice)
	fmt.Println("Value",auth.Value)
	sim := backends.NewSimulatedBackend(alloc, 100)
	address,_,token,err := generateId.DeployGenerateId(auth, clientConnect)
	if err != nil {
		log.Fatal("部署智能合约失败",err)
	}
	simulate = sim
	contractAddress = address
	instance = token
	autheration = auth

}


func increaseId(){

	_, err := instance.GetId(&bind.TransactOpts{
		From:     autheration.From,
		Signer:   autheration.Signer,
		GasLimit: 288162,
		Value:    big.NewInt(30),
	},big.NewInt(1000))
	if err!= nil{
		log.Fatal("增加Id失败",err)
	}
	simulate.Commit()

}

func GetIdFromSmartContract() string{

	increaseId()
	info, err := instance.IntId(&bind.CallOpts{Pending: true})
	if err != nil{
		log.Fatal("获取Id失败",err)
	}
	return info.String()
}