package service

import (
	"log"

	"github.com/imokay5/douyin-demo/repository"
)

// 用户数据信息
type UserData struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	LikedCount    int64  `json:"total_favorited,omitempty"`
	LikeCount     int64  `json:"favorite_count,omitempty"`
}

// 通过 用户id 和 token 查询 用户数据
func QueryUserData(userid int64, token string) (*UserData, error) {
	user, err := repository.NewUserDaoInstance().QueryUserById(userid)
	if err != nil {
		log.Println("query user by id error ", err)
		return nil, err
	}
	////////////////////////////////
	var userData = UserData{Id: user.Id, Name: user.Name}

	return &userData, nil
}
