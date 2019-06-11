package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var client redis.Conn
//var IdChannel chan string

func ConnectRedis(){


	if client == nil {
		mutex.Lock()
		if client == nil{
			fmt.Println( config.Url.RedisUrl)
			conn, err := redis.Dial("tcp", config.Url.RedisUrl)
			if err != nil{
				Log.Error("连接redis失败",err)
			}else{
				client = conn
				//IdChannel = make(chan string)
			}
		}
		mutex.Unlock()
	}


}

func GetValue(key string)(string){
	result, err := redis.String(client.Do("get",key))

	if err != nil {
		Log.Error("获取key失败",err)
	}
	return result
}

func SetValue(key string, value string){
	_, err := client.Do("set",key, value)
	if err != nil{
		Log.Error("插入数据到redis失败",err)
	}
	_,err2 := client.Do("PUBLISH","idChannel",value)
	if err2 != nil {
		Log.Error("发布数据发布数据失败")
	}
}

//func SubscribeChannel(){
//
//	//err := client.Send("SUBSCRIBE","idChannel")
//	client.Do("auth","Abcd1234")
//	psc := redis.PubSubConn{client}
//	err := psc.Subscribe("idChannel")
//	if err != nil{
//		Log.Error("订阅Channel失败")
//	}
//
//	for{
//		switch reply := psc.Receive().(type) {
//		case redis.Message:
//			IdChannel <- string(reply.Data)
//		case redis.Subscription:
//			fmt.Printf("%s: %s %d\n", reply.Channel, reply.Kind, reply.Count)
//		case error:
//			fmt.Println(reply)
//		}
//
//	}
//}

func PublishValue(value string){

	client.Do("PUBLISH","idChannel",value)

}




