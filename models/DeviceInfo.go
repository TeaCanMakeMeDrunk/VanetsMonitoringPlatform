package models

type DeviceInfo struct {
	DeviceType   string `json:"deviceType"`   //设备类型, RSU,OBU..
	RegisterType string `json:"registerType"` //online 或者 exc
	Num          int    `json:"num"`          //在线或者异常数, 根据RegisterType变化
}
