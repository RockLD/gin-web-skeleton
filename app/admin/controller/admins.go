package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Admins struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Status    int    `json:"status"`
	RealName  string `json:"real_name"`
	CreatedAt string `json:"created_at"`
	RoleName  string `json:"role_name"`
}

func AdminsList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
	})
}
