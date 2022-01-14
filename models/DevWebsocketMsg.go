package models

type DevWebsocketMsg struct {
	Message       string `json:"message"`       //发送给信息标志，仅作标志
	MsgType       string `json:"msgType"`       //信息的类型，来自哪个模块
	ChangedDevice string `json:"changedDevice"` //改变的设备
	RegisterType  string `json:"registerType"`  //设备信息登记的类型
	DevExcNum     int    `json:"devExcNum"`     //某个设备异常数
	DevOnlineNum  int    `json:"devOnlineNum"`  //某个设备在线数
	TotalDevExcNum     int    `json:"totalDevExcNum"`     //所有设备异常数
	TotalDevOnlineNum  int    `json:"totalDevOnlineNum"`  //所有设备在线数
}
