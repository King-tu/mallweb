package models

import "github.com/jinzhu/gorm"

type User struct {
	Id        int
	Name      string `gorm:"size(20);unique"` //用户名
	PassWord  string `gorm:"size(20)"`        //登陆密码
	Email     string `gorm:"size(50);null"`   //邮箱
	Active    bool   `gorm:"default(false)"`  //是否激活
	Power     int    `gorm:"default(0)"`      //权限设置  0 表示未激活  1表示激活
	Address   []*Address
	OrderInfo []*OrderInfo
}

type Address struct {
	Id        int
	Receiver  string `gorm:"size(20)"`      //收件人
	Addr      string `gorm:"size(50)"`      //收件地址
	Zipcode   string `gorm:"size(20)"`      //邮编
	Phone     string `gorm:"size(20)"`      //联系方式
	Isdefault bool   `gorm:"defalt(false)"` //是否默认 0 为非默认  1为默认
	UserID    int
	OrderInfo []*OrderInfo
}

func init() {
	// 自动迁移(创建)表
	DB().AutoMigrate(new(User), new(Address))
}

func (user *User) Add() error {
	return DB().Create(user).Error
}

func (user *User) Get() error {
	err := db.Where("name = ?", user.Name).First(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (user *User) Exists() (bool, error) {
	// Select，指定你想从数据库中检索出的字段，默认会选择全部字段。
	err := DB().Select("id").Where("name = ? ", user.Name).First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.Id > 0 {
		return true, nil
	}
	return false, nil
}
