package repository

import (
	"log"
	"sync"
	"time"
)

// 跟随关系表
type Follow struct {
	FromUserId int64     `gorm:"column:user_id"`
	ToUserId   int64     `gorm:"column:to_user_id"`
	CreateTime time.Time `gorm:"column:create_time"`
}

// 表名
func (Follow) TableName() string {
	return "follow"
}

// 初始化 Dao 操作结构体实例
type FollowDao struct {
}

var followDao *FollowDao
var followOnce sync.Once

func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(
		func() {
			followDao = &FollowDao{}
		})
	return followDao
}

// 查询用户与用户之间的跟随关系
func QueryFollowInfo(fromId int64, toId int64) (int, error) {
	var count int64 = 0
	// 查询有几条符合条件的跟随关系
	err := db.Model(&Follow{}).Where(Follow{FromUserId: fromId, ToUserId: toId}).Count(&count).Error
	if err != nil {
		log.Println("query relation err:", err)
		return 0, err
	}
	if count == 0 {
		return 2, nil // 无
	} else {
		return 1, nil // 一次
	}
}
