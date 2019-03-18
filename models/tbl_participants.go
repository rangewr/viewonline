package models

//与会人表
type ParticipantsInfo struct{
	Id int `orm:"column(id)"`	//编号
	Meetingid int `orm:"column(meetingid)"`	//会议编号
	Userid int `orm:"column(userid)"`		//参会用户编号

}