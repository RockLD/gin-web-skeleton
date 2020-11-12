package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"gin-web-skeleton/model"
	"gin-web-skeleton/model/dao"
	"github.com/spf13/viper"
)

type Admins struct {
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	RoleId   string `json:"role_id"`
	UserName string `json:"username"`
	Status   string `json:"status"`
	RealName string `json:"realname"`
}

func GetAdminByUsername(username string) (dao.Admins, error) {
	//return adminDao.GetAdminByUsername(username)
	var admin dao.Admins
	if err := model.DB.Self.Table(admin.TableName()).Where("username=?", username).First(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}

func GetAdminsByWhere(where map[string]interface{}, page, limit int) ([]dao.AdminInfo, error) {

	var list []dao.AdminInfo

	if err := model.DB.Self.Table(dao.Admins{}.TableName()).Where(where).Select("gws_admins.*,gws_roles.role_name").Joins("left join gws_roles on gws_roles.id=gws_admins.role_id").Offset(page - 1).Limit(limit).Order("gws_admins.id desc").Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (admins *Admins) AddAdmin() *Admins {
	status := 2
	if "on" == admins.Status {
		status = 1
	}

	password := viper.GetString("default_password")

	h := md5.New()
	h.Write([]byte(password))
	password = hex.EncodeToString(h.Sum(nil))

	adminsModel := dao.Admins{
		Username: admins.UserName,
		Status:   status,
		Password: password,
		Mobile:   admins.Mobile,
		Email:    admins.Email,
		RoleId:   admins.RoleId,
		RealName: admins.RealName,
	}

	err := model.DB.Self.Table(adminsModel.TableName()).Create(&adminsModel)
	fmt.Println(err)

	return admins
}

func (admins *Admins) EditAdmin(id int) (dao.Admins, error) {
	var count int
	var adminsModel dao.Admins
	model.DB.Self.Table(adminsModel.TableName()).Where("id != ?", id).Where("username=?", admins.UserName).First(&adminsModel).Count(&count)
	if count > 0 {
		fmt.Println(count)
		fmt.Println(adminsModel)
		return adminsModel, errors.New("该用户名已存在111")
	}

	status := 2
	if "on" == admins.Status {
		status = 1
	}
	model.DB.Self.Where("id = ?", id).First(&adminsModel)
	adminsModel.Username = admins.UserName
	adminsModel.Status = status
	adminsModel.Mobile = admins.Mobile
	adminsModel.RoleId = admins.RoleId
	adminsModel.RealName = admins.RealName

	err := model.DB.Self.Table(adminsModel.TableName()).Save(&adminsModel)

	if err != nil {
		return adminsModel, err.Error
	}

	return adminsModel, nil
}
