package test

import (
	"fmt"
	"dotapp_server/models"
)

type User struct {

}

type Person interface {

}

func main() {
	user := models.User{Id:1,Portrait:"http://imgsrc.baidu.com/image/c0%3Dshijue1%2C0%2C0%2C294%2C40/sign=dce088c1a3ec8a1300175fa39f6afbfa/622762d0f703918f2b2091e45b3d269759eec42f.jpg",Nickname:"zhangsan",Aword:"不错",Password:"111111"}
	user2 := models.User{Id:2,Portrait:"http://imgsrc.baidu.com/image/c0%3Dshijue1%2C0%2C0%2C294%2C40/sign=dce088c1a3ec8a1300175fa39f6afbfa/622762d0f703918f2b2091e45b3d269759eec42f.jpg",Nickname:"lisi",Aword:"hello",Password:"111111"}
	user3 := models.User{Id:3,Portrait:"http://imgsrc.baidu.com/image/c0%3Dshijue1%2C0%2C0%2C294%2C40/sign=dce088c1a3ec8a1300175fa39f6afbfa/622762d0f703918f2b2091e45b3d269759eec42f.jpg",Nickname:"wangwu",Aword:"插入成功",Password:"111111"}
	id,err:=user.InsertOrUpdate()
	if err == nil {
		fmt.Println("id=",id)
	}else {
		fmt.Println("err=",err)
	}
	id,err=user2.InsertOrUpdate()
	if err == nil {
		fmt.Println("id=",id)
	}else {
		fmt.Println("err=",err)
	}
	id,err=user3.InsertOrUpdate()
	if err == nil {
		fmt.Println("id=",id)
	}else {
		fmt.Println("err=",err)
	}
	var users []*models.User
	users,err = new(models.User).List()
	if err ==nil &&len(users)>0 {
		for _,user:= range users {
			fmt.Print("%#v",user)
		}
	}
	//user:=new(models.User)
	//user.Aword = "我是老师"
	//user.Id = 2
	//user.Update()
	//user.GetOne(2)
	//fmt.Print("%#v",user)
	//fileName := "wofang.csv"
	//fout,err := os.Create(fileName)
	//defer fout.Close()
	//if err != nil {
	//	fmt.Println(fileName,err)
	//	return
	//}
	//fmt.Println("创建是否成功=","成功")
	//user:=new(User)
	//TestSI(*user,new(User))
	//a:=5
	//TestSI(*user,a)
	//fmt.Print("成功")
	//
	//nums:= []int64{1,2,3,4}
	//numss:=[]interface{}{1,2,3,4}
	//
	//TestArgs(1,nums)
	//TestArgs(2,numss)
	//
	//TestArgs(2,numss...)
	//TestArgs2(1,nums...)
}

func TestSI(use User,obj interface{}){
}

func TestArgs(a int,args ...interface{}){
	fmt.Println(a,args)
}

func TestArgs2(a int,args ...int64){
	fmt.Println(a,args)
}