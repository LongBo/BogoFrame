package config

import (
	"encoding/json"
	"io/ioutil"
	"ncfwxen/common/utils"
	"os"
)

type SysConf struct {
	Domain   string `json:"WebDomain"`
	HttpPort string `json:"WebPort"`
}

type DbConf struct {
	Host string `json:"Host"`
	Name string `json:"DbName"`
	Port string `json:"Port"`
	User string `json:"UserName"`
	Pwd  string `json:"Password"`
}

func readConfig() (data []byte) {
	args := os.Args
	if args == nil || len(args) < 2 {
		utils.Nlog.Panic("config file not exist.")
	}
	file := args[1]
	data, err := ioutil.ReadFile("./" + file)
	if err != nil {
		utils.Nlog.Panic("read config file failed.")
	}
	return
}

func DbConfig() *DbConf {
	data := readConfig()
	conf := DbConf{}
	if err := json.Unmarshal(data, &conf); err != nil {
		utils.Nlog.Panic("json config file is error.")
	}
	return &conf
}

func SysConfig() *SysConf {
	data := readConfig()
	conf := SysConf{}
	if err := json.Unmarshal(data, &conf); err != nil {
		utils.Nlog.Panic("json config file is error.")
	}
	return &conf
}
