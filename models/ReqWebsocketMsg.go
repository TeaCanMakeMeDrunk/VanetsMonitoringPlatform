package models

type ReqWebsocketMsg struct {
	Message   string `json:"message"`   //发送给信息标志，仅作标志
	MsgType   string `json:"msgType"`   //信息的类型，来自哪个模块
	ReqNum    int    `json:"reqNum"`    //总的请求数
	ReqComNum int    `json:"reqComNum"` //已完成的请求数
}
