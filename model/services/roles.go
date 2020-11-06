package services

import (
	"gin-web-skeleton/model"
	"gin-web-skeleton/model/dao"
)

/**
 * 获取所有的角色列表
 */
func GetAllRoles(status int) ([]dao.Roles, error) {
	var roleList []dao.Roles
	if err := model.DB.Self.Table(dao.Roles{}.TableName()).Where("status=?", 1).Find(&roleList).Error; err != nil {
		return roleList, err
	}
	return roleList, nil
}
