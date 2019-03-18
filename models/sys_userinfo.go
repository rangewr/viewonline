package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

//用户表
type UserInfo struct {
	Id int `orm:"column(id);pk,auto"`	//编号
	Loginname string `orm:"column(loginname)"`	//登录名
	Loginpwd string `orm:"column(loginpwd)"`	//登录口令:存储的是MD5后的密码
	Usertype int `orm:"column(usertype);default(1)"`		//用户类型:0:管理员;1:普通用户
	Realname string `orm:"column(realname)"`	//真实姓名
	Mobileno string `orm:"column(mobileno)"`	//手机号码
	Cardid string `orm:"column(cardid)"`		//身份证号码
	Communityid int `orm:"column(communityid)"`	//所属社区编号
	Personliablename string `orm:"column(personliablename)"`	//责任人名称
	Canloginstate int `orm:"column(canloginstate);default(0)"`	//是否可登录直播系统:0:可以登录;1:不可以登录
	Isactive int `orm:"column(isactive);default(0)"`		//可用标志:0:默认值,可用;1:代表已删除
	Registertime string `orm:"column(registertime)"`	//注册时间
	Lastlogintime string `orm:"column(lastlogintime);null"`	//最新登陆时间
}
// 自定义表名（系统自动调用）
func (u *UserInfo) TableName() string {
	return "sys_userinfo"
}
func init(){
	orm.RegisterModel(new(UserInfo))	//把LogInfo注册到orm中
}

//添加UserInfo 添加成功返回true
func AddUserInfo(userInfo UserInfo)error{

	if userInfo.Loginname==""{

		return errors.New("UserInfo为空")
	}
	o:=orm.NewOrm()
	o.Using("default")
	var userInfos []UserInfo
	var strSql="select * from sys_userinfo where isactive=0 and loginname='"+userInfo.Loginname+"'"
	num,err:=o.Raw(strSql).QueryRows(&userInfos)
	if err!=nil{
		fmt.Println("AddUserInfo[check exist loginname] Error:"+err.Error())
		return err
	}
	//判断用户名是否存在
	if num>0 {
		return errors.New("用户名已存在")
	}
	userInfo.Lastlogintime=time.Now().Format("2006-01-02 15:04:05")
	_,err=o.Insert(&userInfo)	//向数据库中插入记录
	if err!=nil{
		AddLog("AddUserInfo Error:"+err.Error(),0)
		return errors.New("新增用户内部错误")
	}
	return err
}
//删除一个用户
func DelUserInfo(userID int) (err error){
	o:=orm.NewOrm()
	o.Using("default")
	var strSql="update sys_userinfo set isactive=1 where id='"+strconv.Itoa(userID)+"'"
	//user:=UserInfo{Id:userID,Isactive:1}
	//num:err:=o.Update(&user)
	//num,err:=o.Update(&user)
	_,err=o.Raw(strSql).Exec()
	if err!=nil{
		AddLog("删除用户["+strconv.Itoa(userID)+"]错误:"+err.Error(),0)
		return  errors.New("删除用户错误")
	}else{
		AddLog("删除用户["+ strconv.Itoa(userID) +"]成功",0)
		return nil
	}
}
//获得用户列表
//LoginName 可按照用户登录名查询,userType 用户类型(包含普通用户和管理员)
func GetUserInfos(strWhere string) (userInfos []UserInfo,err error){

	var strSql="select * from sys_userinfo where 1=1 "
	if strWhere!=""{
		strSql+=" and "+strWhere
	}
	o:=orm.NewOrm()
	o.Using("default")
	num,err:=o.Raw(strSql).QueryRows(&userInfos)//获取记录总数

	if err!=nil{
		AddLog("GetUserInfos 出现错误,影响行数["+strconv.Itoa(int(num))+"]:"+err.Error(),0)
		return userInfos,err
	}
	return userInfos,nil
}
//获取用户详细信息,需提供用户编号ID
func GetUserInfo(intID int) (userInfo UserInfo,err error){

	if intID<=0 {
		return userInfo,errors.New("用户编号小于等于零")
	}
	var strSql string ="select * from sys_userinfo where id="+strconv.Itoa(intID)
	o:=orm.NewOrm()
	o.Using("default")
	err=o.Raw(strSql).QueryRow(&userInfo)
	if err!=nil{
		AddLog("GetUserInfo 出现错误:"+err.Error(),0)
	}
	return
}
//更新用户信息
func UpdateUserInfo(userInfo UserInfo) (error){
	o:=orm.NewOrm()
	o.Using("default")
	var uInfo UserInfo
	uInfo,err:=GetUserInfo(userInfo.Id)
	if err!=nil{
		AddLog("UpdateUserInfo 调用GetUserInfo出现错误:"+err.Error(),0)
		return errors.New("内部错误")
	}
	if userInfo.Loginpwd!=""{
		uInfo.Loginpwd=userInfo.Loginpwd
	}
	uInfo.Usertype=userInfo.Usertype
	uInfo.Realname=userInfo.Realname
	uInfo.Mobileno=userInfo.Mobileno
	uInfo.Cardid=userInfo.Cardid
	uInfo.Communityid=userInfo.Communityid
	uInfo.Personliablename=userInfo.Personliablename
	uInfo.Canloginstate=userInfo.Canloginstate

	num,err:=o.Update(&uInfo)
	if err!=nil{
		AddLog("UpdateUserInfo 更新用户信息出现错误,影响行数["+strconv.Itoa(int(num))+"]:"+err.Error(),0)
	}
	return err
}
//变更密码
func ChangePassWord(intID int,strPwd string) (bool,error){
	user:=UserInfo{Id:intID,Loginpwd:strPwd}
	o:=orm.NewOrm()
	o.Using("default")
	num,err:=o.Update(&user)
	if err!=nil{
		AddLog("ChangePassWord 变更用户密码出现错误,影响行数["+strconv.Itoa(int(num))+"]"+err.Error(),0)
		return false,err
	}
	return true,nil
}
//用户登录
func Login(strLoginName string,strLoginPwd string) (UserInfo,error){
	user:=UserInfo{Loginname:strLoginName,Loginpwd:strLoginPwd,Isactive:0}
	//user:=usr{Loginname:strLoginName,Loginpwd:strLoginPwd,Isactive:0}
	if strLoginName=="" || strLoginPwd==""{
		return user,errors.New("用户名和密码不允许为空")
	}
	o:=orm.NewOrm()
	o.Using("default")
	var strSql string ="select * from sys_userinfo where isactive=0 and loginname='"+strLoginName+"' and loginpwd='"+strLoginPwd+"' "
	err:=o.Raw(strSql).QueryRow(&user)
	//err:=o.Read(&user) //这种方法只适合主键查询
	if err!=nil{
		AddLog("Login 方法出现错误:"+err.Error(),0)
	}
	if user.Id==0 {
		AddLog("用户["+strLoginName+"]登录失败",1)
		return user,errors.New("用户名或口令错误")
	}else{
		AddLog("用户["+strLoginName+"]登录成功!",1)
		return user,nil
	}
}