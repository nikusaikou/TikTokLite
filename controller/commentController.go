package controller

import (
	"TikTokLite/log"
	"TikTokLite/response"
	"TikTokLite/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//发布评论
func CommentAction(ctx *gin.Context) {
	var err error
	token := ctx.Query("token")
	video_id := ctx.Query("video_id")
	comment_text := ctx.Query("comment_text")
	actionType := ctx.Query("action_type")
	comment_id := ctx.Query("comment_id")
	commentId := int64(0)
	if actionType == "2" {
		commentId, err = strconv.ParseInt(comment_id, 10, 64)
		if err != nil {
			log.Errorf("commentId error : %s", err)
			response.Fail(ctx, err.Error(), nil)
			return
		}
	}
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		log.Errorf("videoId error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	commentResponse, err := service.CommentAction(commentId, videoId, token, comment_text, actionType)
	if err != nil {
		log.Errorf("comment error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", commentResponse)
}

//获取评论列表
func GetCommentList(ctx *gin.Context) {
	var err error
	video_id := ctx.Query("video_id")
	token := ctx.Query("token")
	videoId, _ := strconv.ParseInt(video_id, 10, 64)

	listResponse, err := service.CommentList(token, videoId)
	if err != nil {
		log.Infof("list error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", listResponse)
}
