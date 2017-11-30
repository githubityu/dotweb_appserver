package service

import (
	"dotapp_server/models"
	"github.com/henrylee2cn/faygo/errors"
)

type UserService struct {
	
}

func (*UserService) Insert(user *models.User)(newId int64,err error)  {
	newId,err = user.Insert()
	return
}
func (*UserService) Update(user *models.User) (updId int64,err error)  {
	updId,err = user.Update()
	return
}
func (*UserService) Delete(user *models.User) (delId int64,err error)  {
	delId,err = user.Delete()
	return
}

func (*UserService) getUser(userId int64) (user *models.User,err error)  {
	var ok bool
	ok,err = user.GetOne(userId)
	if !ok {
		errors.New("没有找到用户")
		return
	}
	return
}