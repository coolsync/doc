package logs_source

import (
	"os"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type LogConfig struct {
	LogPath string `json:"log_path"`
	LogName string `json:"log_name"`
	LogLevel string `json:"log_level"`
}

func LoadLogConfig()  *LogConfig{

	log_conf := LogConfig{}


	file,err := os.Open("conf/log_conf.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	byte_datas,err2 := ioutil.ReadAll(file)


	if err2 != nil {
		panic(err2)
	}

	err3 := json.Unmarshal(byte_datas,&log_conf)
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println(log_conf)

	return &log_conf
}
