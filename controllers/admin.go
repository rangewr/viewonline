package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"viewonline/models"
)

type GetUserList struct {
	beego.Controller
}
type AddUserInfo struct {
	beego.Controller
}
type GetUserInfo struct{
	beego.Controller
}
type UpdateUserInfo struct {
	beego.Controller
}
type DelUserInfo struct {
	beego.Controller
}
type GetCommunityList struct {
	beego.Controller
}
type AddCommunityInfo struct{
	beego.Controller
}
type GetCommunityInfo struct {
	beego.Controller
}
type UpdateCommunityInfo struct {
	beego.Controller
}
type DelCommunityInfo struct {
	beego.Controller
}
type GetMeetingList struct {
	beego.Controller
}
type AddMeetingInfo struct{
	beego.Controller
}
type GetMeetingInfo struct{
	beego.Controller
}
type UpdateMeetingInfo struct {
	beego.Controller
}
type DelMeetingInfo struct {
	beego.Controller
}
type GetMeetingJoinByMeetingID struct {
	beego.Controller
}
type AddMeetingJoin struct {
	beego.Controller
}
type GetMeetings struct {
	beego.Controller
}
type PostLiveData struct{
	beego.Controller
}
type ChageMeetingState struct{
	beego.Controller
}
type GetMeetingChart struct{
	beego.Controller
}

type GetViewers struct{
	beego.Controller
}

//获取已参会在线用户数据
func (this *GetViewers)Get(){
	var result models.ResultInfo
	var strWhere string
	//myInfo,err:=checkLogin(this.GetSession("userinfo"))
	//if  err!=nil{
	//	models.AddLog(result.ErrMsg,0)
	//	result.ErrCode=-5
	//	result.ErrMsg=err.Error()
	//	this.Data["json"]=result
	//	this.ServeJSON()
	//	return
	//}
	//if myInfo.Usertype!=0{
	//	models.AddLog("GetMeetingJoinByMeetingID Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
	//	result.ErrCode=-39
	//	result.ErrMsg="没有权限查询评议记录"
	//	this.Data["json"]=result
	//	this.ServeJSON()
	//	return
	//}
	intMeetingId,err:=this.GetInt("meetingid")

	if err!=nil ||intMeetingId<1{
		models.AddLog("GetViewers Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	strWhere="userexperience = 0 and meetingid="+strconv.Itoa(intMeetingId) +" order by jointime desc "
	mjInfo,err:=models.GetMeetingJoinInfos(strWhere)
	//mjInfo,err:=models.GetMeetingJoinInfoByWhere(strWhere)
	if err!=nil {
		models.AddLog("GetViewers Controller:查询评议记录失败",0)
		result.ErrCode=-40
		result.ErrMsg="查询评议记录失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	userInfos := make([]models.UserInfo, 0)
	for i:=0; i< len(mjInfo);i++{
		user,err1:=models.GetUserInfo(mjInfo[i].Userid)
		if err1!=nil{
			models.AddLog("GetUserInfo Controller:查询用户信息失败",0)
			result.ErrCode=-40
			result.ErrMsg="查询用户信息失败"
			this.Data["json"]=result
			this.ServeJSON()
			return
		}
		userInfos = append(userInfos, user)
	}
	info :=make(map[string]interface{})
	info["userInfo"] = userInfos
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=info
	this.Data["json"]=result
	this.ServeJSON()
}


//获取某场会议的评价信息
func (this *GetMeetingChart)Get(){
	var result models.ResultInfo
	intMeetingId,err1:=this.GetInt("MeetingId")
	if  err1!=nil ||intMeetingId<1 {
		models.AddLog("MeetingChart Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	cInfo,err:=models.GetChartInfo(intMeetingId)
	if err!=nil{
		models.AddLog("PostLiveData 获取Chart信息错误:"+err.Error(),0)
		result.ErrCode=-45
		result.ErrMsg="获取Chart信息错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=cInfo
	this.Data["json"]=result
	this.ServeJSON()
}
//变更会议状态(包括开始和结束) Post
func (this *ChageMeetingState)Post(){
	var result models.ResultInfo
	intMeetingId,err:=this.GetInt("MeetingId")
	intMeetingState,err1:=this.GetInt("MeetingState")
	if err!=nil || err1!=nil ||intMeetingId<1 || intMeetingState<0{
		models.AddLog("GetMeetingJoinByMeetingID Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	mInfo,err:=models.GetMeetingInfo(intMeetingId)
	if err!=nil{
		models.AddLog("PostLiveData 获取会议信息错误:"+err.Error(),0)
		result.ErrCode=-43
		result.ErrMsg="获取会议信息错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	//mInfo.Nowcontent=strMsg
	mInfo.Meetingstate=intMeetingState
	mInfo.Updatetime=time.Now().Format("2006-01-02 02:02:03")
	err=models.UpdateMeetingInfo(mInfo)
	if err!=nil{
		models.AddLog("PostLiveData 同步数据库错误:"+err.Error(),0)
		result.ErrCode=-44
		result.ErrMsg="同步数据库错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}
//同步会议现场数据
func (this *PostLiveData)Post(){
	var result models.ResultInfo
	intMeetingId,err:=this.GetInt("MeetingId")
	strMsg:=this.GetString("Result")
	if err!=nil ||intMeetingId<1{
		models.AddLog("GetMeetingJoinByMeetingID Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	mInfo,err:=models.GetMeetingInfo(intMeetingId)
	if err!=nil{
		models.AddLog("PostLiveData 获取会议信息错误:"+err.Error(),0)
		result.ErrCode=-42
		result.ErrMsg="获取会议信息错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	mInfo.Nowcontent=strMsg
	//mInfo.Meetingstate=1
	mInfo.Updatetime=time.Now().Format("2006-01-02 02:02:03")
	err=models.UpdateMeetingInfo(mInfo)
	if err!=nil{
		models.AddLog("PostLiveData 同步数据库错误:"+err.Error(),0)
		result.ErrCode=-42
		result.ErrMsg="同步数据库错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}
//获取直播列表
func (this *GetMeetings)Get(){
	var result models.ResultInfo
	var strWhere=" isactive=0 and DATE_FORMAT(begintime,'%Y-%m-%d')='"+time.Now().Format("2006-01-02")+"' "
	var mInfos []models.MeetingInfo
	mInfos,err:=models.GetMeetingInfos(strWhere)
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=mInfos
	this.Data["json"]=result
	this.ServeJSON()


}

//添加一条评议记录
func (this *AddMeetingJoin)Post(){
	var result models.ResultInfo
	//var strWhere string
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("GetMeetingJoinByMeetingID Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-41
		result.ErrMsg="没有权限查询评议记录"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intMeetingId,err:=this.GetInt("id")

	if err!=nil ||intMeetingId<1{
		models.AddLog("GetMeetingJoinByMeetingID Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	var mjInfo models.MeetingJoinInfo
	mjInfo.Meetingid=intMeetingId
	mjInfo.Userid=myInfo.Id
	mjInfo.Jointime=time.Now().Format("2006-01-02 15:04:05")
	mjInfo.Userexperience=0
	mjInfo.Votetime=time.Now().Format("2006-01-02 15:04:05")
	err=models.AddMeetingJoinInfo(mjInfo)
	if err!=nil {
		models.AddLog("GetMeetingJoinByMeetingID Controller:查询评议记录失败",0)
		result.ErrCode=-40
		result.ErrMsg="删除会议信息失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=mjInfo
	this.Data["json"]=result
	this.ServeJSON()
}

//根据会议ID获取评议记录 Get
func (this *GetMeetingJoinByMeetingID)Get(){
	var result models.ResultInfo
	var strWhere string
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("GetMeetingJoinByMeetingID Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-39
		result.ErrMsg="没有权限查询评议记录"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intMeetingId,err:=this.GetInt("id")

	if err!=nil ||intMeetingId<1{
		models.AddLog("GetMeetingJoinByMeetingID Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	strWhere="userexperience = 2 and meetingid="+strconv.Itoa(intMeetingId)
	mjInfo,err:=models.GetMeetingJoinInfos(strWhere)
	//mjInfo,err:=models.GetMeetingJoinInfoByWhere(strWhere)
	if err!=nil {
		models.AddLog("GetMeetingJoinByMeetingID Controller:查询评议记录失败",0)
		result.ErrCode=-40
		result.ErrMsg="查询评议记录失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	userInfo := make([]models.UserInfo, 0)
	for i:=0; i< len(mjInfo);i++{
		user,err1:=models.GetUserInfo(mjInfo[i].Userid)
		if err1!=nil{
			models.AddLog("GetUserInfo Controller:查询用户信息失败",0)
			result.ErrCode=-40
			result.ErrMsg="查询用户信息失败"
			this.Data["json"]=result
			this.ServeJSON()
			return
		}
		userInfo = append(userInfo, user)
	}
	info :=make(map[string]interface{})
	info["mjInfo"] = mjInfo
	info["userInfo"] = userInfo
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=info
	this.Data["json"]=result
	this.ServeJSON()
}
//删除会议
func (this *DelMeetingInfo)Post(){
	var result models.ResultInfo
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("DelMeetingInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-37
		result.ErrMsg="没有权限删除会议信息"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intID,err:=this.GetInt("id")
	if err!=nil ||intID<1{
		models.AddLog("DelMeetingInfo Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	err=models.DelMeetingInfo(intID)
	if err!=nil {
		models.AddLog("DelMeetingInfo Controller:删除会议信息失败",0)
		result.ErrCode=-38
		result.ErrMsg="删除会议信息失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}
//获取会议详细信息
func (this *GetMeetingInfo)Get(){
	var result models.ResultInfo
	var strWhere string
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("GetMeetingInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-33
		result.ErrMsg="没有权限查询会议信息"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intID,err:=this.GetInt("id")
	if err!=nil ||intID<1{
		models.AddLog("GetMeetingInfo Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	strWhere=" isactive=0 and id="+strconv.Itoa(intID)
	mInfo,err:=models.GetMeetingInfoByWhere(strWhere)
	if err!=nil ||mInfo.Id<1{
		models.AddLog("GetMeetingInfo Controller:获取会议信息失败",0)
		result.ErrCode=-34
		result.ErrMsg="获取会议信息失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=mInfo
	this.Data["json"]=result
	this.ServeJSON()
}
//更新会议信息
func (this *UpdateMeetingInfo)Post(){
	var result models.ResultInfo
	//var strWhere string
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("UpdateMeetingInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-35
		result.ErrMsg="没有更新查询会议信息"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intID,err2:=this.GetInt("id")
	strMeetingName:=this.GetString("meetingname")
	strMeetingContent:=this.GetString("meetingcontent")
	intCommunityId,err1:=this.GetInt("communityid")
	strBeginTime:=this.GetString("begintime")
	strEndTime:=this.GetString("endtime")

	if err1!=nil || err2!=nil||intID<1|| intCommunityId<1 || strMeetingName=="" || strBeginTime==""{
		models.AddLog("UpdateMeetingInfo Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	sourceInfo,err:=models.GetMeetingInfo(intID)
	if err!=nil{
		models.AddLog("UpdateMeetingInfo Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	sourceInfo.Id=intID
	sourceInfo.Meetingname=strMeetingName
	sourceInfo.Meetingcontent=strMeetingContent
	sourceInfo.Communityid=intCommunityId
	sourceInfo.Begintime=strBeginTime
	sourceInfo.Endtime=strEndTime
	sourceInfo.Addtime=time.Now().Format("2006-01-02 15:04:05")
	sourceInfo.Updatetime=time.Now().Format("2006-01-02 15:04:05")
	sourceInfo.Isactive=0
	//mInfo.Addtime=time.Now().Format("2006-01-02 15:04:05")
	//mInfo.Meetingstate=0
	err=models.UpdateMeetingInfo(sourceInfo)
	if err!=nil{
		models.AddLog("UpdateMeetingInfo Controller:更新会议信息失败",0)
		result.ErrCode=-36
		result.ErrMsg="更新会议信息失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}
//增加一个会议
func (this *AddMeetingInfo)Post(){
	var result models.ResultInfo
	//var strWhere string
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("AddMeetingInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-31
		result.ErrMsg="没有权限新增会议信息"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	strMeetingName:=this.GetString("meetingname")
	strMeetingContent:=this.GetString("meetingcontent")
	intCommunityId,err1:=this.GetInt("communityid")
	strBeginTime:=this.GetString("begintime")
	strEndTime:=this.GetString("endtime")
	if err1!=nil || intCommunityId<1 || strMeetingName=="" || strBeginTime==""{
		models.AddLog("AddMeetingInfo Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	var mInfo models.MeetingInfo
	mInfo.Meetingname=strMeetingName
	mInfo.Meetingcontent=strMeetingContent
	mInfo.Communityid=intCommunityId
	mInfo.Begintime=strBeginTime
	mInfo.Endtime=strEndTime
	mInfo.Isactive=0
	mInfo.Addtime=time.Now().Format("2006-01-02 15:04:05")
	mInfo.Updatetime=time.Now().Format("2006-01-02 15:04:05")
	mInfo.Meetingstate=0
	err=models.AddMeetingInfo(mInfo)
	if err!=nil{
		models.AddLog("AddMeetingInfo Controller:参数传递错误",0)
		result.ErrCode=-32
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}

	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}
//获取会议信息列表
func (this *GetMeetingList)Get(){
	var result models.ResultInfo
	var strWhere string
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("GetMeetingList Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-29
		result.ErrMsg="没有权限删除会议信息"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	strMeetingName:=this.GetString("meetingname")
	intCommunityId,err1:=this.GetInt("communityid")
	strBeginTime:=this.GetString("begintime") //查询会议开始日期
	strEndTime:=this.GetString("endtime")		//查询会议结束日期
	intMeetingState,err2:=this.GetInt("meetingstate")
	if (err1!=nil || err2!=nil || (intCommunityId<0) || (intMeetingState<(-1))) {
		models.AddLog("GetMeetingList Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if strMeetingName!=""{
		if len(strWhere)>0{
			strWhere+=" and "
		}
		strWhere+=" meetingname like '%"+strMeetingName+"%' "
	}
	if intCommunityId>0{
		if len(strWhere)>0{
			strWhere+=" and "
		}
		strWhere+=" communityid="+strconv.Itoa(intCommunityId)
	}
	if intMeetingState>=0{
		if len(strWhere)>0{
			strWhere+=" and "
		}
		strWhere+=" meetingstate="+strconv.Itoa(intMeetingState)
	}
	if strBeginTime!="" && strEndTime!=""{
		if len(strWhere)>0{
			strWhere+=" and "
		}
		strWhere+=" (DATE_FORMAT(begintime,'%Y-%m-%d')>='"+strBeginTime+"' and DATE_FORMAT(begintime,'%Y-%m-%d')<='"+strEndTime+"') "
	}
	mInfos,err:=models.GetMeetingInfos(strWhere)
	if err!=nil{
		models.AddLog("GetMeetingList Controller:获取会议列表失败",0)
		result.ErrCode=-30
		result.ErrMsg="获取会议列表失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=mInfos
	this.Data["json"]=result
	this.ServeJSON()
}
//删除社区信息
func (this *DelCommunityInfo)Post(){
	var result models.ResultInfo
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("DelCommunityInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-27
		result.ErrMsg="没有权限删除社区信息"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intId,err:=this.GetInt("id")
	if err!=nil || intId<1 {
		models.AddLog("DelCommunityInfo Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	err=models.DeleteCommunityInfo(intId)
	if err!=nil{
		models.AddLog("DelCommunityInfo Controller:删除社区["+strconv.Itoa(intId)+"]失败",0)
		result.ErrCode=-28
		result.ErrMsg="删除社区信息失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}
//更新社区信息
func (this *UpdateCommunityInfo)Post(){
	var result models.ResultInfo
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("UpdateCommunityInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-25
		result.ErrMsg="没有权限更新社区信息"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intId,err:=this.GetInt("id")
	strCommunityName:=this.GetString("communityname")
	if err!=nil || intId<1 || strCommunityName==""{
		models.AddLog("UpdateCommunityInfo Controller:参数传递错误",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	var cInfo models.CommunityInfo
	cInfo.Id=intId
	cInfo.Communityname=strCommunityName
	cInfo.Isactive=0
	err=models.UpdateCommunityInfo(cInfo)
	if err!=nil{
		models.AddLog("UpdateCommunityInfo Controller:更新社区["+strconv.Itoa(intId)+"]失败",0)
		result.ErrCode=-26
		result.ErrMsg="更新社区信息失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}
//获取社区详细信息
func (this *GetCommunityInfo)Get(){
	var result models.ResultInfo
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("GetCommunityInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-21
		result.ErrMsg="没有权限获取社区信息"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intId,err:=this.GetInt("id")
	if err!=nil || intId<1{
		models.AddLog("GetCommunityInfo Controller:社区名称为空",0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	cInfo,err:=models.GetCommunityInfo(intId)
	if err!=nil{
		models.AddLog("GetCommunityInfo Controller:获取社区["+strconv.Itoa(intId) +"]错误",0)
		result.ErrCode=-22
		result.ErrMsg="获取社区信息错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=cInfo
	this.Data["json"]=result
	this.ServeJSON()
}
//新增一个社区
func (this *AddCommunityInfo)Post(){
	var result models.ResultInfo
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("AddCommunityInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-19
		result.ErrMsg="没有权限新增社区"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	strCommunityName:=this.GetString("communityname")
	if strCommunityName==""{
		models.AddLog("AddCommunityInfo Controller:社区名称为空",0)
		result.ErrCode=-20
		result.ErrMsg="社区名称不能为空"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	var cInfo models.CommunityInfo
	cInfo.Communityname=strCommunityName
	cInfo.Isactive=0
	err=models.AddCommunityInfo(cInfo)
	if err!=nil{
		models.AddLog("AddCommunityInfo Controller:添加新社区失败",0)
		result.ErrCode=-20
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}
//获取社区列表
func (this *GetCommunityList)Get(){
	var strWhere string=" isactive=0 "
	var result models.ResultInfo
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("GetCommunityInfos Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-17
		result.ErrMsg="没有权限更新用户"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	cInfos,err:=models.GetCommunityInfosByWhere(strWhere)
	if err!=nil{
		result.ErrCode=-18
		result.ErrMsg="获取社区列表失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=cInfos
	this.Data["json"]=result
	this.ServeJSON()
}
//删除一个用户
func (this *DelUserInfo)Post(){
	var result models.ResultInfo
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("DelUserInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-23
		result.ErrMsg="没有权限删除用户"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intId,err:=this.GetInt("id")
	if err!=nil ||intId<1{
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	err=models.DelUserInfo(intId)
	if err!=nil{
		result.ErrCode=-24
		result.ErrMsg="删除用户内部错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}
//更新用户信息
func (this *UpdateUserInfo)Post(){
	var result models.ResultInfo
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("UpdateUserInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-17
		result.ErrMsg="没有权限更新用户"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intID,err4:=this.GetInt("id")
	strLoginName:=this.GetString("loginname")
	strLoginPwd:=this.GetString("loginpwd")
	intUserType,err1:=this.GetInt("usertype")
	strRealname:=this.GetString("realname")
	strMobileNo:=this.GetString("mobileno")
	strCardId:=this.GetString("cardid")
	intCommunityId,err2:=this.GetInt("communityid")
	strPersonliableName:=this.GetString("personliablename")
	intCanLoginState,err3:=this.GetInt("canloginstate")
	if err1!=nil ||err2!=nil ||err3!=nil ||err4!=nil|| strLoginName=="" || strRealname==""||strMobileNo==""||strCardId==""||strPersonliableName==""||intUserType<0||intCommunityId==0||intCanLoginState<0 {
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	var userInfo models.UserInfo
	userInfo.Loginname=strLoginName
	if strLoginPwd==""{
		userInfo.Loginpwd="" //不更新密码
	}else{
		userInfo.Loginpwd=fmt.Sprintf("%x",md5.Sum([]byte("strLoginPwd")))
	}
	userInfo.Id=intID
	userInfo.Usertype=intUserType
	userInfo.Realname=strRealname
	userInfo.Mobileno=strMobileNo
	userInfo.Cardid=strCardId
	userInfo.Communityid=intCommunityId
	userInfo.Personliablename=strPersonliableName
	userInfo.Canloginstate=intCanLoginState
	userInfo.Registertime=time.Now().Format("2006-01-02 15:04:05")
	errUp:=models.UpdateUserInfo(userInfo)
	if errUp!=nil{
		result.ErrCode=-16
		result.ErrMsg="更新用户信息错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}
//获取一个用户的详细信息
func (this *GetUserInfo)Get(){
	var result models.ResultInfo
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("AddUserInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-14
		result.ErrMsg="没有权限查看用户信息"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intUserId,err:=this.GetInt("id")
	if err!=nil ||intUserId<1{
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	userInfo,err:=models.GetUserInfo(intUserId)
	if err!=nil{
		result.ErrCode=-15
		result.ErrMsg="查看用户信息内部错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	userInfo.Loginpwd=""
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=userInfo
	this.Data["json"]=result
	this.ServeJSON()


}
//新增一个用户
func (this *AddUserInfo)Post(){
	var result models.ResultInfo
	myInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if myInfo.Usertype!=0{
		models.AddLog("AddUserInfo Controller 用户["+myInfo.Loginname+"]用户类型["+strconv.Itoa( myInfo.Usertype)+"]鉴权失败",0)
		result.ErrCode=-13
		result.ErrMsg="没有权限新增用户"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	strLoginName:=this.GetString("loginname")
	strLoginPwd:=this.GetString("loginpwd")
	intUserType,err1:=this.GetInt("usertype")
	strRealname:=this.GetString("realname")
	strMobileNo:=this.GetString("mobileno")
	strCardId:=this.GetString("cardid")
	intCommunityId,err2:=this.GetInt("communityid")
	strPersonliableName:=this.GetString("personliablename")
	intCanLoginState,err3:=this.GetInt("canloginstate")
	if err1!=nil ||err2!=nil ||err3!=nil || strLoginName=="" || strLoginPwd==""||strRealname==""||strMobileNo==""||strCardId==""||strPersonliableName==""||intUserType<0||intCommunityId==0||intCanLoginState<0 {
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	var userInfo models.UserInfo
	userInfo.Loginname=strLoginName
	userInfo.Loginpwd=fmt.Sprintf("%x",md5.Sum([]byte(strLoginPwd)))
	userInfo.Usertype=intUserType
	userInfo.Realname=strRealname
	userInfo.Mobileno=strMobileNo
	userInfo.Cardid=strCardId
	userInfo.Communityid=intCommunityId
	userInfo.Personliablename=strPersonliableName
	userInfo.Canloginstate=intCanLoginState
	userInfo.Registertime=time.Now().Format("2006-01-02 15:04:05")
	userInfo.Lastlogintime=time.Now().Format("2006-01-02 15:04:05")
	err=models.AddUserInfo(userInfo)
	if err!=nil{
		result.ErrCode=-12
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}
//获取用户列表
func (this *GetUserList)Get(){
	var result models.ResultInfo
	userInfo,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	if userInfo.Usertype!=0{ //普通用户
		models.AddLog("用户["+userInfo.Loginname+"]尝试获取用户列表:失败",0)
		result.ErrCode=-10
		result.ErrMsg="没有权限进行此操作"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	strLoginName:=this.GetString("loginname")
	intUserType,err1:=this.GetInt("usertype")
	intCommunityId,err2:=this.GetInt("communityid")
	if err1 !=nil || err2!=nil{
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	var strWhere string =" isactive=0 "
	if strLoginName!=""{
		if len(strWhere)>0{
			strWhere+=" and "
		}
		strWhere+=" loginname like '%"+strLoginName+"%' "
	}
	if intUserType>-1{
		if len(strWhere)>0{
			strWhere+=" and "
		}
		strWhere+=" usertype="+strconv.Itoa(intUserType)+" "
	}
	if intCommunityId>0{
		if len(strWhere)>0{
			strWhere+=" and "
		}
		strWhere+=" communityid="+strconv.Itoa(intCommunityId)+" "
	}

	userInfos,err:=models.GetUserInfos(strWhere)
	if err!=nil{
		result.ErrCode=-11
		result.ErrMsg="获取用户列表错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=userInfos
	this.Data["json"]=result
	this.ServeJSON()
}

