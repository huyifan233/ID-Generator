package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"sync"
)




type Client struct {
	rpcClient *rpc.Client
	EthClient *ethclient.Client
}

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
var clientConnect *rpc.Client
var mutex sync.Mutex

func ConnectEthereum()(*rpc.Client){

	if clientConnect == nil {
		mutex.Lock()
		if clientConnect == nil {
			conn, err := rpc.Dial(config.Url.EthereumUrl)
			if err != nil{
				log.Fatal(err)
			}else{
				clientConnect = conn
			}
		}
		mutex.Unlock()
	}

	return clientConnect
}


func DeployContract(conn *rpc.Client){




}

func SubmitTransaction(conn *rpc.Client){



}

