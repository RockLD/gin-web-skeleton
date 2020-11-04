package dao

import (
	"fmt"
	"gin-web-skeleton/model"
)

type Admins struct {
	model.BaseModel
	ID       int64  `gorm:"primary_key;column:id;type:bigint(20);not null"`
	Username string `gorm:"column:username;type:varchar(20);not null"`  // 用户名
	Password string `gorm:"column:password;type:varchar(128);not null"` // 密码
	Email    string `gorm:"column:email;type:varchar(255);not null"`    // 邮箱
	Status   int    `gorm:"column:status;type:int(10);not null"`        // 状态
	RealName string `gorm:"column:realname;type:varchar(255);not null"` // 真实姓名
	RoleId   string `gorm:"column:role_id;type:int(10);not null"`       // 角色ID
	Mobile   string `gorm:"column:mobile;type:varchar(32);not null"`    // 手机
	Roles    Roles
}

func (admin Admins) GetAdminByUsername(username string) (Admins, error) {
	d := model.DB.Self.Where("username=?", username).First(&admin)
	return admin, d.Error
}

func (admin Admins) GetAdminsByWhere(where map[string]string, page, limit int) ([]Admins, error) {
	admins := []Admins{}
	model.DB.Self.Offset(page - 1).Limit(limit).Order("id desc").Find(&admins)
	fmt.Println(admins)
	return admins, nil
}
