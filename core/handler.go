package core

import (
	"../utils"
	"net/http"
	"sync"
)



func setHeaders(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
}

var Id [32]byte
var mutex sync.Mutex
var channelMutex sync.Mutex




func GetId(w http.ResponseWriter, r *http.Request){

	setHeaders(w)

	mutex.Lock()
	utils.WaitPendingTrasnactions()
	id := utils.GetIdFromEthereum()
	//if strings.Compare(string(Id[:len(Id)]),id) != 0 {
	//	Id = utils.ConvertStringTobyte32(id)
	//	utils.Log.Info("Id值被篡改，从区块链中恢复")
	//}
	validLen := utils.GetValidLen(utils.ConvertStringTobyte32(id))
	utils.Log.Info("当前的Id值: ",id)
	Id = utils.ConvertStringTobyte32(utils.Add(string(id[:validLen])))


	validLen = utils.GetValidLen(Id)
	//utils.Log.Info("有效长度: ",validLen)
	utils.Log.Info("增加后的Id值: ",string(Id[:validLen]))

	//channelMutex.Lock()
	//utils.IdChannel <- string(Id[:validLen])
	//channelMutex.Unlock()
	utils.SetIdToEthereum(Id)
	utils.WaitUtilHavePendingTrasnactions()
	mutex.Unlock()
	validLen = utils.GetValidLen(Id)
	_, err := w.Write([]byte(string(Id[:validLen])))
	if err != nil{
		utils.Log.Error("写返回值出错: ",err)
	}
}







