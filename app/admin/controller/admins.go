package controller

import (
	"gin-web-skeleton/model/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Admins struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Status    int       `json:"status"`
	RealName  string    `json:"realname"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
	RoleName  string    `json:"role_name"`
}

/**
 * 获取管理员列表
 */
func AdminsList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	where := make(map[string]interface{})

	if c.Query("username") != "" {
		where["username"] = strings.TrimSpace(c.Query("username"))
	}

	if c.Query("mobile") != "" {
		where["mobile"] = strings.TrimSpace(c.Query("mobile"))
	}

	if c.Query("email") != "" {
		where["email"] = strings.TrimSpace(c.Query("email"))
	}

	if roleId, _ := strconv.Atoi(c.Query("role_id")); roleId != 0 {
		where["role_id"] = roleId
	}

	list, err := services.GetAdminsByWhere(where, page, limit)
	if err != nil {
		sendResponse(c, -1, err.Error(), nil)
		return
	}
	var admins []Admins
	if len(list) != 0 {
		for _, vo := range list {
			admins = append(admins, Admins{
				ID:       vo.ID,
				Username: vo.Username,
				Email:    vo.Email,
				Status:   vo.Status,
				RealName: vo.RealName, CreatedAt: vo.CreatedAt, Mobile: vo.Mobile, RoleName: vo.RoleName})
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

func AdminsAdd(c *gin.Context) {
	var admins services.Admins

	admins.UserName = c.PostForm("username")
	admins.Mobile = c.PostForm("mobile")
	admins.Email = c.PostForm("email")
	admins.RoleId = c.PostForm("role_id")
	admins.Status = c.PostForm("status")
	admins.RealName = c.PostForm("realname")
	res := admins.AddAdmin()
	c.JSON(http.StatusOK, gin.H{"data": res})
	c.Abort()
	return
}

func AdminsEdit(c *gin.Context) {
	var admins services.Admins

	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		sendResponse(c, -1, err.Error(), nil)
		return
	}
	admins.UserName = c.PostForm("username")
	admins.Mobile = c.PostForm("mobile")
	admins.Email = c.PostForm("email")
	admins.RoleId = c.PostForm("role_id")
	admins.Status = c.PostForm("status")
	admins.RealName = c.PostForm("realname")
	res, err := admins.EditAdmin(id)
	if err != nil {
		sendResponse(c, -2, err.Error(), nil)
		return
	}
	sendResponse(c, 0, "success", res)
	return
}

func AdminsDel(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		sendResponse(c, -1, err.Error(), nil)
		return
	}
	var admins services.Admins
	err = admins.DelAdmin(id)
	if err != nil {
		sendResponse(c, -2, err.Error(), nil)
	} else {
		sendResponse(c, 0, "success", nil)
	}
	return
}
