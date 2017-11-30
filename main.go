package main

import (
	"fmt"
	"dotweb-start/core/exception"
	"dotweb-start/const"
	"github.com/devfeel/dotlog"
	"dotweb-start/util/file"
	"os"
	"flag"
	"dotweb-start/config"
	"dotapp_server/server"
)
var (
	innerLogger dotlog.Logger
	configFile  string
	configPath  string

)
func init() {
	//start log service
	configPath = file.GetCurrentDirectory()
	err := dotlog.StartLogService(configPath + "/dotlog.conf")
	if err != nil {
		os.Stdout.Write([]byte("log service start error => " + err.Error()))
	}
	innerLogger = dotlog.GetLogger(_const.LoggerName_Inner)
}
func main() {
	// defer-->finally   recover --catch
	defer func() {
		if err := recover(); err != nil {
			ex := exception.CatchError(_const.Global_ProjectName+":main", err)
			innerLogger.Error(fmt.Errorf("%v", err), ex.GetDefaultLogString())
			os.Stdout.Write([]byte(ex.GetDefaultLogString()))
		}
	}()
	//load app config
	flag.StringVar(&configFile, "config", "", "配置文件路径")
	if configFile == "" {
		configFile = configPath + "/app.conf"
	}

	//加载xml配置文件
	config.InitConfig(configFile)

	err := server.StartServer(configPath)
	if err != nil {
		innerLogger.Warn("HttpServer.StartServer失败 " + err.Error())
		fmt.Println("HttpServer.StartServer失败 " + err.Error())
	}

}
