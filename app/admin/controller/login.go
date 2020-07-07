package controller

import (
	"gin-web-skeleton/util"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
)

type LoginField struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RespData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	if "POST" == c.Request.Method {
		var loginField LoginField
		if err := c.BindJSON(&loginField); err != nil {
			util.SendResponse(c, err, nil)
			c.Abort()
		}
		var respData RespData
		respData.Username = loginField.Username
		respData.Password = loginField.Password
		util.SendResponse(c, nil, respData)
	}
	csrfToken := csrf.GetToken(c)
	c.HTML(http.StatusOK, "login/login.html", gin.H{"csrfToken": csrfToken})
}
