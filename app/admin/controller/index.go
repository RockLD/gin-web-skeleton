package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {

	if "POST" == c.Request.Method {

	}

	c.HTML(http.StatusOK, "index/index.html", nil)
}

func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "index/welcome.html", nil)
}
