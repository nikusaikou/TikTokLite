package ProtoMessage

import (
	"TikTokLite/log"
	"TikTokLite/proto/pkg"
	"encoding/json"
	"testing"
)

func TestProto(t *testing.T) {
	user := &message.User{
		Id:            123,
		Name:          "someName",
		FollowCount:   12,
		FollowerCount: 123,
		IsFollow:      false,
	}
	data, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Marshal error\n")
	}
	newUser := &message.User{}
	err = json.Unmarshal(data, newUser)
	log.Infof("%+v", newUser)
	if err != nil {
		t.Errorf("Unmarshal error\n")
	}
	if user.GetId() != newUser.GetId() {
		t.Errorf("user:%+v,newUser:%+v\n", user, newUser)
	}
}
