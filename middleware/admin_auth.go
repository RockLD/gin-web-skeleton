package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func AdminAuth(g *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := cookie.NewStore([]byte("secret"))
		g.Use(sessions.Sessions("mysession", store))
		g.Use(csrf.Middleware(csrf.Options{
			Secret: "secret123",
			ErrorFunc: func(c *gin.Context) {
				c.String(401, "CSRF token mismatch")
				c.Abort()
			},
		}))
		c.Next()
	}
}
