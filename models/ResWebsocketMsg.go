package models

type ResWebsocketMsg struct {
	Message       string `json:"message"`       //发送给信息标志，仅作标志
	MsgType       string `json:"msgType"`       //信息的类型，来自哪个模块
	ResourceType string `json:"resourceType"` //资源类别, 带宽, 算力等
	Percentage   int    `json:"percentage"`   //资源使用百分比
}
