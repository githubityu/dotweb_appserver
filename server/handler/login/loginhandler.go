package login

import (
	"github.com/devfeel/dotweb"

	"dotapp_server/controllers"
)

func Login(ctx dotweb.Context) error {
	user := new(controllers.UserController)
	err := user.Login(ctx)
	if err != nil {
		return err
	}
	return nil
}
func Register(ctx dotweb.Context) error {
	user := new(controllers.UserController)
	err := user.Register(ctx)
	if err != nil {
		return err
	}
	return nil
}
func UpdatePwd(ctx dotweb.Context) error {
	user := new(controllers.UserController)
	err := user.UpdatePwd(ctx)
	if err != nil {
		return err
	}
	return nil
}

func FileUpload(ctx dotweb.Context) error {
	user := new(controllers.PublicFunc)
	err := user.FileUpload(ctx)
	if err != nil {
		return err
	}
	return nil
}
func FileUploadMore(ctx dotweb.Context) error {
	user := new(controllers.PublicFunc)
	err := user.FileUploadMore(ctx)
	if err != nil {
		return err
	}
	return nil
}
func UpdateUser(ctx dotweb.Context) error {
	user := new(controllers.UserController)
	err := user.UpdateUser(ctx)
	if err != nil {
		return err
	}
	return nil
}
func UpdateNicke(ctx dotweb.Context) error {
	user := new(controllers.UserController)
	err := user.UpdateNicke(ctx)
	if err != nil {
		return err
	}
	return nil
}
func TestBind(ctx dotweb.Context) error {
	user := new(controllers.UserController)
	err := user.TestBind(ctx)
	if err != nil {
		return err
	}
	return nil
}