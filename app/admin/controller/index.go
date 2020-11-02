package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "index/welcome.html", nil)
}
