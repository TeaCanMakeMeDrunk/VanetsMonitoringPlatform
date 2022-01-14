package service

import (
	"MonitoringPlatform/models"
)

// DeviceRegister
// @Description 设备登记功能模块，包含设备在线登记，离线登记，异常登记等
// @Param devInfo是用户传来的设备信息
// @return (A,B,C,D,E) A是错误信息，如果为空则没错,B是改变了数值的某个设备在线数,C是改变了数值的某个设备异常数,D是更新后的所有设备在线数,E是更新后的所有设备异常数
// @Author Mateo
// @Date 11:26 2022/1/13
func DeviceRegister(devInfo models.DeviceInfo) (string, int, int, int, int) {
	switch devInfo.RegisterType {
	case "online":
		return addOnlineDevNum(devInfo)
	case "exc":
		return addExcDevNum(devInfo)
	}
	return "", 0, 0, 0, 0
}

// addOnlineDevNum
// @Description 设备在线情况统计功能,设备上线在线数+1,异常数保持不变
// @Param devInfo是用户传来的设备信息
// @return (A,B,C,D,E) A是错误信息，如果为空则没错,B是改变了数值的某个设备在线数,C是改变了数值的某个设备异常数,D是更新后的所有设备在线数,E是更新后的所有设备异常数
// @Author Mateo
// @Date 16:31 2022/1/12
func addOnlineDevNum(devInfo models.DeviceInfo) (string, int, int, int, int) {
	devOnlineNum := 0
	devExcNum := 0
	switch devInfo.DeviceType {
	case "OBU":
		models.StaVariables.ObuOnlineNum += devInfo.Num
		devOnlineNum = models.StaVariables.ObuOnlineNum
		devExcNum = models.StaVariables.ObuExcNum
	case "RSU":
		models.StaVariables.RsuOnlineNum += devInfo.Num
		devOnlineNum = models.StaVariables.RsuOnlineNum
		devExcNum = models.StaVariables.RsuExcNum
	case "MEC":
		models.StaVariables.MecOnlineNum += devInfo.Num
		devOnlineNum = models.StaVariables.MecOnlineNum
		devExcNum = models.StaVariables.MecExcNum
	case "camera":
		models.StaVariables.CameraOnlineNum += devInfo.Num
		devOnlineNum = models.StaVariables.CameraOnlineNum
		devExcNum = models.StaVariables.CameraExcNum
	case "core_pad":
		models.StaVariables.CorePadOnlineNum += devInfo.Num
		devOnlineNum = models.StaVariables.CorePadOnlineNum
		devExcNum = models.StaVariables.CorePadExcNum
	case "edge_pad":
		models.StaVariables.EdgePadOnlineNum += devInfo.Num
		devOnlineNum = models.StaVariables.EdgePadOnlineNum
		devExcNum = models.StaVariables.EdgePadExcNum
	default:
		return "设备类型或Json字段错误", devExcNum, devOnlineNum, -1, -1
	}
	models.StaVariables.TotalDevOnlineNum += devInfo.Num
	return "", devExcNum, devOnlineNum, models.StaVariables.TotalDevOnlineNum, models.StaVariables.TotalDevExcNum
}

// addExcDevNum
// @Description 设备异常情况统计功能,设备上线在线数-1,异常数+1, TODO解决异常应该重新写个API
// @Param devInfo是用户传来的设备信息
// @return (A,B,C,D,E) A是错误信息，如果为空则没错,B是改变了数值的某个设备在线数,C是改变了数值的某个设备异常数,D是更新后的所有设备在线数,E是更新后的所有设备异常数
// @Author Mateo
// @Date 16:31 2022/1/12
func addExcDevNum(devInfo models.DeviceInfo) (string, int, int, int, int) {
	devExcNum := 0
	devOnlineNum := 0
	switch devInfo.DeviceType {
	case "OBU":
		models.StaVariables.ObuOnlineNum -= devInfo.Num
		devOnlineNum = models.StaVariables.ObuOnlineNum
		models.StaVariables.ObuExcNum += devInfo.Num
		devExcNum = models.StaVariables.ObuExcNum
	case "RSU":
		models.StaVariables.RsuOnlineNum -= devInfo.Num
		devOnlineNum = models.StaVariables.RsuOnlineNum
		models.StaVariables.RsuExcNum += devInfo.Num
		devExcNum = models.StaVariables.RsuExcNum
	case "MEC":
		models.StaVariables.MecOnlineNum -= devInfo.Num
		devOnlineNum = models.StaVariables.MecOnlineNum
		models.StaVariables.MecExcNum += devInfo.Num
		devExcNum = models.StaVariables.MecExcNum
	case "camera":
		models.StaVariables.CameraOnlineNum -= devInfo.Num
		devOnlineNum = models.StaVariables.CameraOnlineNum
		models.StaVariables.CameraExcNum += devInfo.Num
		devExcNum = models.StaVariables.CameraExcNum
	case "core_pad":
		models.StaVariables.CorePadOnlineNum -= devInfo.Num
		devOnlineNum = models.StaVariables.CorePadOnlineNum
		models.StaVariables.CorePadExcNum += devInfo.Num
		devExcNum = models.StaVariables.CorePadExcNum
	case "edge_pad":
		models.StaVariables.EdgePadOnlineNum -= devInfo.Num
		devOnlineNum = models.StaVariables.EdgePadOnlineNum
		models.StaVariables.EdgePadExcNum += devInfo.Num
		devExcNum = models.StaVariables.CorePadExcNum
	default:
		return "设备类型或Json字段错误", devExcNum, devOnlineNum, -1, -1
	}
	models.StaVariables.TotalDevOnlineNum -= devInfo.Num
	models.StaVariables.TotalDevExcNum += devInfo.Num
	return "", devExcNum, devOnlineNum, models.StaVariables.TotalDevOnlineNum, models.StaVariables.TotalDevExcNum
}
