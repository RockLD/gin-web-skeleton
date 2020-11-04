package controller

import (
	"gin-web-skeleton/model/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Admins struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Status    int       `json:"status"`
	RealName  string    `json:"real_name"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
	RoleName  string    `json:"role_name"`
}

type Resp struct {
}

func AdminsList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	list, err := services.GetAdminsByWhere(nil, page, limit)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
		c.Abort()
		return
	}
	var admins []Admins
	if len(list) != 0 {
		for _, vo := range list {
			admins = append(admins, Admins{
				ID:       vo.ID,
				Username: vo.Username,
				Email:    vo.Email,
				Status:   vo.Status, RealName: vo.RealName, CreatedAt: vo.CreatedAt, Mobile: vo.Mobile})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "ok",
		"count": len(admins),
		"data":  admins,
	})
	c.Abort()
	return
}
