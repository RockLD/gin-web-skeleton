package middleware

import (
	"fmt"
	"gin-web-skeleton/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		requestUri := c.Request.RequestURI
		if "/admin/login" != requestUri {
			accessToken := c.Request.Header.Get("access_token")
			if accessToken == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Unauthorized"})
				c.Abort()
				return
			}

			//parts := strings.SplitN(accessToken," ", 2)
			//if !(len(parts) == 2 && parts[0] == "Bearer") {
			//	c.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"请求头中auth格式错误"})
			//	c.Abort()
			//	return
			//}

			mc, err := util.ParseToken(accessToken)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "无效的token", "data": accessToken})
				c.Abort()
				return
			}
			c.Set("username", mc.Username)
		}

		c.Next()
	}
}
