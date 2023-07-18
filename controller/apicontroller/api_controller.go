package apicontroller

import (
	"net/http"
	cf "open_ai_chat/config"
	"open_ai_chat/db"
	"open_ai_chat/db/tables"

	"github.com/gin-gonic/gin"
)

var (
	RegisterTableName = cf.Tables.GetString("newRegisterTable.name")
)

// 用户注册
func Register(ctx *gin.Context) {
	regBody := tables.Register{}
	err := ctx.Bind(&regBody)
	// TODO验证参数
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors":  err.Error(),
			"message": "bind failed!",
		})
	}
	//查询表中是否有相同的邮箱
	items, err := db.DB.Query("email", regBody.Email, RegisterTableName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors":  err.Error(),
			"message": "query failed!",
		})
	}
	//fmt.Println(items)
	if len(items) != 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors":  nil,
			"message": "this email has been registered!",
		})
	} else {
		//添加入表
		err = db.DB.AddRegisterInfo(regBody, RegisterTableName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errors":  err.Error(),
				"message": "add failed!",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"errors":  nil,
			"message": "register successful",
		})
	}
}

// 用户删除
func Delete(ctx *gin.Context) {
	regBody := tables.Register{}
	err := ctx.Bind(&regBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors":  err.Error(),
			"message": "bind failed!",
		})
	}
	//查询表中是否有相同的邮箱
	items, err := db.DB.Query("email", regBody.Email, RegisterTableName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors":  err.Error(),
			"message": "query failed!",
		})
	}
	if len(items) == 1 {
		//删除信息
		err = db.DB.DeleteUserInfo(regBody, RegisterTableName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errors":  err.Error(),
				"message": "delete failed!",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors":  nil,
			"message": "delete successful!",
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors":  nil,
			"message": "数据不唯一,表错误!",
		})
	}
}

// 用户更新
func Update(ctx *gin.Context) {
	regBody := tables.Register{}
	err := ctx.Bind(&regBody)
	// TODO验证参数
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors":  err.Error(),
			"message": "bind failed!",
		})
		return
	}
	err = db.DB.UpdateUserInfo(regBody, regBody.UserName, RegisterTableName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors":  err.Error(),
			"req":     regBody,
			"message": "update failed!",
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"errors":  nil,
		"message": "update successful!",
	})
}
