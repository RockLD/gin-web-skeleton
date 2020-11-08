package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gin-web-skeleton/model/services"
	"gin-web-skeleton/util"
	"github.com/gin-gonic/gin"
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
		sendResponse(c, -1, err.Error(), nil)
		return
	}
	if 1 != admin.Status {
		sendResponse(c, -2, "账号状态异常", nil)
		return
	}
	h := md5.New()
	h.Write([]byte(loginField.Password))
	md5Password := hex.EncodeToString(h.Sum(nil))

	if admin.Password != md5Password {
		sendResponse(c, -3, "密码错误", nil)
		return
	}

	token, err := util.GenerateToken(admin.Username, admin.Password)
	if err != nil {
		sendResponse(c, -4, "token创建失败，请重新登录", nil)
		return
	}

	respData := RespData{Username: admin.Username, Password: admin.Password, AccessToken: token}

	fmt.Println(respData)
	sendResponse(c, 0, "success", respData)
	return
}
