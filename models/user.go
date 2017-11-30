package models

import (
	"dotapp_server/db"
	"time"
	"fmt"
)
type User struct {
	Id int64 `json:"id"`
	Portrait string `json:"portrait"` //头像
	Nickname string `json:"nickename" xorm:"varchar(11)"` //昵称
	Password string `json:"-"` //密码
	Aword    string  `json:"aword"` //一句话
	Phone    string  `json:"phone"` //手机号
	Token    string  `json:"token"` //token
	Createtime time.Time  `json:"-" xorm:"created 'createtime'"` //注册时间
}

func init()  {
	//db.GetEngine().CreateTables(new(User))
	err := db.GetEngine().Sync2(new(User))
	if err != nil {
		fmt.Println("err=",err)
	}
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) Insert()(newId int64,err error){
	engine:= db.GetEngine()
	newId,err = engine.Insert(user)
	return
}
func (user *User) List()(users []*User,err error){
	 engine:=db.GetEngine()
	 err = engine.Find(&users)
	 return
}

func (user *User) Delete()(delId int64,err error)  {
	engine:=db.GetEngine()
	delId,err = engine.Delete(user)
	return
}

func (user *User) Update()(updId int64,err error)  {
	engine:=db.GetEngine()
	updId,err = engine.Id(user.Id).Update(user)
	return
}

/**
 ok  false 未找到  true 找到
 */
func  (user *User) GetOne(id int64)(ok bool,err error)  {
	engine:=db.GetEngine()
	ok,err = engine.Id(id).Get(user)
	return
}
func  (user *User) GetByPhone(phone string)(ok bool,err error)  {
	engine:=db.GetEngine()
	ok,err = engine.Where("phone=?",phone).Get(user)
	return
}
func  (user *User) Get()(ok bool,err error)  {
	engine:=db.GetEngine()
	ok,err = engine.Get(user)
	return
}
func  (user *User) GetOne2(aword string)(ok bool,err error)  {
	engine:=db.GetEngine()
	ok,err = engine.Where("aword=?",aword).Get(user)
	return
}
func (user *User) InsertOrUpdate()(inorupdateId int64,err error)  {
	var ok bool
	if user.Id<=0 {
		inorupdateId,err  = user.Insert()
		return
	}
	ok,err = new(User).GetOne(user.Id)
	 if err==nil && ok {
	 	inorupdateId,err = user.Update()
	 	return
	 }
	 inorupdateId ,err = user.Insert()
	return
}


