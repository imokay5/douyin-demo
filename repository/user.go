package repository

import (
	"log"
	"sync"
	"time"

	"gorm.io/gorm"
)

var UsersLoginInfo map[string]User

type User struct {
	Id         int64     `gorm:"column:user_id"`
	Name       string    `gorm:"column:username"`
	Password   string    `gorm:"column:password"`
	CreateTime time.Time `gorm:"column:create_time"`
	ModifyTime time.Time `gorm:"column:modify_time"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

// 创建一个 UserDao指针，指向UserDao实例
func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		},
	)
	return userDao
}

// 当初始化数据库时，创建 map 的 toekn-user 映射，以方便后续的操作。
func (*UserDao) TokenMap() {
	UsersLoginInfo = make(map[string]User)
	result := make([]*User, 0)
	db.Find(&result)
	for _, val := range result {
		UsersLoginInfo[val.Name+val.Password] = *val // UserLoginInfo[token] = *user
	}

}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Where("user_id=?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Println("query user by id error", err)
		return nil, err
	}
	return &user, nil
}
