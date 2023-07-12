package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"user/internal/service"
)

type User struct {
	UserId         uint   `gorm:"primarykey"`
	UserName       string `gorm:"unique"`
	NickName       string
	PasswordDigest string
}

const (
	PasswordCost = 12
)

// 判断用户是否存在
func (user *User) CheckUserExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name=?", req.UserName).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// 获取用户信息
func (user *User) ShowUserInfo(req *service.UserRequest) (err error) {
	if exist := user.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("UserName Not Exist")
}

func (*User) UserCreate(req *service.UserRequest) (err error) {
	var count int64
	DB.Where("user_name=?", req.UserName).Count(&count)
	if count != 0 {
		return errors.New("UserName Exist")
	}
	user := User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	_ = user.SetPassword(req.Password)
	err = DB.Create(&user).Error
	return err
}

// 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		panic(err)
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// 检验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func BuildUser(item User) *service.UserModel {
	userModel := service.UserModel{
		UserID:   uint32(item.UserId),
		UserName: item.UserName,
		NickName: item.NickName,
	}
	return &userModel
}
