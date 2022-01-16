package controllers

import (
	"MonitoringPlatform/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

// ResourceRegisterController 资源使用情况登记
type ResourceRegisterController struct {
	beego.Controller
}

// Post /resource
// @Description 带宽、算力等资源使用情况登记调用此API进行登记
// @Author Mateo
// @Date 11:03 2022/1/13
// 使用的json数据格式
// {
// 	 "resourceType":"bandwidth", //bandwidth可替换为comPower,uplink,downlink
// 	 "percentage":10   			//资源使用百分比
// }
func (this *ResourceRegisterController) Post() {
	//解析发送过来的Json数据为models.DeviceInfo对象
	resourceInfo := models.ResourceInfo{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &resourceInfo); err == nil {
		//通过websocket发送消息给前端，更新数据
		msg := models.ResWebsocketMsg{
			Message:      "资源使用统计模块发送给前端的消息",
			MsgType:      "ResourceStatisticsModule",
			ResourceType: resourceInfo.ResourceType,
			Percentage:   resourceInfo.Percentage,
		}
		SendResWebsocketMsg(msg)
		//返回正确状态
		this.returnStatus200()
	} else {
		//返回错误状态
		this.returnStatusError("传入的Json数据格式错误")
	}
}

func (this *ResourceRegisterController) returnStatus200() {
	this.Data["json"] = map[string]interface{}{"status": "200"}
	err := this.ServeJSON()
	if err != nil {
		return
	}
}

func (this *ResourceRegisterController) returnStatusError(errStr string) {
	this.Data["json"] = map[string]interface{}{"errMsg": errStr}
	err := this.ServeJSON()
	if err != nil {
		return
	}
}
