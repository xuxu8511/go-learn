package global

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
	"zinx-demo/config"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	G_ZinxConfig config.ServerConfig
)

func init() {
	v := viper.New()
	v.SetConfigFile("resources/server_config.yaml")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := v.Unmarshal(&G_ZinxConfig); err != nil {
		log.Error(err)
	}

	logger := &lumberjack.Logger{
		Filename:   G_ZinxConfig.Logrus.FileName,
		MaxSize:    G_ZinxConfig.Logrus.MaxSize,  // 日志文件大小，单位是 MB
		MaxBackups: G_ZinxConfig.Logrus.MaxCount, // 最大过期日志保留个数
		MaxAge:     28,                           // 保留过期文件最大时间，单位天
		//Compress:   true,                         // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
	}

	writers := []io.Writer{
		logger,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	log.SetReportCaller(true)
	log.SetFormatter(&LogrusFormatter{})
	log.SetOutput(fileAndStdoutWriter)
}
