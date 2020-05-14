package router

import (
	"gin-web-skeleton/app/admin/controller"
	"gin-web-skeleton/app/index/api"
	"gin-web-skeleton/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(g *gin.Engine, m ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())
	g.Use(m...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The router is not exist!")
	})

	g.Any("/", api.Hello)

	u := g.Group("/admin")
	u.Use(middleware.AdminAuth())
	{
		u.GET("/login", controller.Login)
	}
	return g
}
