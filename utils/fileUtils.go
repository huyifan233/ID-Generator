package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	IdGenerator IDGen `yaml:"id-generator"`
	Url UrlConfig `yaml:"url-config"`
	Ec EthereumConfig `yaml:"ethereum-config"`
}

type IDGen struct {
	Step string `yaml:"step"`
}

type UrlConfig struct {
	RedisUrl string `yaml:"redis-url"`
	EthereumUrl string `yaml:"ethereum-url"`
}

type EthereumConfig struct {
	EthereumAdminAccount string `yaml:"ethereum-admin-account"`
	EthereumAdminAddress string `yaml:"ethereum-admin-address"`
	EthereumAdminPassword string `yaml:"ethereum-admin-password"`
	EthereumIdStep string `yaml:"ethereum-id-step"`
	EthereumContractAddress string `yaml:"ethereum-contract-address"`
}

var config Config

func GetYaml()(){

	//execpath, err := os.Executable()
	//configFilePath := filepath.Join(execpath,"../conf/conf.yaml")
	abspath, _ := filepath.Abs("../ID-Generator/conf/conf.yaml")
	//abspath := "C:\\Users\\huyifan01\\Documents\\ID-Generator\\conf\\conf.yaml"
	yamlFile, err := ioutil.ReadFile(abspath)
	if err != nil{
		Log.Error("获取日志文件出错",err)
	}

	yaml.Unmarshal(yamlFile,&config)


}