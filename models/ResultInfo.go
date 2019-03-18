package models
//与前台交互的json结构
type ResultInfo struct {
	Result interface{}
	ErrCode int
	ErrMsg string
}