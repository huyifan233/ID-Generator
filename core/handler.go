package core

import (
	"../utils"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net/http"
)


func setHeaders(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
}

func GetId(w http.ResponseWriter, r *http.Request){

	setHeaders(w)
	id := utils.GetIdFromSmartContract()
	_, err := w.Write([]byte(id))
	if err != nil{
		fmt.Println("写返回值出错")
	}
}

func getStId(endId string, step string)(string){


	return ""
}

func getSubId(currentId string, endId string)(string){
	return "";
}

func getRedisId(conn redis.Conn)(string){

	id := utils.GetValue("generateId")

	if id == "" {
		utils.SetValue("generateId","1")
		return "1"
	}

	return id
}



