package db

import (
	"fmt"
	"dotapp_server/conf"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
)
var(
	engine  *xorm.Engine

)

func init() {
	var err error
	dataSource:= fmt.Sprintf(conf.DataSource,"root","","127.0.0.1",3306,"dotapp")+"&loc=Asia%2FShanghai"
	fmt.Println("dataSource:",dataSource)
	engine,err = xorm.NewEngine("mysql",dataSource)

	if err != nil {
		fmt.Println("初始化数据库连接失败，err:",err)
		return
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(30)
}

func GetEngine() *xorm.Engine  {
	return  engine
}