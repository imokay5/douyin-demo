package service

import (
	"log"
	"time"

	"github.com/imokay5/douyin-demo/repository"
)

// 视频数据
type VideoData struct {
	Id            int64    `json:"id,omitempty"`
	Author        UserData `json:"author"`
	PlayUrl       string   `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount int64    `json:"favorite_count,omitempty"`
	CommentCount  int64    `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
	Title         string   `json:"title,omitempty"`
}

var serverAddr string = "http://172.31.102.143:8080/"

// 上传视频
func PostVideo(fileName string, title string, userId int64) (*repository.Video, error) {
	playUrl := serverAddr + "public/" + fileName
	//////////////
	coverUrl := serverAddr + "public/bear-1283347_1280.jpg"
	newVideo := repository.Video{PlayUrl: playUrl, CoverUrl: coverUrl, Title: title, UserId: userId, CreateTime: time.Now()}
	video, err := repository.NewVideoDaoInstance().AddVideo(newVideo)
	if err != nil {
		log.Println("post video to db err:" + err.Error())
		return nil, err
	}
	return video, nil
}

func GetPublishList(userID int64, token string) ([]VideoData, error) {
	videoList, err := repository.NewVideoDaoInstance().SearchVideoById(userID)
	if err != nil {
		log.Println("search video by id err:", err)
		return nil, err
	}
	videoDataLsit, err := prepareVideoData(videoList, token)
	if err != nil {
		log.Println("prepare video data err:", err)
		return nil, err
	}
	return videoDataLsit, nil
}

func prepareVideoData(videoList []*repository.Video, token string) ([]VideoData, error) {
	var videoDataList []VideoData
	for _, k := range videoList {
		author, _ := QueryUserData(k.UserId, token)
		videoData := VideoData{Id: k.Id, Author: *author, PlayUrl: k.PlayUrl, CoverUrl: k.CoverUrl, Title: k.Title}
		videoDataList = append(videoDataList, videoData)
	}
	return videoDataList, nil
}
