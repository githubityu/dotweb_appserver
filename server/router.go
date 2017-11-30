package server

import (
	"github.com/devfeel/dotweb"
	"dotapp_server/server/handler/login"
	"dotapp_server/controllers"
)

func InitRoute(server *dotweb.HttpServer) {


	dotapp := server.Group("/api")
	//dotapp.POST("/login", login.Login)
	dotapp.POST("/register", login.Register)
	dotapp.GET("/register", login.Register)
	dotapp.POST("/login", login.Login)
	dotapp.GET("/login", login.Login)
	dotapp.GET("/updatepwd", login.UpdatePwd)
	dotapp.POST("/updatepwd", login.UpdatePwd)
	dotapp.GET("/updateuser", login.UpdateUser)
	dotapp.GET("/updatenick", login.UpdateNicke)

	dotapp.POST("/uploadFile", login.FileUpload)
	dotapp.POST("/uploadFiles", login.FileUploadMore)

	dotapp.GET("/testbind", login.TestBind)
	//设置文件的实际路径 第一个参数为访问路径，第二个参数为静态文件的实际目录
	server.ServerFile("/static/*filepath", "../static")
	viewController:=new(controllers.ViewController)
	adminapp := server.Group("/admin")
	adminapp.GET("/home", viewController.Home)
	adminapp.GET("/admin", viewController.Admin)
	adminapp.GET("/button", viewController.Button)
	adminapp.GET("/login", viewController.Login)

}
