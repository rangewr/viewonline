package main

import (
	"github.com/astaxie/beego/orm"
	_ "viewonline/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	//************************临时测试接口使用,解决跨域问题***************************
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH,POST, GET, OPTIONS, DELETE"},
		AllowHeaders:     []string{"Origin,Content-Type,XFILENAME,XFILECATEGORY,XFILESIZE,x-requested-with"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,}))
	//********************************************************************************
	beego.Run()
}

//初始化orm,并设置数据库连接
func init(){
	orm.RegisterDriver("mysql",orm.DRMySQL)
	connectionString := "admin:Hecha1809@tcp(222.74.254.17:1366)/viewonline?charset=utf8" //外网访问端口
	//connectionString := "admin:Hecha1809@tcp(10.100.111.6:1366)/viewonline?charset=utf8" //更新部署后的内部访问端口
	orm.RegisterDataBase("default","mysql",connectionString)
	//orm.RegisterDataBase("default","mysql","admin:Hecha1809@tcp(222.74.254.17:1366/viewonline?charset=utf8")
	orm.Debug=true
	//orm.SetMaxIdleConns("default",100)
	//orm.SetMaxOpenConns("default",100)
}