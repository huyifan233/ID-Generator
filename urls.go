package main

import (
	"./core"
	"./utils"
	"log"
	"net/http"
)



func startWebService(){


	//fs := http.FileServer(http.Dir("static"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	//http.Handle("./template", http.FileServer(http.Dir("template")))

	http.HandleFunc("/ids",core.GetId)

	log.Println("Start Listen")
	err := http.ListenAndServe(":9095", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func init(){

	utils.GetYaml()
	utils.ConnectEthereum()
	//utils.ConnectRedis()
	//utils.DeployContract()
	utils.NewContract()
	utils.InitLogFactory()
	//utils.ListeningTrasactionStatus()
	utils.IdChannel = make(chan string)

	core.Id[0] = '0'

}


func main(){

	startWebService()

}
