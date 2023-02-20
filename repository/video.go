package repository

import (
	"log"
	"sync"
	"time"

	"gorm.io/gorm"
)

// 视频 表结构体
type Video struct {
	Id         int64     `gorm:"column:video_id"`
	PlayUrl    string    `gorm:"column:play_url"`
	CoverUrl   string    `gorm:"column:cover_url"`
	Title      string    `gorm:"column:title"`
	CreateTime time.Time `gorm:"column:create_time"`
	UserId     int64     `gorm:"column:user_id"`
}

// 视频 表名
func (Video) TableName() string {
	return "video"
}

// Dao 视频数据库操作集
type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

// 初始化 Dao 操作结构体实例
func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		},
	)
	return videoDao
}

// 通过 用户ID 查询其发布的视频列表
func (*VideoDao) SearchVideoById(userID int64) ([]*Video, error) {
	var videoList []*Video
	err := db.Model(&Video{}).Where("user_id=?", userID).Find(&videoList).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Println("query video list by user id error", err)
		return nil, err
	}
	return videoList, nil
}
