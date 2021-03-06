package controller

import (
	"TikTokLite/log"
	"TikTokLite/response"
	"TikTokLite/service"
	"TikTokLite/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path/filepath"
	"strconv"
)

//视频发布
func PublishAction(ctx *gin.Context) {
	// publishResponse := &message.DouyinPublishActionResponse{}
	token := ctx.PostForm("token")
	log.Infof("token:%s", token)
	data, err := ctx.FormFile("data")
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	_, err = service.CheckCurrentUser(token)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	filename := filepath.Base(data.Filename)

	finalName := fmt.Sprintf("%s_%s", util.RandomString(), filename)
	videopath := viper.GetString("videofile")
	saveFile := filepath.Join(videopath, finalName)
	log.Debug(saveFile)
	if err := ctx.SaveUploadedFile(data, saveFile); err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	publish, err := service.PublishVideo(token, saveFile)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", publish)

}

//获取视频列表
func GetPublishList(ctx *gin.Context) {
	token := ctx.Query("token")
	id := ctx.Query("user_id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
	}
	list, err := service.PublishList(token, userId)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", list)
}
