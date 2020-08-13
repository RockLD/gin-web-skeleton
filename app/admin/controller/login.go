package controller

import (
	"fmt"
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
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	if "POST" == c.Request.Method {
		var loginField LoginField
		if err := c.Bind(&loginField); err != nil {
			fmt.Println(err)
			util.SendResponse(c, err, nil)
			return
		}
		var respData RespData
		respData.Username = loginField.Username
		respData.Password = loginField.Password
		fmt.Println(respData)
		util.SendResponse(c, nil, respData)
		return
	}
	//token := csrf.GetToken(c)
	c.HTML(http.StatusOK, "login/login.html", nil)
}
