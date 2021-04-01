package logs_source

import (
	"github.com/sirupsen/logrus"
	"os"
	"fmt"
)


// 创建一个实例
var Log = logrus.New()

func init()  {

	log_conf := LoadLogConfig()

	// 设置日志输出文件
	log_dir := log_conf.LogPath + "/" +log_conf.LogName
	fmt.Println(log_dir)
	file,err := os.OpenFile(log_dir,os.O_APPEND|os.O_WRONLY,os.ModeAppend)

	if err != nil {
		panic(err)
	}

	Log.Out = file


	// 设置日志級別

	level_mapping := map[string]logrus.Level{
		"trace":logrus.TraceLevel,
		"debug":logrus.DebugLevel,
		"info":logrus.InfoLevel,
		"warn":logrus.WarnLevel,
		"error":logrus.ErrorLevel,
		"fatal":logrus.FatalLevel,
		"panic":logrus.PanicLevel,
	}

	Log.SetLevel(level_mapping[log_conf.LogLevel])

	// 日志格式化
	Log.SetFormatter(&logrus.TextFormatter{})


}

//func main() {
//	Log.Info("xxxxx")
//}