package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

//会议表
type MeetingInfo struct{
	Id int `orm:"column(id);pk,auto"`	//编号
	Meetingname string `orm:"column(meetingname)"`	//会议名称
	Meetingcontent string `orm:"column(meetingcontent);null"`	//会议内容
	Communityid int `orm:"column(communityid)"`		//所属社区编号
	Addtime string `orm:"column(addtime)"`			//创建时间
	Begintime string `orm:"column(begintime)"`		//开始时间
	Endtime string `orm:"column(endtime);null"`			//结束时间
	Meetingstate int `orm:"column(meetingstate);default(0)"`	//会议状态:0:未开始;1:进行中;2.已结束
	Nowcontent string `orm:"column(nowcontent);null"`	//最新信息:存储前端发送的Json
	Updatetime string `orm:"column(updatetime);null"`	//更新时间
	Isactive int `orm:"column(isactive);default(0)"`			//删除标记:0:正常;1:删除
}
// 自定义表名（系统自动调用）
func (u *MeetingInfo) TableName() string {
	return "tbl_meeting"
}
func init(){
	orm.RegisterModel(new(MeetingInfo))	//把LogInfo注册到orm中
}
//新增一个会议
func AddMeetingInfo(mInfo MeetingInfo)(error){
	o:=orm.NewOrm()
	o.Using("default")
	num,err:=o.Insert(&mInfo)	//向数据库中插入记录
	if err!=nil{
		fmt.Println("number="+strconv.Itoa(int(num)))
		fmt.Println("AddUserInfo Error:"+err.Error())
		return err
	}
	return nil
}

//查询会议列表
func GetMeetingInfos(strWhere string)([]MeetingInfo,error){
	var mInfos []MeetingInfo
	var strSql string="select * from tbl_meeting where 1=1 and isactive=0 "
	if strWhere!=""{
		strSql+=" and "+strWhere
	}
	strSql+=" order by id "
	o:=orm.NewOrm()
	o.Using("default")
	_,err:=o.Raw(strSql).QueryRows(&mInfos)
	if err!=nil{
		AddLog("GetMeetingInfos 错误:"+err.Error(),0)
		return mInfos,err
	}
	return mInfos,nil
}
//根据查询条件获取会议信息
func GetMeetingInfoByWhere(strWhere string)(MeetingInfo, error){
	var strSql string ="select * from tbl_meeting where 1=1 "
	if len(strWhere)>0{
		strSql+=" and "+strWhere
	}
	o:=orm.NewOrm()
	o.Using("default")
	var mInfo MeetingInfo
	err:=o.Raw(strSql).QueryRow(&mInfo)
	if err!=nil{
		AddLog("GetMeetingInfoByWhere 错误:"+err.Error(),0)
	}
	return mInfo,err
}
//获取一个会议的详细信息
func GetMeetingInfo(intID int)(MeetingInfo,error){
	mInfo:=MeetingInfo{Id:intID}
	o:=orm.NewOrm()
	o.Using("default")
	err:=o.Read(&mInfo)
	if err!=nil{
		AddLog("GetMeetingInfo 错误:"+err.Error(),0)

	}else{
		err=nil
	}
	return mInfo,err
}

//删除一个会议
func DelMeetingInfo(intID int)error{
	o:=orm.NewOrm()
	o.Using("default")
	//mInfo:=MeetingInfo{Id:intID,Isactive:1}
	//strWhere:="update tbl_meeting set isactive=1 where id="+strconv.Itoa(intID)
	//_,err:=o.Update(&mInfo)
	mInfo:=MeetingInfo{Id:intID}
	err:=o.Read(&mInfo)
	mInfo.Isactive=1
	mInfo.Updatetime=time.Now().Format("2006-01-02 15:04:05")
	_,err=o.Update(&mInfo)

	//err:=o.Raw(strWhere).QueryRow(&mInfo)
	if err!=nil{
		AddLog("DelmeetingInfo 错误:"+err.Error(),0)
	}
	return err
}

//更新一个会议信息
func UpdateMeetingInfo(mInfo MeetingInfo)error{
	o:=orm.NewOrm()
	o.Using("default")
	_,err:=o.Update(&mInfo)
	if err!=nil{
		AddLog("UpdateMeetingInfo 错误:"+err.Error(),0)
	}
	return err
}
