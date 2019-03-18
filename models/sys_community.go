package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

//社区表
type CommunityInfo struct {
	Id            int `orm:"column(id);pk,auto"`			//社区编号
	Communityname string `orm:"column(communityname)"`		//社区名称
	Isactive	  int `orm:"column(isactive);default(0)"`	//删除标记
}
// 自定义表名（系统自动调用）
func (u *CommunityInfo) TableName() string {
	return "sys_community"
}
func init(){
	orm.RegisterModel(new(CommunityInfo))	//把LogInfo注册到orm中
}
//根据条件查询社区名
func GetCommunityInfoByWhere(strWhere string)(CommunityInfo,error){
	var strSql string ="select * from sys_community where 1=1 "
	var cInfo CommunityInfo
	if strWhere!=""{
		strSql+=" and "+strWhere
	}
	strSql+="order by id "

	o:=orm.NewOrm()
	o.Using("default")
	err:=o.Raw(strSql).QueryRow(&cInfo)
	if err!=nil{
		AddLog("GetCommunityInfoByWhere Error:"+err.Error(),0)
		err=nil
	}
	return cInfo,err
}
//根据条件查询社区列表
func GetCommunityInfosByWhere(strWhere string)([]CommunityInfo,error){
	var strSql string ="select * from sys_community where 1=1 "
	var cInfos []CommunityInfo
	if strWhere!=""{
		strSql+=" and "+strWhere
	}
	strSql+="order by id "

	o:=orm.NewOrm()
	o.Using("default")
	_,err:=o.Raw(strSql).QueryRows(&cInfos)
	if err!=nil{
		AddLog("GetCommunityInfoByWhere Error:"+err.Error(),0)
	}
	return cInfos,err
}
func GetCommunityInfo(intID int)(CommunityInfo,error){
	var cInfo CommunityInfo
	o:=orm.NewOrm()
	o.Using("default")
	err:=o.Read(&cInfo)
	if err!=nil{
		AddLog("GetCommunityInfo Error:"+err.Error(),0)
	}
	return cInfo,err
}
//新增社区名
func AddCommunityInfo(cInfo CommunityInfo)error{
	//检查社区名称重复
	ccInfo,err:=GetCommunityInfoByWhere(" communityname='"+cInfo.Communityname+"'")
	if err!=nil{
		AddLog("AddCommunityInfo Error:"+err.Error(),0)
		return errors.New("社区名称重复请重新输入")
	}
	if ccInfo.Id>0{
		AddLog("AddCommunityInfo Error:社区名称重复["+cInfo.Communityname+"]",0)
		return errors.New("社区名称重复请重新输入")
	}
	o:=orm.NewOrm()
	o.Using("default")
	_,err=o.Insert(&cInfo)
	if err!=nil{
		AddLog("AddCommunityInfo 插入错误:"+err.Error(),0)
	}
	return err
}
//更新社区信息
func UpdateCommunityInfo(cInfo CommunityInfo)error{
	o:=orm.NewOrm()
	o.Using("default")
	_,err:=o.Update(&cInfo)
	if err!=nil{
		AddLog("UpdateCommunityInfo Error:"+err.Error(),0)
	}
	return err
}
//删除一条社区信息
func DeleteCommunityInfo(intID int)error{
	cInfo:=CommunityInfo{Id:intID}
	o:=orm.NewOrm()
	o.Using("default")
	err:=o.Read(&cInfo)
	if err!=nil{
		AddLog("DeleteCommunityInfo read Error:"+err.Error(),0)
	}
	cInfo.Isactive=1
	_,err=o.Update(&cInfo)
	if err!=nil{
		AddLog("DeleteCommunityInfo Error:"+err.Error(),0)
	}
	return err
}
