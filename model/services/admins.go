package services

import "gin-web-skeleton/model/dao"

var adminDao dao.Admins

func GetAdminByUsername(username string) (dao.Admins, error) {
	return adminDao.GetAdminByUsername(username)
}

func GetAdminsByWhere(where map[string]interface{}, page, limit int) ([]dao.AdminInfo, error) {
	return adminDao.GetAdminsByWhere(where, page, limit)
}
