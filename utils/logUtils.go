package utils

import (
	"fmt"
	"github.com/op/go-logging"
	"os"
)

var Log = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

//type Password string
//
//func (p Password) Redacted() interface{} {
//	return logging.Redact(string(p))
//}


func InitLogFactory(){

	logFile, err := os.OpenFile("../logs/id_generate_log.log", os.O_WRONLY,0666)
	if err != nil{
		fmt.Println(err)
	}
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.INFO, "")

	logging.SetBackend(backend1Leveled, backend2Formatter)


}

func main() {

	
}
