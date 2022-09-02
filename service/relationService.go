package service

import (
	"TikTokLite/log"
	message "TikTokLite/proto/pkg"
	"TikTokLite/repository"
	"errors"
)

// 关注行为
// 参数： 对方 id， 自己 id， 关注动作标记
func RelationAction(toUserId, tokenUserId int64, action string) error {
	if tokenUserId == toUserId {
		return errors.New("you can't follow yourself")
	}
	if action == "1" {
		log.Infof("follow action id:%v,toid:%v", tokenUserId, toUserId)
		err := repository.FollowAction(tokenUserId, toUserId)
		if err != nil {
			return err
		}
	} else {
		log.Infof("unfollow action id:%v,toid:%v", tokenUserId, toUserId)
		err := repository.UnFollowAction(tokenUserId, toUserId)
		if err != nil {
			return err
		}
	}
	return nil
}

// 关注列表
// 参数 对方id 自己id
// 读取双方的关注列表，如果存在共同关注则进行标记。
func RelationFollowList(userId int64, tokenUserId int64) (*message.DouyinRelationFollowListResponse, error) {
	followList, err := repository.GetFollowList(userId, "follow")
	if err != nil {
		return nil, err
	}
	log.Infof("user:%v, followList:%+v", userId, followList)
	list, err := tokenFollowList(tokenUserId)
	if err != nil {
		return nil, err
	}
	followListResponse := message.DouyinRelationFollowListResponse{
		UserList: make([]*message.User, len(followList)),
	}
	for i, u := range followList {
		follow := messageUserInfo(u)
		if _, ok := list[follow.Id]; ok {
			follow.IsFollow = true
		}
		followListResponse.UserList[i] = follow
	}

	return &followListResponse, nil
}

// 粉丝列表
func RelationFollowerList(userId int64, tokenUserId int64) (*message.DouyinRelationFollowerListResponse, error) {
	followList, err := repository.GetFollowList(userId, "follower")
	if err != nil {
		return nil, err
	}
	log.Infof("user:%v, followerList:%+v", userId, followList)
	list, err := tokenFollowList(tokenUserId)
	if err != nil {
		return nil, err
	}
	followListResponse := message.DouyinRelationFollowerListResponse{
		UserList: make([]*message.User, len(followList)),
	}
	for i, u := range followList {
		follow := messageUserInfo(u)
		if _, ok := list[follow.Id]; ok {
			follow.IsFollow = true
		}
		followListResponse.UserList[i] = follow
	}

	return &followListResponse, nil
}

func tokenFollowList(userId int64) (map[int64]struct{}, error) {
	m := make(map[int64]struct{})
	list, err := repository.GetFollowList(userId, "follow")
	if err != nil {
		return nil, err
	}
	for _, u := range list {
		m[u.Id] = struct{}{}
	}
	return m, nil
}
