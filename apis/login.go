package apis

import (
	myjwt "../middleware/jwt"
	. "../models"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func generateToken(c *gin.Context, user User) {
	j := myjwt.NewJWT()
	claims := myjwt.CustomClaims{
		user.Username,
		user.Nickname,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600),
		},
	}

	token, err := j.CreateToken(claims)
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
			"token": token,
		},
	})
	return
}

// HandleLogin check user and return token
func HandleLogin(c *gin.Context) {
	var l LoginReq

	err := c.ShouldBindJSON(&l)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": gin.H{},
		})
		return
	}

	isPass, user, err := LoginCheck(l)
	if !isPass {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "check failed" + err.Error(),
			"data": gin.H{},
		})
		return
	}

	generateToken(c, user)
}
