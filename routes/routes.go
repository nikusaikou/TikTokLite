package routes

import (
	"TikTokLite/common"
	"TikTokLite/controller"

	"github.com/gin-gonic/gin"
)

// feed 流单独不做鉴权，其他设置一个分组路由做鉴权
func SetRoute(r *gin.Engine) *gin.Engine {
	douyin := r.Group("/douyin")
	{
		UserRoutes(douyin)
		PublishRoutes(douyin)
		CommentRoutes(douyin)
		FavoriteRoutes(douyin)
		RelationRoutes(douyin)
		douyin.GET("/feed/", common.AuthWithOutMiddleware(), controller.Feed)
	}

	return r
}
