package router

import (
	"open_ai_chat/middleware"
	"open_ai_chat/router/ai"
	"open_ai_chat/router/api"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.Use(middleware.LogInput)
	api.InitApi(r) //路由注册
	ai.InitAi(r)
}
