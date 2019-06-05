package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	IdGenerator IDGen `yaml:"id-generator"`
	Url UrlConfig `yaml:"url-config"`
}

type IDGen struct {
	Step string `yaml:"step"`
}

type UrlConfig struct {
	RedisUrl string `yaml:"redis-url"`
	EthereumUrl string `yaml:"ethereum-url"`
}

var config Config

func GetYaml()(){

	//abspath,_ := filepath.Abs("conf/conf.yaml")
	yamlFile, err := ioutil.ReadFile("C:\\Users\\huyifan01\\Documents\\ID-Generator\\conf\\conf.yaml")
	if err != nil{
		log.Fatal("获取日志文件出错",err)
	}

	yaml.Unmarshal(yamlFile,&config)


}