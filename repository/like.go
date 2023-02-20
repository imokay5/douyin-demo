package repository

import (
	"log"
	"sync"
	"time"
)

type Like struct {
	UserId     int64     `gorm:"column:user_id"`
	VideoId    int64     `gorm:"column:video_id"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Like) TableName() string {
	return "likes"
}

type LikeDao struct {
}

var likeDao *LikeDao
var likeOnce sync.Once

func NewLikeDaoInstance() *LikeDao {
	likeOnce.Do(
		func() {
			likeDao = &LikeDao{}
		})
	return likeDao
}

// 查询 用户 是否喜欢某 视频
func QueryLike(videoId int64, userId int64) (bool, error) {
	var count int64 = 0
	err := db.Model(&Like{}).Where(Like{UserId: userId, VideoId: videoId}).Count(&count).Error
	if err != nil {
		log.Println("query relation err:", err)
		return false, err
	}
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

// 统计 某视频 被 多少人 喜欢
func CountLike(videoId int64) (int64, error) {
	var count int64
	err := db.Model(&Like{}).Where("video_id = ?", videoId).Count(&count).Error
	if err != nil {
		log.Println("count like err:", err)
		return -1, err
	}
	return count, nil
}
