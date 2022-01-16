package models

type RequestInfo struct {
	//RequestFrom string `json:"requestFrom"` //预留字段, 暂时无用, 表示哪个应用发来的请求相关信息
	ReqNum    int `json:"reqNum"`    //总的请求数
	ReqComNum int `json:"reqComNum"` //已完成的请求数
}
