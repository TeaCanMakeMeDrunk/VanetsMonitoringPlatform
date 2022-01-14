package main

import (
	_ "MonitoringPlatform/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
