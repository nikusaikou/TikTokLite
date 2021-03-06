package controller

import (
	"TikTokLite/response"
	"TikTokLite/service"

	"github.com/gin-gonic/gin"
)

type FavActionParams struct {
	// 暂时没 user_id ，因为客户端出于安全考虑没给出
	Token      string `form:"token" binding:"required"`
	VideoId    int64  `form:"video_id" binding:"required"`
	ActionType int8   `form:"action_type" binding:"required,oneof=1 2"`
}

type FavListParams struct {
	Token  string `form:"token" binding:"required"`
	UserId int64  `form:"user_id" binding:"required"`
}

//点赞视频
func FavoriteAction(ctx *gin.Context) {
	var favInfo FavActionParams
	err := ctx.ShouldBindQuery(&favInfo)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	err = service.FavoriteAction(favInfo.Token, favInfo.VideoId, favInfo.ActionType)

	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", nil)
}

//获取点赞列表
func GetFavoriteList(ctx *gin.Context) {
	var listInfo FavListParams
	err := ctx.ShouldBindQuery(&listInfo)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	favList, err := service.FavoriteList(listInfo.Token, listInfo.UserId)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", favList)
}
