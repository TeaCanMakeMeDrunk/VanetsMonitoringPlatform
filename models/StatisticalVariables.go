package models

// StatisticalVariables 统计变量类 先不作持久化。
type StatisticalVariables struct {
	//总的请求数
	ReqNum int
	//已完成的请求数
	ReqComNum int
	//设备在线数
	TotalDevOnlineNum int
	//设备异常数
	TotalDevExcNum int
	//OBU在线数
	ObuOnlineNum int
	//OBU异常数
	ObuExcNum int
	//RSU在线数
	RsuOnlineNum int
	//RSU异常数
	RsuExcNum int
	//MEC在线数
	MecOnlineNum int
	//MEC异常数
	MecExcNum int
	//camera在线数
	CameraOnlineNum int
	//camera异常数
	CameraExcNum int
	//核心网pad在线数
	CorePadOnlineNum int
	//核心网pad异常数
	CorePadExcNum int
	//边缘网pad在线数
	EdgePadOnlineNum int
	//边缘网pad异常数
	EdgePadExcNum int
}

// StaVariables 定义一个全局变量,暂未做持久化处理
var StaVariables = StatisticalVariables{
	ReqNum:            0,
	ReqComNum:         0,
	TotalDevOnlineNum: 0,
	TotalDevExcNum:    0,
	ObuOnlineNum:      0,
	ObuExcNum:         0,
	RsuOnlineNum:      0,
	RsuExcNum:         0,
	MecOnlineNum:      0,
	MecExcNum:         0,
	CameraOnlineNum:   0,
	CameraExcNum:      0,
	CorePadOnlineNum:  0,
	CorePadExcNum:     0,
	EdgePadOnlineNum:  0,
	EdgePadExcNum:     0,
}
