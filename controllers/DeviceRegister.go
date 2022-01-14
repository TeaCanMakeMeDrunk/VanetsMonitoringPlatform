package controllers

import (
	"MonitoringPlatform/models"
	"MonitoringPlatform/service"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

// DeviceRegisterController 设备在线登记
type DeviceRegisterController struct {
	beego.Controller
}

// Post /device
// @Description 设备在线、离线等状态调用此API进行登记
// @Author Mateo
// @Date 11:03 2022/1/13
// 设备在线登记使用的json数据格式
// {
// 	 "deviceType":"OBU", //OBU可替换为RSU,MEC,camera,core_pad,edge_pad
//   "registerType":"online",
// 	 "num":1
// }
// 设备异常登记使用的json数据格式
// {
// 	 "deviceType":"OBU",
//   "registerType":"exc",
// 	 "num":1
// }
func (this *DeviceRegisterController) Post() {
	//解析发送过来的Json数据为models.DeviceInfo对象
	devInfo := models.DeviceInfo{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &devInfo); err == nil {
		//调用设备登记功能模块，进行设备在线、异常等信息记录
		errStr, devExcNum, devOnlineNum, totalDevOnlineNum, totalDevExcNum := service.DeviceRegister(devInfo)
		if errStr != "" {
			//返回错误状态
			this.returnStatusError(errStr)
		} else {
			//通过websocket发送消息给前端，更新数据
			msg := models.DevWebsocketMsg{
				Message:        "设备情况统计模块发送给前端的消息",
				MsgType:        "DeviceStatisticsModule",
				RegisterType:   devInfo.RegisterType,
				ChangedDevice:  devInfo.DeviceType,
				DevExcNum:      devExcNum,
				DevOnlineNum:   devOnlineNum,
				TotalDevOnlineNum: totalDevOnlineNum,
				TotalDevExcNum: totalDevExcNum,
			}
			SendDevWebsocketMsg(msg)
			//返回正确状态
			this.returnStatus200()
		}
	} else {
		//返回错误状态
		this.returnStatusError("传入的Json数据格式错误")
	}
}

func (this *DeviceRegisterController) returnStatus200() {
	this.Data["json"] = map[string]interface{}{"status": "200"}
	err := this.ServeJSON()
	if err != nil {
		return
	}
}

func (this *DeviceRegisterController) returnStatusError(errStr string) {
	this.Data["json"] = map[string]interface{}{"errMsg": errStr}
	err := this.ServeJSON()
	if err != nil {
		return
	}
}
