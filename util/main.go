package main

import (
	 "log"
	"fmt"
	"os"
	"bytes"
	"encoding/csv"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"dotapp_server/models"
	"sort"
)

func main() {
	//t1 := time.Now() // get current time
	//p()
	//elapsed := time.Since(t1)
	//fmt.Println("")
	//fmt.Println("爬虫结束,总共耗时: ", elapsed)
	//insertUser()
	// user:=new(models.User)
	// ok,err:=user.GetOne(1)
	// if ok {
	// 	fmt.Println("%#v",user)
	// }else{
	//	 fmt.Println("err=",err)
	// }
	//insertMessageCode("18939887456","111111",1)
	//insertMessageCode("18939887455","222222",1)
	//insertMessageCode("18939887454","333333",2)

	//insertMessageCode("18745998745","111111",2)
	//insertMessageCode("18939887454","000000",2)
	//获取验证码实例
	//path:="../static/upload/1511864526.png"
	//data:=strings.Replace(path,"../","",1)
	//fmt.Println("data=",data)

	//testmap:=make(map[int64]int64)
	//
	//testmap[0] = 0
	//testmap[1] = 1
	//testmap[2] = 2
	//testmap[3] = 3
	//
	//for _,v := range  testmap {
	//	fmt.Println("v=",v)
	//}

	qp:=make([]string,0)
	qp =append(qp,"c","a","b","d")

	sort.Strings(qp)
	fmt.Println("qb",qp)
	

}


func insertUser(){
	user := models.User{Phone:"18939559563",Portrait:"http://imgsrc.baidu.com/image/c0%3Dshijue1%2C0%2C0%2C294%2C40/sign=dce088c1a3ec8a1300175fa39f6afbfa/622762d0f703918f2b2091e45b3d269759eec42f.jpg",Nickname:"zhangsan",Aword:"不错",Password:"111111"}
	_,err :=user.InsertOrUpdate()
	if err != nil {
		fmt.Println("err=",err)
	}
}

func insertMessageCode(phone ,code string ,mtype int64){
	mc := models.MessageCode{Phone:phone,CodeTpye:mtype,Code:code}
	_,err :=mc.InsertOrUpdate()
	if err != nil {
		fmt.Println("err=",err)
	}
}




func p() {
	a:=0
	fileName := "wofang.csv"
	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)
	for i := 1; i <=202; i++ {
		fmt.Println("正在抓取第" + strconv.Itoa(i) + "页......")
		url := "http://www.wofang.com/building/p/" + strconv.Itoa(i) + "/"
		if i==1{
			url= "http://www.wofang.com/building/"
		}
		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find(".m ul li").Each(func(i int, s *goquery.Selection) {
			name:= s.Find(".title a").Text()
			location:= s.Find(".time").Text()
			price:=s.Find(".sale-price font").Text()
			if price!="" {
				a++
				s := make([]string,3)
				s[0] = name
				s[1] = price
				s[2] = location
				r2.Write(s)
				r2.Flush()
				fmt.Printf("%s,%s,%s\n", name,price, location)
			}
		})
	}
	fout,err := os.Create(fileName)
	defer fout.Close()
	if err != nil {
		fmt.Println(fileName,err)
		return
	}
	fout.WriteString(buf.String())
	fmt.Print(a)
}
