package core

import (
	"../utils"
	"fmt"
	"net/http"
	"strings"
)



func setHeaders(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
}

var Id [32]byte

func GetId(w http.ResponseWriter, r *http.Request){

	setHeaders(w)
	utils.WaitPendingTrasnactions()
	id := utils.GetIdFromEthereum()
	if strings.Compare(string(Id[:len(Id)]),id) != 0 {
		fmt.Println("Id 值被篡改，从区块链中恢复")
		Id = utils.ConvertStringTobyte32(id)
	}

	Id = utils.ConvertStringTobyte32(utils.Add(string(Id[:len(Id)]),"1000"))

	go utils.SetIdToEthereum(Id)

	_, err := w.Write([]byte(string(Id[:len(Id)])))
	if err != nil{
		fmt.Println("写返回值出错")
	}
}







