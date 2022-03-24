package router

import (
	"gin_chat/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("template/html/*")
	r.Static("/assets", "template/assets")

	chat := r.Group("")
	chat.GET("/ws", controller.Chat)
	chat.GET("/", controller.ChatRoomPage)

	return r
}
