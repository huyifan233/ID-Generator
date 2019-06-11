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

	fmt.Println("ID1: ",Id)
	setHeaders(w)
	utils.WaitPendingTrasnactions()
	id := utils.GetIdFromEthereum()
	fmt.Println("id: ",id)
	fmt.Println("比较值：",strings.Compare(string(Id[:len(Id)]),id))
	if strings.Compare(string(Id[:len(Id)]),id) != 0 {
		fmt.Println("Id 值被篡改，从区块链中恢复")
		Id = utils.ConvertStringTobyte32(id)
	}
	validLen := utils.GetValidLen(Id)
	fmt.Println(validLen)
	Id = utils.ConvertStringTobyte32(utils.Add(string(Id[:validLen]),"1000"))

	go utils.SetIdToEthereum(Id)

	validLen = utils.GetValidLen(Id)
	_, err := w.Write([]byte(string(Id[:validLen])))
	if err != nil{
		fmt.Println("写返回值出错")
	}
}







