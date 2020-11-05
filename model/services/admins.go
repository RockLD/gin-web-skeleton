package services

import "gin-web-skeleton/model/dao"

var adminDao dao.Admins

func GetAdminByUsername(username string) (dao.Admins, error) {
	return adminDao.GetAdminByUsername(username)
}

func GetAdminsByWhere(where map[string]string, page, limit int) ([]dao.AdminList, error) {
	return adminDao.GetAdminsByWhere(where, page, limit)
}
