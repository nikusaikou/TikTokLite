package routes

import (
	"TikTokLite/common"
	"TikTokLite/controller"

	"github.com/gin-gonic/gin"
)

// action 发布视频逻辑
// list 查看已发布的视频
func PublishRoutes(r *gin.RouterGroup) {
	publish := r.Group("publish")
	{
		publish.POST("/action/", common.AuthMiddleware(), controller.PublishAction)
		publish.GET("/list/", common.AuthWithOutMiddleware(), controller.GetPublishList)
	}
}
