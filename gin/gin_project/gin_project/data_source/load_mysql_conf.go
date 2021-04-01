package data_source

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type MysqlConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	LogoMode bool `json:"logo_mode"`
}

func LoadMysqlConf() *MysqlConf {

	mysql_conf := MysqlConf{}

	file,err := os.Open("conf/mysql_conf.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	byte_data,err2 := ioutil.ReadAll(file)

	if err2 != nil {
		panic(err2)
	}

	err3 := json.Unmarshal(byte_data,&mysql_conf)

	if err3 != nil {
		panic(err3)
	}


	return &mysql_conf

}

