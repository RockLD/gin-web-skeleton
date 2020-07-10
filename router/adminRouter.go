package router

import (
	"gin-web-skeleton/app/admin/controller"
	"gin-web-skeleton/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	csrf "github.com/utrack/gin-csrf"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadAdminRouter(g *gin.Engine, c *gin.Context) {
	g.LoadHTMLGlob("app/admin/templates/**/*")
	r := g.Group("/admin")

	store := cookie.NewStore([]byte("secret"))
	g.Use(sessions.Sessions("gin-web-skeleton", store))
	g.Use(csrf.Middleware(csrf.Options{
		Secret:        "123abc456def",
		IgnoreMethods: nil,
		ErrorFunc: func(c *gin.Context) {
			c.String(http.StatusBadRequest, "CSRF token mismatch")
			c.Abort()
		},
		TokenGetter: nil,
	}))

	csrf.GetToken(c)

	r.Use(middleware.AdminAuth())
	{
		r.GET("/login", controller.Login)
		r.POST("/login", controller.Login)
		r.GET("/index", controller.Index)
		r.POST("/index", controller.Index)
		r.GET("/welcome", controller.Welcome)
	}
}
