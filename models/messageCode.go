package models

import (
	"time"
	"dotapp_server/db"
)

type MessageCode struct {
	Id int64 `json:"id"`
	Phone    string  `json:"phone" xorm:"unique varchar(11)"` //手机号
	Code    string  `json:"code"` //验证码
	CodeTpye int64   `json:"type"` //验证码类型
	MessageTime time.Time `json:"-" xorm:"updated 'time'"` //短信时间
}
func init()  {
	db.GetEngine().CreateTables(new(MessageCode))
	//db.GetEngine().Sync2(new(MessageCode))

}
func (mc *MessageCode) TableName() string {
	return "messagecode"
}

func (mc *MessageCode) Insert()(newId int64,err error){
	engine:= db.GetEngine()
	newId,err = engine.Insert(mc)
	return
}
func (mc *MessageCode) Delete()(delId int64,err error)  {
	engine:=db.GetEngine()
	delId,err = engine.Delete(mc)
	return
}
func (mc *MessageCode) Update()(updId int64,err error)  {
	engine:=db.GetEngine()
	updId,err = engine.Id(mc.Id).Cols("code").Update(mc)
	return
}

func  (mc *MessageCode) GetOne(phone string)(ok bool,err error)  {
	engine:=db.GetEngine()
	ok,err = engine.Where("phone=?",phone).Get(mc)
	return
}

//通过手机号查询判断是否存在，不存在就插入，存在就更新验证码
func (mc *MessageCode) InsertOrUpdate()(inorupdateId int64,err error)  {
	var ok bool
	mcc:=new(MessageCode)
	ok,err = mcc.GetOne(mc.Phone)
	mcc.Code = mc.Code
	if err==nil && ok {
		inorupdateId,err = mcc.Update()
		return
	}
	inorupdateId ,err = mc.Insert()
	return
}