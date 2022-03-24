package controller

import (
	"gin_chat/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChatRoomPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chatroom.html", nil)
}

func Chat(ctx *gin.Context) {
	global.Melody.HandleRequest(ctx.Writer, ctx.Request)
}
