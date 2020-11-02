package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gin-web-skeleton/model/services"
	"gin-web-skeleton/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginField struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Csrf     string `form:"csrf"`
}

type RespData struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

func Login(c *gin.Context) {
	var loginField LoginField
	if err := c.Bind(&loginField); err != nil {
		fmt.Println(err)
		util.SendResponse(c, err, nil)
		return
	}
	fmt.Println(loginField)
	admin, err := services.GetAdminByUsername(loginField.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		c.Abort()
		return
	}
	if 1 != admin.Status {
		c.JSON(http.StatusOK, gin.H{"code": -2, "msg": "账号状态异常"})
		c.Abort()
		return
	}
	h := md5.New()
	h.Write([]byte(loginField.Password))
	md5Password := hex.EncodeToString(h.Sum(nil))

	if admin.Password != md5Password {
		c.JSON(http.StatusOK, gin.H{"code": -3, "msg": "密码错误"})
		c.Abort()
		return
	}

	token, err := util.GenerateToken(admin.Username, admin.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -4, "msg": "token创建失败，请重新登录"})
		c.Abort()
		return
	}

	respData := RespData{Username: admin.Username, Password: admin.Password, AccessToken: token}

	fmt.Println(respData)
	util.SendResponse(c, nil, respData)
	return
}
