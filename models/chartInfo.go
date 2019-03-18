package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

//与前台交互的json结构,描述统计数据
type ChartInfo struct {
	MeetingId       int `orm:"column(meetingid);pk;auto"`			//会议编号
	Allvote 		int `orm:"column(allvote)"`		//所有投票
	Dissentingvote	  int `orm:"column(dissentingvote)"`	//反对票
	Affirmativevote	  int `orm:"column(affirmativevote)"`	//赞成票
}
// 自定义表名（系统自动调用）
func (u *ChartInfo) TableName() string {
	return "tbl_meetingjoin1"
}
func init(){
	orm.RegisterModel(new(ChartInfo))	//把LogInfo注册到orm中
}
//根据条件获取投票信息
func GetChartInfo(MeetingId int)(ChartInfo,error){
	var strSql string ="SELECT meetingid as meetingid, count(*) as allvote,count(case when userexperience>=2 then userexperience end) as dissentingvote,count(case when userexperience<=1 then userexperience end) as affirmativevote FROM tbl_meetingjoin where meetingid= "
	strSql+=strconv.Itoa( MeetingId)
	var cInfo ChartInfo
	o:=orm.NewOrm()
	o.Using("default")
	err:=o.Raw(strSql).QueryRow(&cInfo)
	if err!=nil{
		AddLog("GetMeetingJoinInfoByWhere Error:"+err.Error(),0)
	}
	return cInfo,err
}