package service

import "MonitoringPlatform/models"

// AddServiceRequestNum
// @Description 请求数和请求完成数统计功能
// @Param reqInfo是用户传来的请求相关信息
// @return (A,B) A是总的请求数,B是请求完成数
// @Author Mateo
// @Date 16:48 2022/1/16
func AddServiceRequestNum(reqInfo models.RequestInfo) (int, int) {
	models.StaVariables.ReqNum += reqInfo.ReqNum
	models.StaVariables.ReqComNum += reqInfo.ReqComNum
	return models.StaVariables.ReqNum, models.StaVariables.ReqComNum
}
