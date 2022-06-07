package controller

import (
	"github.com/RaymondCode/simple-demo/middleware"
	"net/http"
	"time"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

var usersLoginInfo map[string]model.User = make(map[string]model.User)

type UserLoginResponse struct {
	model.Response
	UserId uint64 `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	model.Response
	User model.User `json:"user"`
}

// Register 用户注册接口
// @Summary 用户注册接口
// @Description
// @Tags 基本接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} UserLoginResponse
// @Router /user/register/ [post]
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//密码加盐存储
	md5Password := middleware.MD5_SALT(password)

	uid, ok := service.Register(username, md5Password)
	if !ok {
		println("Register fail")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "Register fail"},
		})
		return
	}

	token, err := service.GenerateToken(uid)
	if err != nil {
		println("Generate token fail")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "Generate token fail"},
		})
		return
	}

	newUser := model.User{
		Id:   uid,
		Name: username,
	}
	usersLoginInfo[token] = newUser
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: model.Response{StatusCode: 0, StatusMsg: "注册成功"},
		UserId:   uid,
		Token:    token,
	})

}

// Login 用户登录接口
// @Summary 用户登录接口
// @Description
// @Tags 基本接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} UserLoginResponse
// @Router /user/login/ [post]
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//当前登录密码加盐后与db中的密码进行比较
	md5Password := middleware.MD5_SALT(password)

	user, ok := service.Login(username, md5Password)
	if !ok {
		println("Login fail")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "Login fail! Check your username and password"},
		})
		return
	}
	// token := username + password
	token, err := service.GenerateToken(user.Id)
	if err != nil {
		println("Generate token fail", err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "Generate token fail"},
		})
		return
	}
	usersLoginInfo[token] = user
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: model.Response{StatusCode: 0, StatusMsg: "Login Successfully"},
		UserId:   user.Id,
		Token:    token,
	})
}

// UserInfo 获取当前登录用户信息
// @Summary 获取当前登录用户接口
// @Description
// @Tags 基本接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} UserResponse
// @Router /user/ [get]
func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "No Token!"})
		return
	} else {
		claims, err := service.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "Parse Token fail!"})
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "Token Timeout! Please login again"})
			return
		}
	}

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "Please Login again"},
		})
	}
}
