package main

import (

	"../utils"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/garyburd/redigo/redis"
	"strings"
)




func getId(conn *rpc.Client) string{

	return  "1000"

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


func main(){

	utils.GetYaml()
	ethereumConn := utils.ConnectEthereum()

	redisConn := utils.ConnectRedis()
	endId := getId(ethereumConn)
	fmt.Println(redisConn)

	utils.SubscribeChannel()
	for {
		currentId := <- utils.IdChannel
		if strings.Compare(getSubId(currentId , endId),"300") == -1{
			go getId(ethereumConn)
		}

	}


}


