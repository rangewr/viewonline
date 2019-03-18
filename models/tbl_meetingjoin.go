package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

//会议反馈信息表
type MeetingJoinInfo struct{
	Id int `orm:"column(id);pk,auto"`	//编号
	Meetingid int `orm:"column(meetingid)"`	//会议室编号
	Userid int `orm:"column(userid)"`		//用户编号
	Jointime string `orm:"column(jointime)"`	//参与时间
	Userexperience int `orm:"column(userexperience);default(0)"`	//用户体验结果:0:默认满意;1:满意;2:不满意
	Experiencecontent string `orm:"column(experiencecontent);null"`	//评价内容
	Votetime string `orm:"column(votetime);null"`	//评议时间
}
// 自定义表名（系统自动调用）
func (u *MeetingJoinInfo) TableName() string {
	return "tbl_meetingjoin"
}
func init(){
	orm.RegisterModel(new(MeetingJoinInfo))	//把LogInfo注册到orm中
}

//添加一个新反馈
func AddMeetingJoinInfo(mjInfo MeetingJoinInfo)error{
	if mjInfo.Meetingid<=0{
		AddLog("AddMeetingJoinInfo 传入参数错误",0)
		return errors.New("AddMeetingJoinInfo 传入参数错误")

	}
	o:=orm.NewOrm()
	o.Using("default")
	_,err:=o.Insert(&mjInfo)	//向数据库中插入记录
	if err!=nil{
		//fmt.Println("number="+strconv.Itoa(int(num)))
		AddLog("AddMeetingJoinInfo Error:"+err.Error(),0)
		return err
	}
	return nil
}
//查询反馈信息集合
func GetMeetingJoinInfos(strWhere string)([]MeetingJoinInfo,error){
	var mjInfos []MeetingJoinInfo
	//if strWhere=="" {
	//	AddLog("GetMeetingJoinInfos 传入参数错误",0)
	//	return mjInfos,errors.New("GetMeetingJoinInfos 传入参数错误")
	//}
	var strSql string ="select * from tbl_meetingjoin where 1=1 "
	if strWhere!="" {
		strSql+= " and "+strWhere
	}
	o:=orm.NewOrm()
	o.Using("default")
	_,err:=o.Raw(strSql).QueryRows(&mjInfos)
	if err!=nil{
		AddLog("GetMeetingJoinInfos Error:"+err.Error(),0)
		return mjInfos,err
	}
	return mjInfos,nil
}

//查询指定条目反馈信息
func GetMeetingJoinInfo(intID int)(MeetingJoinInfo,error){
	var mjInfo =MeetingJoinInfo{Id:intID}
	o:=orm.NewOrm()
	o.Using("default")
	err:=o.Read(&mjInfo)
	if err!=nil{
		AddLog("GetMeetingJoinInfo Error:"+err.Error(),0)
	}
	return mjInfo,err
}
//根据条件获取反馈信息
func GetMeetingJoinInfoByWhere(strWhere string)(MeetingJoinInfo,error){
	var strSql string ="select * from tbl_meetingjoin where 1=1 "
	if strWhere!="" {
		strSql+= " and "+strWhere
	}
	var mjInfo MeetingJoinInfo
	o:=orm.NewOrm()
	o.Using("default")
	err:=o.Raw(strSql).QueryRow(&mjInfo)
	if err!=nil{
		AddLog("GetMeetingJoinInfoByWhere Error:"+err.Error(),0)
	}
	return mjInfo,err
}

//更新反馈信息
func UpdateMeetingJoinInfo(mjInfo MeetingJoinInfo)error{
	o:=orm.NewOrm()
	o.Using("default")
	_,err:=o.Update(&mjInfo)
	if err!=nil{
		AddLog("UpdateMeetingJoinInfo Error:"+err.Error(),0)
	}
	return err
}