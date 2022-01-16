package controllers

import (
	"MonitoringPlatform/models"
	"MonitoringPlatform/service"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type RequestRegisterController struct {
	beego.Controller
}

// Post /request
// @Description 调用post请求的时候，某个MEC发来的请求数与请求完成数自己统计后发给前端页面显示
// @Author Mateo
// @Date 17:05 2022/1/16
// 使用的json数据格式
// {
// 	 "reqNum": 10,
//   "reqComNum": 20,
// }
func (this *RequestRegisterController) Post() {
	//解析发送过来的Json数据为models.RequestInfo对象
	reqInfo := models.RequestInfo{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &reqInfo); err == nil {
		//调用设备登记功能模块，进行设备在线、异常等信息记录
		reqNum, reqComNum := service.AddServiceRequestNum(reqInfo)
		//通过websocket发送消息给前端，更新数据
		msg := models.ReqWebsocketMsg{
			Message:   "服务请求统计模块发送给前端的消息",
			MsgType:   "ServiceReqStatisticsModule",
			ReqNum:    reqNum,
			ReqComNum: reqComNum,
		}
		SendReqWebsocketMsg(msg)
		//返回正确状态
		this.returnStatus200()
	} else {
		//返回错误状态
		this.returnStatusError("传入的Json数据格式错误")
	}
}

func (this *RequestRegisterController) returnStatus200() {
	this.Data["json"] = map[string]interface{}{"status": "200"}
	err := this.ServeJSON()
	if err != nil {
		return
	}
}

func (this *RequestRegisterController) returnStatusError(errStr string) {
	this.Data["json"] = map[string]interface{}{"errMsg": errStr}
	err := this.ServeJSON()
	if err != nil {
		return
	}
}
