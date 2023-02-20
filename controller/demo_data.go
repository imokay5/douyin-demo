package controller

import "github.com/imokay5/douyin-demo/service"

type UserData service.UserData

var DemoUser = UserData{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
	LikedCount:    0,
	LikeCount:     0,
}

var DemoVideos = []VideoData{
	{
		Id:            1,
		Author:        service.UserData(DemoUser),
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}
