package controller

import (
	"gin-web-skeleton/model/services"
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ok",
		})
	} else {
		for _, vo := range list {
			roles = append(roles, Roles{ID: vo.ID, RoleName: vo.RoleName})
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": roles,
		})
	}

	c.Abort()
	return
}
