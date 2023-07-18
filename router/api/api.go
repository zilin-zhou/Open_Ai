package api

import (
	"open_ai_chat/controller/apicontroller"

	"github.com/gin-gonic/gin"
)

func InitApi(r *gin.Engine) {
	//分组
	api := r.Group("/api")
	{
		api.POST("/register", apicontroller.Register)
		api.POST("/delete", apicontroller.Delete)
		api.POST("/update", apicontroller.Update)
	}

}
