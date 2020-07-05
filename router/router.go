package router

import (
	"fmt"
	"gin-web-skeleton/app/index/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine, m ...gin.HandlerFunc) *gin.Engine {
	store := cookie.NewStore([]byte("secret"))
	g.Use(sessions.Sessions("mysession", store))
	g.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(401, "CSRF token mismatch")
			c.Abort()
		},
	}))

	g.Static("/static", "./public/static")
	// 格式化日志输出格式
	g.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))

	g.Use(gin.Recovery())
	g.Use(m...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The router is not exist!")
	})

	// 服务器健康检查
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	g.GET("/", api.Hello)

	// 管理后台路由
	LoadAdminRouter(g)
	return g
}
