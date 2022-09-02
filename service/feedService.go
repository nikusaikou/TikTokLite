package service

import (
	message "TikTokLite/proto/pkg"
	"TikTokLite/repository"
	"TikTokLite/util"
)

// 参数： 时间戳， 当前用户 id
// 返回 feed 响应结构
// 按照当前时间从数据库读取视频列表，封装进 feed 响应结构，更新时间戳
func GetFeedList(currentTime, tokenUserId int64) (*message.DouyinFeedResponse, error) {
	videoList, err := repository.GetVideoListByFeed(currentTime)
	if err != nil {
		return nil, err
	}
	feed := &message.DouyinFeedResponse{
		VideoList: VideoList(videoList, tokenUserId),
	}

	nextTime := util.GetCurrentTime()
	if len(videoList) == 20 {
		nextTime = videoList[len(videoList)-1].PublishTime
	}
	feed.NextTime = nextTime
	return feed, nil
}

// 该函数用于打包视频响应结构体，去数据库取得视频信息，封装在视频 protobuf 的 video 结构内
func VideoList(videoList []repository.Video, userId int64) []*message.Video {
	var err error
	FollowList := make(map[int64]struct{})
	favList := make(map[int64]struct{})
	if userId != int64(0) {
		FollowList, err = tokenFollowList(userId)
		if err != nil {
			return nil
		}
		favList, err = tokenFavList(userId)
		if err != nil {
			return nil
		}
	}
	lists := make([]*message.Video, len(videoList))
	for i, video := range videoList {
		v := &message.Video{
			Id:            video.Id,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Author:        messageUserInfo(video.Author),
			Title:         video.Title,
		}
		if _, ok := FollowList[video.AuthorId]; ok {
			v.Author.IsFollow = true
		}
		if _, ok := favList[video.Id]; ok {
			v.IsFavorite = true
		}
		lists[i] = v
	}
	return lists
}
