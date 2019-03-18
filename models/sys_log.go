package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

//日志表
type LogInfo struct {
	Id int `orm:"column(id)"`	//编号
	Addtime string `orm:"column(addtime)"`	//记录时间
	Logtype int `orm:"column(logtype)"`		//日志类型:0:系统日志;1:用户日志
	Logcontent string `orm:"column(logcontent)"`	//日志内容
}
// 自定义表名（系统自动调用）
func (u *LogInfo) TableName() string {
	return "sys_log"
}
func init(){
	orm.RegisterModel(new(LogInfo))	//把LogInfo注册到orm中
	//配置log输出选项
	beego.SetLogger("file", `{"filename":"logs/logInfo.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()

}
//获取Log信息,strBegintime日志开始时间;strEndtime日志结束时间,logType:日志类型,pageNo:需要的页码编号;pageSize:每页大小,返回的记录总数需要计算
func GetLogInfos(strBegintime string,strEndtime string,intLogType int,intPageNo int,intPageSize int) (Logs []LogInfo,totalItem int,totalpages int,rtnErr error ){
	var strSql string
	strSql="select * from sys_log where 1=1 "
	if !(strBegintime=="" || strEndtime==""){
		strSql+=" addtime >="+strBegintime +" and addtime<=" + strEndtime
	}
	if intLogType>=0 && intLogType <=1{
		strSql+=" logtype="+strconv.Itoa(intLogType)
	}
	strSql+=" order by addtime desc "
	strSqlCount:=strSql

	if(intPageNo>=1){
		if intPageSize>0 {
			var intBegin int =(intPageNo-1)*intPageSize
			var intEnd int =intPageNo*intPageSize
			strSql+="limit "+strconv.Itoa(intBegin)+","+strconv.Itoa(intEnd)
		}
	}

	o:=orm.NewOrm()
	o.Using("default")
	//o.Begin()
	//var Logs []LogInfo
	totalItem=0
	totalpages =0
	o.Raw(strSqlCount).QueryRow(&totalItem)//获取记录总数
	if totalItem <= intPageSize {
		totalpages = 1
	} else if totalItem > intPageSize {
		temp := totalItem / intPageSize
		if (totalItem % intPageSize) != 0 {
			temp = temp + 1
		}
		totalpages = temp
	}



	_,err:=o.Raw(strSql).QueryRows(&Logs)
	if err!=nil{
		fmt.Println("GetLogInfos Error:"+err.Error())
		return Logs,totalItem,totalpages,err
	}
	return Logs,totalItem,totalpages,nil
}

//根据ID获取LogInfo
func GetLogInfobyID(intID int) (loginfo LogInfo,err error){
	if(intID<=0){

		return loginfo,errors.New("编号小于零")
	}
	o:=orm.NewOrm()
	o.Using("default")
	var strSql string
	strSql="select * from sys_log where id="+strconv.Itoa(intID)
	var logs []LogInfo
	num,err:=o.Raw(strSql).QueryRows(&logs)
	fmt.Println("GetLogInfobyID Error:"+strconv.Itoa(int(num)))
	if err!=nil{
		fmt.Println("GetLogInfos Error:"+err.Error())
		return loginfo,err
	}
	if len(logs) >0{
		return logs[0],nil
	}else{
		return loginfo,err
	}
}

//添加Log 添加成功返回true
func addLogInfo(logInfo LogInfo) (bolResult bool,err error){

	if logInfo.Logcontent==""{
		return false,errors.New("LogInfo为空")
	}
	o:=orm.NewOrm()
	o.Using("default")

	num,err:=o.Insert(&logInfo)
	if err!=nil{
		fmt.Println("number="+strconv.Itoa(int(num)))
		fmt.Println("addLogInfo Error:"+err.Error())
		return false,err
	}
	logs.Debug(logInfo.Logcontent)
	return true,nil
}
//对外公开方法,添加一行个Log
func AddLog(strContent string,intLogType int){
	var logInfo LogInfo
	logInfo.Logcontent=strContent
	logInfo.Addtime=time.Now().Format("2006-01-02 15:04:05")
	logInfo.Logtype=intLogType
	addLogInfo(logInfo)

}