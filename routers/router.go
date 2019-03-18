package routers

import (
	"viewonline/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})//测试用
	//=============前台所需页面===============
	beego.Router("/proc/Login", &controllers.UserLogin{})//登录方法 Post
	beego.Router("/proc/Logout", &controllers.UserLogout{})//登出方法 Get
	beego.Router("/proc/GetMyMeetingInfo", &controllers.GetMyMeetingInfo{})//获取我的会议信息列表 Get
	beego.Router("/proc/JoinMeeting", &controllers.JoinMeeting{})//参加一个会议 Get
	beego.Router("/proc/RefreshMeeting", &controllers.RefreshMeeting{})//刷新一个会议状态 Get
	beego.Router("/proc/VoteMeeting", &controllers.VoteMeeting{})//对参加的会议进行打分 Post
	//==============后台所需页面===============
	beego.Router("/admin/GetUserList", &controllers.GetUserList{})//获取用户列表 Get
	beego.Router("/admin/AddUserInfo", &controllers.AddUserInfo{})//添加一个用户 Post
	beego.Router("/admin/GetUserInfo", &controllers.GetUserInfo{})//查询一个用户的详细信息 Get
	beego.Router("/admin/DelUserInfo", &controllers.DelUserInfo{})//删除一个用户的详细信息 Post
	beego.Router("/admin/UpdateUserInfo", &controllers.UpdateUserInfo{})//更新一个用户的详细信息 Post

	beego.Router("/admin/GetCommunityList", &controllers.GetCommunityList{})//获取社区列表 Get
	beego.Router("/admin/AddCommunityInfo", &controllers.AddCommunityInfo{})//添加一个新社区 Post
	beego.Router("/admin/GetCommunityInfo", &controllers.GetCommunityInfo{})//查询社区详细信息 Get
	beego.Router("/admin/UpdateCommunityInfo", &controllers.UpdateCommunityInfo{})//更新社区详细信息 Post
	beego.Router("/admin/DelCommunityInfo", &controllers.DelCommunityInfo{})//删除社区详细信息 Post

	beego.Router("/admin/GetMeetingList", &controllers.GetMeetingList{})//获取会议列表 Get
	beego.Router("/admin/AddMeetingInfo", &controllers.AddMeetingInfo{})//添加一个新会议 Post
	beego.Router("/admin/GetMeetingInfo", &controllers.GetMeetingInfo{})//查询会议详细信息 Get
	beego.Router("/admin/UpdateMeetingInfo", &controllers.UpdateMeetingInfo{})//更新会议详细信息 Post
	beego.Router("/admin/DelMeetingInfo", &controllers.DelMeetingInfo{})//删除会议详细信息 Post

	beego.Router("/admin/GetMeetingJoinByMeetingID", &controllers.GetMeetingJoinByMeetingID{})//根据会议ID获取评议记录 Get
	beego.Router("/admin/AddMeetingJoin", &controllers.AddMeetingJoin{})//添加一个新评议 Post
	beego.Router("/admin/MeetingChart", &controllers.GetMeetingChart{})//获取某场会议的统计信息 Get
	//beego.Router("/admin/UpdateMeetingJoin", &controllers.UpdateMeetingJoin{})//更新评议详细信息 Post
	//beego.Router("/admin/DelMeetingJoin", &controllers.DelMeetingJoin{})//删除评议详细信息 Post

	//======================直播接口================================

	beego.Router("/Live/GetMeetings", &controllers.GetMeetings{})//添加一个新评议 Get
	beego.Router("/Live/PostLiveData", &controllers.PostLiveData{})//同步会议现场数据 Post
	beego.Router("/Live/ChageMeetingState", &controllers.ChageMeetingState{})//变更会议状态(包括开始和结束) Post

	//==============为前端提供已参会在线用户的数据===================
	beego.Router("/Live/GetViewers", &controllers.GetViewers{})//获取已参会在线用户数据 Get
}
