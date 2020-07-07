package router

import (
	"gin-web-skeleton/app/admin/controller"
	"gin-web-skeleton/middleware"

	"github.com/gin-gonic/gin"
)

func LoadAdminRouter(g *gin.Engine) {
	g.LoadHTMLGlob("app/admin/templates/**/*")
	r := g.Group("/admin")
	r.Use(middleware.AdminAuth(g))
	{
		r.GET("/login", controller.Login)
		r.POST("/login", controller.Login)
		r.GET("/index", controller.Index)
		r.POST("/index", controller.Index)
		r.GET("/welcome", controller.Welcome)
	}
}
