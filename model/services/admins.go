package services

import (
	"gin-web-skeleton/model"
	"gin-web-skeleton/model/dao"
)

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
