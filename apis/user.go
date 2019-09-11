package apis

import (
	. "../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleUserGetAll handle user getall
func HandleUserGetAll(c *gin.Context) {
	var u User
	users, err := u.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": gin.H{},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": users,
	})
}

// HandleUserGetOne handle user getone
func HandleUserGetOne(c *gin.Context) {
	uid := c.Param("uid")
	var u User
	user, err := u.GetOne(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": gin.H{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": user,
	})
}

// HandleUserCreate handle user create
func HandleUserCreate(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": gin.H{},
		})
		return
	}

	uid, err := user.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": gin.H{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"uid": uid,
		},
	})
}

// HandleUserUpdate handle user update
func HandleUserUpdate(c *gin.Context) {
	uid := c.Param("uid")
	var u User
	u.Username = c.DefaultPostForm("username", "")
	u.Password = c.DefaultPostForm("password", "")
	u.Nickname = c.DefaultPostForm("nickname", "")

	user, err := u.Update(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": gin.H{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": user,
	})
}

// HandleUserDelete handle user delete
func HandleUserDelete(c *gin.Context) {
	uid := c.Param("uid")

	var u User
	affected, err := u.Delete(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": gin.H{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"affected": affected,
		},
	})
}
