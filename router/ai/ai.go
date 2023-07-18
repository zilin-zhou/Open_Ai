package ai

import (
	"open_ai_chat/controller/aicontroller"

	"github.com/gin-gonic/gin"
)

func InitAi(r *gin.Engine) {
	//分组
	api := r.Group("/ai")
	{
		api.POST("/chatGTP", aicontroller.ChatGTPRequest)

	}

}
