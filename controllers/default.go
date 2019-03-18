package controllers

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"viewonline/models"
)

type MainController struct {
	beego.Controller
}
//测试页面/test
type TestController struct {
	beego.Controller
}

type UserLogin struct{
	beego.Controller
}
type UserLogout struct {
	beego.Controller
}
type GetMyMeetingInfo struct {
	beego.Controller
}
type JoinMeeting struct{
	beego.Controller
}
type RefreshMeeting struct{
	beego.Controller
}
type VoteMeeting struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "../views/login.tpl"
}
//测试页面具体方法
func (c *TestController) Get(){
	//models.AddLog("测试Log内容",0)
	//var userInfo models.UserInfo
	//userInfo.Loginname="qiutian"
	//userInfo.Loginpwd=strings.ToUpper(fmt.Sprintf("%x",md5.Sum([]byte("111111"))))//生成MD5,所有字母大写
	//userInfo.Realname="张三"
	//userInfo.Registertime=time.Now().Format("2006-01-02 15:04:05")
	//userInfo.Canloginstate=0
	//userInfo.Cardid="123456"
	//userInfo.Communityid=0
	//userInfo.Isactive=0
	//userInfo.Lastlogintime=time.Now().Format("2006-01-02 15:04:05")
	//err:=models.AddUserInfo(userInfo)
	//var data string
	//if err!=nil{
	//	data=err.Error()
	//}else{
	//	data=strconv.FormatBool(bolReturn)
	//}

	//LogInfos,intCount,err:=models.GetLogInfos("","",-1,0,10)
	//if err!=nil{
	//	fmt.Println("TestController Error:"+err.Error())
	//}
	//data, err := json.Marshal(LogInfos)//返回json格式
	//if err != nil {
	//	//log.Println(err)
	//	fmt.Println("change json Error:"+err.Error())
	//	return
	//}

	//fmt.Println("LogInfos="+string(data))
	//c.Ctx.WriteString(string(data))
	//fmt.Println(data)
}
//处理用户登录
func (this *UserLogin)Post(){

	strLoginName:=this.GetString("loginname")
	strLoginPwd:=this.GetString("loginpwd")
	var result models.ResultInfo
	if strLoginName=="" || strLoginPwd==""{
		result.ErrCode=-1
		result.ErrMsg="用户名和口令不能为空"
		//data,err:=json.Marshal(result)
		this.Data["json"]=result
		this.ServeJSON()
		return
	}

	userinfo,err:=models.Login(strLoginName,fmt.Sprintf("%x",md5.Sum([]byte("strLoginPwd"))))
	if err!=nil{
		result.ErrCode=-2
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	userinfo.Loginpwd=""
	if userinfo.Id>0 {
		result.ErrCode=0
		result.ErrMsg="登录成功"
		this.SetSession("userinfo",userinfo)	//加入session管理
	}else{
		result.ErrCode=-3
		result.ErrMsg="登录失败,内部错误"
	}
	result.Result=userinfo
	this.Data["json"]=result
	this.ServeJSON()
}
//退出登录
func (this *UserLogout) Get(){
	this.SetSession("userinfo",nil)
	var result models.ResultInfo
	result.ErrCode=0
	result.ErrMsg="退出登录成功"
	this.Data["json"]=result
	this.ServeJSON()
}
//获取我的会议列表(正在直播和未开始直播的)
func (this *GetMyMeetingInfo) Get(){
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
	//fmt.Println("GetMyMeetingInfo usertype="+strconv.Itoa(userInfo.Usertype))
	result.ErrCode=0
	result.ErrMsg=""
	result.Result=userInfo

	//获取用户可参加的当天的未开始和正在直播的会议
	var strSql string=" communityid="+strconv.Itoa(userInfo.Communityid)+" and meetingstate < 2 and begintime>=DATE_FORMAT(NOW(),'%Y-%m-%d') "
	mInfos,err:=models.GetMeetingInfos(strSql)
	if err!=nil{
		models.AddLog("GetMyMeetingInfo Controllers Error:"+err.Error(),0)
		result.ErrCode=-6
		result.ErrMsg="查询会议信息失败,请刷新页面或重新登录"
		result.Result=""
	}else{
		result.ErrMsg=""
		result.ErrCode=0
		result.Result=mInfos
	}
	this.Data["json"]=result
	this.ServeJSON()
}
//加入一场会议
func (this *JoinMeeting)Get(){
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
	intMeetingID,err:=this.GetInt("id")
	if err!=nil{
		models.AddLog("JoinMeeting Controller Error:"+err.Error(),0)
		result.ErrCode=-7
		result.ErrMsg="接收参数失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}

	//获取会议id,通过用户所在社区和会议社区来判定用户是否有权参加(管理员有权参加所有会议)
	//如果相符则返回会议所有内容
	mInfo,err:=models.GetMeetingInfo(intMeetingID)
	if err!=nil{
		models.AddLog("JoinMeeting Controller GetMeetingInfo Error:"+err.Error(),0)
		result.ErrCode=-8
		result.ErrMsg="获取会议信息出现错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	var bolCheck=false
	if userInfo.Usertype==0{
		if mInfo.Communityid==userInfo.Communityid{
			bolCheck=true
		}
	}else{
		bolCheck=true
	}
	if !bolCheck {
		result.ErrCode=-9
		result.ErrMsg="你没有权限加入这场会议"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	//如果会议状态为:0:未开始;1:进行中 ,则在meetingjoin表中插入新纪录
	var mjInfo models.MeetingJoinInfo
	mjInfo.Meetingid=intMeetingID
	mjInfo.Userid=userInfo.Id
	mjInfo.Jointime=time.Now().Format("2006-01-02 15:04:05")
	mjInfo.Userexperience=0
	mjInfo.Votetime=time.Now().Format("2006-01-02 15:04:05")
	mjInfo.Experiencecontent=""
	err=models.AddMeetingJoinInfo(mjInfo) //插入失败暂不做处理
	if err!=nil{
		models.AddLog("JoinMeeting Controller AddMeetingJoinInfo Error:"+err.Error(),0)
	}
	result.ErrMsg=""
	result.ErrCode=0
	result.Result=mInfo
	this.Data["json"]=result
	this.ServeJSON()
}

//刷新一场会议状态,只返回meetingcontent
func (this *RefreshMeeting)Get(){
	var result models.ResultInfo
	_,err:=checkLogin(this.GetSession("userinfo"))
	if  err!=nil{
		models.AddLog(result.ErrMsg,0)
		result.ErrCode=-5
		result.ErrMsg=err.Error()
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	intMeetingID,err:=this.GetInt("id")
	if err!=nil{
		models.AddLog("RefreshMeeting Controller Error:"+err.Error(),0)
		result.ErrCode=-7
		result.ErrMsg="参数传递错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	mInfo,err:=models.GetMeetingInfo(intMeetingID)
	if err!=nil{
		models.AddLog("RefreshMeeting Controller GetMeetingInfo Error:"+err.Error(),0)
		result.ErrCode=-8
		result.ErrMsg="获取会议信息出现错误"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	//result.Result=mInfo.Meetingcontent
	result.Result=mInfo
	this.Data["json"]=result
	this.ServeJSON()
}
//对一场会议进行投票
func(this *VoteMeeting)Post(){
	var result models.ResultInfo
	//_,err:=checkLogin(this.GetSession("userinfo"))
	//if  err!=nil{
	//	models.AddLog(result.ErrMsg,0)
	//	result.ErrCode=-5
	//	result.ErrMsg=err.Error()
	//	this.Data["json"]=result
	//	this.ServeJSON()
	//	return
	//}
	intMeetingID,err:=this.GetInt("id")
	intUserID,err1:=this.GetInt("userid")
	intUserExperience,err2:=this.GetInt("userexperience")
	strExperienCecontent:=this.GetString("experiencecontent")

	if err!=nil || err1!=nil || err2!=nil || strExperienCecontent==""{
		models.AddLog("RefreshMeeting Controller Error:"+err.Error(),0)
		result.ErrCode=-7
		result.ErrMsg="接收参数失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	//var mjInfo models.MeetingJoinInfo
	mjInfo,err:=models.GetMeetingJoinInfoByWhere(" meetingid="+strconv.Itoa(intMeetingID)+" and userid="+strconv.Itoa( intUserID))
	if err!=nil{
		result.ErrCode=-8
		result.ErrMsg="更新反馈信息失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	mjInfo.Userexperience=intUserExperience
	mjInfo.Experiencecontent=strExperienCecontent
	mjInfo.Votetime=time.Now().Format("2006-01-02 15:04:05")
	err=models.UpdateMeetingJoinInfo(mjInfo)
	if err!=nil{
		result.ErrCode=-9
		result.ErrMsg="更新反馈信息失败"
		this.Data["json"]=result
		this.ServeJSON()
		return
	}
	result.ErrCode=0
	result.ErrMsg=""
	this.Data["json"]=result
	this.ServeJSON()
}


/*检测用户登录状态
  传入参数
  ss 用户获取的session
  传出参数
  UserInfo 用户信息
  error 错误信息
 */
func checkLogin(ss interface{})(models.UserInfo,error){

	var userss models.UserInfo
	if ss==nil{
		return userss,errors.New("用户未登录")
	}
	userss=ss.(models.UserInfo)
	if userss.Id<=0{
		return userss,errors.New("登录超时,请重新登录")
	}
	return userss,nil
}
