package v1

import (
	"msisensor-rna/middleware"
	"msisensor-rna/model"
	"msisensor-rna/utils/errmsg"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var uid uint

func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int

	formData, code = model.CheckLogin(formData.Username, formData.Password)

	if code == errmsg.SUCCESS {
		setToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": errmsg.GetErrMsg(code),
			"token":   token,
		})
	}

}

// LoginFront 前台登录
func LoginFront(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var code int
	var token string

	formData, code = model.CheckLoginFront(formData.Username, formData.Password)

	if code == errmsg.SUCCESS {
		setToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": errmsg.GetErrMsg(code),
			"token":   token,
		})
	}
}

func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	Claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "MSIsensor-RNA",
		},
	}

	token, err := j.CreateToken(Claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   token,
		})
	}
	uid = model.GetUserID(Claims.Username)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.Username,
		"id":      user.ID,
		"message": errmsg.GetErrMsg(200),
		"token":   token,
	})
}
