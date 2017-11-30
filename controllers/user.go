package controllers

import (
	"strconv"
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/toolkit/encrypt"
	"dotapp_server/models"
	"fmt"
	"dotapp_server/const"
)

type UserController struct {
	BaseController
}

type UserBody struct {
	Portrait string `json:"portrait"`                     //头像
	Nickname string `json:"nickname"  xorm:"varchar(11)"` //昵称
	Password string `json:"password" form:"password"`     //密码
	Aword    string `json:"aword"`                        //一句话
	Phone    string `json:"phone" form:"phone"`           //一句话
}

func (this *UserController) InsetOrUpdate(ctx dotweb.Context) (err error) {
	var (
		newId int64
	)
	userBody := &UserBody{}
	this.DecodeJSONReq(ctx, userBody)
	if len(userBody.Phone) == 0 {
		this.Respone(ctx, 1, "手机号不能为空", nil)
		return
	}
	if len(userBody.Password) == 0 {
		this.Respone(ctx, 1, "密码不能为空", nil)
		return
	}
	password, err := encrypt.Encryption(userBody.Password, "llkj")
	if err != nil {
		this.Respone(ctx, 1, err.Error(), nil)
		return
	}
	user := &models.User{Portrait: userBody.Portrait, Nickname: userBody.Nickname,
		Password: password, Aword: userBody.Aword, Token: this.CreateToken(),
	}
	newId, err = user.InsertOrUpdate()
	if err != nil {
		this.Respone(ctx, 1, err.Error(), nil)
		return
	}
	this.Respone(ctx, 0, "请求成功", newId)
	return
}

func (this *UserController) Delete(ctx dotweb.Context) (err error) {
	var (
		id int64
	)

	uid, err := strconv.ParseInt(ctx.QueryString("uid"), 10, 64)
	if uid <= 0 || err != nil {
		this.Respone(ctx, 1, "uid获取出错,请传入数字", nil)
		return
	}
	user := &models.User{}
	ok, _ := user.GetOne(uid)
	if !ok {
		this.Respone(ctx, 1, "数据不存在", nil)
		return
	}
	id, err = user.Delete()
	if err != nil {
		this.Respone(ctx, -1, "删除失败", nil)
		return
	}
	this.Respone(ctx, 0, "删除成功", id)
	return
}

func (this *UserController) Register(ctx dotweb.Context) (err error) {
	pwd := ctx.QueryString("password")
	phone := ctx.QueryString("phone")

	if len(phone) == 0 {
		this.Respone(ctx, _const.CODE_FIAL, "手机号不能为空", "")
		return
	}
	if len(pwd) == 0 {
		this.Respone(ctx, _const.CODE_FIAL, "密码不能为空", "")
		return
	}
	user := &models.User{}
	var (
		ok   bool
		pwdd string
	)
	ok, err = user.GetByPhone(phone)
	if ok {
		this.Respone(ctx, 1, "该用户已存在", "")
		return
	} else {
		pwdd, err = encrypt.Encryption(pwd, _const.JM)
		this.DecodeJSONReq(ctx, user)
		user.Phone = phone
		user.Password = pwdd
		user.Insert()
		this.Respone(ctx, 0, "注册成功", "")
		return
	}
	return
}

func (this *UserController) Login(ctx dotweb.Context) (err error) {
	pwd := ctx.QueryString("password")
	phone := ctx.QueryString("phone")
	flatform := ctx.QueryString("flatform") //判断是app还是web

	if len(phone) == 0 {
		this.Respone(ctx, _const.CODE_FIAL, "手机号不能为空", new(models.User))
		return
	}
	if len(pwd) == 0 {
		this.Respone(ctx, _const.CODE_FIAL, "密码不能为空", new(models.User))
		return
	}
	user := &models.User{}
	var (
		ok   bool
		pwdd string
	)
	ok, err = user.GetByPhone(phone)
	if ok {
		pwdd, err = encrypt.Encryption(pwd, "llkj")
		fmt.Print("pwdd=", pwdd)
		if err != nil {
			fmt.Print("err=", err)
			this.Respone(ctx, _const.CODE_FIAL, "密码错误", new(models.User))
			return
		}
		if user.Password == pwdd {
			user.Token = this.CreateToken()
			user.Update()
			this.Respone(ctx, _const.CODE_SUCCESS, "登录成功", user)
		} else {
			this.Respone(ctx, _const.CODE_FIAL, "密码错误", new(models.User))
		}
	} else {
		if flatform == "1" {
			this.Respone(ctx, _const.CODE_FIAL, "没有找到该用户", new(models.User))
		} else {

		}
	}
	return
}

//修改用户信息
func (this *UserController) UpdateNicke(ctx dotweb.Context) (err error) {
	//
	uid := ctx.FormValue("uid")
	nickname := ctx.FormValue("nickname")
	if len(uid) == 0 {
		this.Respone(ctx, _const.CODE_FIAL, "用户id不能为空", nil)
		return
	}
	if len(nickname) == 0 {
		this.Respone(ctx, _const.CODE_FIAL, "昵称不能为空", nil)
		return
	}
	user := new(models.User)
	id, _ := strconv.ParseInt(uid, 10, 64)
	ok, err := user.GetOne(id)
	if ok {
		user.Nickname = nickname
		user.Update()
		this.Respone(ctx, _const.CODE_SUCCESS, "昵称修改成功", user)
	} else {
		this.Respone(ctx, _const.CODE_FIAL, "当前账号未注册", nil)
	}
	return
}

//修改用户信息
func (this *UserController) UpdateUser(ctx dotweb.Context) (err error) {
	//
	uid := ctx.FormValue("uid")
	content := ctx.FormValue("data")
	updatetype := ctx.FormValue("type")
	if len(uid) == 0 {
		this.Respone(ctx, _const.CODE_FIAL, "用户id不能为空", nil)
		return
	}
	if len(content) == 0 {
		this.Respone(ctx, _const.CODE_FIAL, "输入内容不能为空", nil)
		return
	}
	user := new(models.User)
	id, _ := strconv.ParseInt(uid, 10, 64)
	ok, err := user.GetOne(id)
	if ok {
		//0修改昵称。1修改头像....
		switch updatetype {
		case "0":
			user.Nickname = content
		case "1":
			user.Portrait = content
		}
		user.Update()
		this.Respone(ctx, _const.CODE_SUCCESS, "修改成功", user)
	} else {
		this.Respone(ctx, _const.CODE_FIAL, "当前账号未注册", nil)
	}
	return
}

//修改密码
func (this *UserController) UpdatePwd(ctx dotweb.Context) (err error) {
	//判断字段
	pwd := ctx.FormValue("password")     //原密码
	newPwd := ctx.FormValue("password2") //新密码
	phone := ctx.FormValue("phone")
	uid := ctx.FormValue("uid")
	if len(uid) == 0 {
		this.Respone(ctx, _const.CODE_FIAL, "用户id不能为空", nil)
		return
	}
	if len(phone) == 0 {
		this.Respone(ctx, _const.CODE_FIAL, "手机号不能为空", nil)
		return
	}
	if len(pwd) == 0 {
		this.Respone(ctx, _const.CODE_FIAL, "密码不能为空", nil)
		return
	}
	var (
		ok          bool
		pwdd, pwdd2 string //加密后的新密码
	)
	//和原密码比较
	user := new(models.User)
	id, _ := strconv.ParseInt(uid, 10, 64)
	ok, err = user.GetOne(id)
	//更改
	if ok {
		pwdd, err = encrypt.Encryption(pwd, _const.JM)
		if user.Password == pwdd {
			pwdd2, err = encrypt.Encryption(newPwd, _const.JM)
			user.Password = pwdd2
			user.Update()
			this.Respone(ctx, _const.CODE_SUCCESS, "密码修改成功", nil)
		} else {
			this.Respone(ctx, _const.CODE_FIAL, "原密码不正确", nil)
		}
	} else {
		this.Respone(ctx, _const.CODE_FIAL, "当前账号未注册", nil)
	}
	return
}
func (this *UserController) TestBind(ctx dotweb.Context) (err error) {
	ub := new(UserBody)
	this.DecodeJSONReq(ctx, ub)
	fmt.Println("%#v", ub)
	return
}
