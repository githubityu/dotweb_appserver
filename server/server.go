package server

import (
	"fmt"
	"github.com/devfeel/dotlog"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/config"
	"strconv"
	"dotweb-start/global"
	"dotweb-start/const"
)

func StartServer(configPath string) error {
	//初始化DotServer
	global.DotApp = dotweb.New()
	global.DotApp.SetDevelopmentMode()
	appConfig := config.MustInitConfig(configPath + "/dotweb.conf")
	global.DotApp.SetConfig(appConfig)

	//设置路由
	InitRoute(global.DotApp.HttpServer)

	//设置模板目录
	global.DotApp.HttpServer.Renderer().SetTemplatePath("../static")

	innerLogger := dotlog.GetLogger(_const.LoggerName_Inner)

	innerLogger.Debug("dotweb.StartServer => " + strconv.Itoa(appConfig.Server.Port))
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(appConfig.Server.Port))
	err := global.DotApp.Start()
	return err
}
