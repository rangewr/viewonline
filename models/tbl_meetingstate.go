package models
//实时会议状态表
type MeetingState struct{
	Meetingroomid int `orm:"column(meetingroomid)"`	//会议室编号
	Nowcontent string `orm:"column(nowcontent)"`	//最新信息
	Meetingstate int `orm:"column(meetingstate)"`	//会议状态:0:未开始;1:进行中;2.结束
	Updatetime string `orm:"column(updatetime)"`	//更新时间
}
