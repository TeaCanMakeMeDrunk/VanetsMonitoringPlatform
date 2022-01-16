package routers

import (
	"MonitoringPlatform/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//websocket的路由配置
	beego.Router("/ws", &controllers.MyWebSocketController{})
	//设备统计模块的post请求路由配置
	beego.Router("/device", &controllers.DeviceRegisterController{})
	//资源使用情况统计模块的post请求路由配置
	beego.Router("/resource", &controllers.ResourceRegisterController{})
	//服务请求统计模块的post请求路由配置
	beego.Router("/request", &controllers.RequestRegisterController{})
}
