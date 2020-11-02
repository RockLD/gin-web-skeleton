package router

import (
	"gin-web-skeleton/app/admin/controller"
	"gin-web-skeleton/middleware"
	"github.com/gin-gonic/gin"
)

func LoadAdminRouter(g *gin.Engine, c *gin.Context) {

	r := g.Group("/admin")

	r.Use(middleware.AdminAuth())
	{
		r.GET("/login", controller.Login)
		r.POST("/login", controller.Login)
		r.GET("/welcome", controller.Welcome)
	}
}
