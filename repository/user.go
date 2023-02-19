package repository

import (
	"log"
	"sync"
	"time"

	"gorm.io/gorm"
)

// 用户登录信息
var UsersLoginInfo map[string]User

// 用户信息
type User struct {
	Id         int64     `gorm:"column:user_id"`     // id
	Name       string    `gorm:"column:username"`    // 用户名
	Password   string    `gorm:"column:password"`    // 密码
	CreateTime time.Time `gorm:"column:create_time"` // 创建时间
	ModifyTime time.Time `gorm:"column:modify_time"` // 修改时间
}

// db 表名
func (User) TableName() string {
	return "user"
}

// Dao 数据库操作 结构体
type UserDao struct {
}

// Dao 结构体指针
var userDao *UserDao

// 使用一次的方法
var userOnce sync.Once

// 初始化 Dao 操作结构体实例
func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		},
	)
	return userDao
}

// 创建 token 和 User 的映射，当初始化数据库时使用
func (*UserDao) TokenMap() {
	UsersLoginInfo = make(map[string]User)
	result := make([]*User, 0)
	db.Find(&result)
	for _, val := range result {
		UsersLoginInfo[val.Name+val.Password] = *val // UserLoginInfo[token] = *user
	}
}

// 通过 id 查询 User
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
