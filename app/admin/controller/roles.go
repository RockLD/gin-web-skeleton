package controller

import (
	"gin-web-skeleton/model/services"
	"github.com/gin-gonic/gin"
)

type Roles struct {
	ID       int64  `json:"id"`
	RoleName string `json:"role_name"`
	Rule     string `json:"rule"`
	Status   string `json:"status"`
}

func RolesAll(c *gin.Context) {
	var roles []Roles

	if list, err := services.GetAllRoles(1); err != nil {
		sendResponse(c, -1, err.Error(), nil)
	} else {
		for _, vo := range list {
			roles = append(roles, Roles{ID: vo.ID, RoleName: vo.RoleName})
		}
		sendResponse(c, 0, "ok", roles)
	}

	c.Abort()
	return
}
