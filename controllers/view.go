package controllers

import "github.com/devfeel/dotweb"

type ViewController struct{}

// Home 首页
func (*ViewController) Home(ctx dotweb.Context) (err error) {
	return ctx.View("views/index.html")
}
func (*ViewController) Admin(ctx dotweb.Context) (err error) {
	return ctx.View("views/admin.html")
}
func (*ViewController) Button(ctx dotweb.Context) (err error) {
	return ctx.View("views/button.html")
}
func (*ViewController) Login(ctx dotweb.Context) (err error) {
	return ctx.View("assets/page/login/login.html")
}
