package controllers

import (
	"MonitoringPlatform/models"
	"log"
)

//init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
func init() {
	go handleDevMessages()
}

//广播发送至页面
func handleDevMessages() {
	for {
		msg := <-devBroadcast
		//广播给所有连接的用户（我们只有index.html页面）
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// SendDevWebsocketMsg 发送消息给channel
func SendDevWebsocketMsg(msg models.DevWebsocketMsg) {
	devBroadcast <- msg
}
