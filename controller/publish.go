package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

// Publish 发布视频
// @Summary 发布视频接口
// @Description
// @Tags 基本接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /publish/action/ [post]
func Publish(c *gin.Context) {
	// 部分手机、模拟器可能会有问题 我的就发布不成功...
	token, ok := c.GetPostForm("token")
	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "Public Request Token Fail",
		})
		return
	}

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

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "Please login again"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		println("data error!\n")
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	user := usersLoginInfo[token]
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		println("save error!\n")
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	} else {
		success, errstring := service.Publish(saveFile, finalName, user)
		if success {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: 0,
				StatusMsg:  finalName + " uploaded successfully",
			})
		} else {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: 1,
				StatusMsg:  errstring,
			})
		}
	}
}

// PublishList 获取视频列表
// @Summary 获取视频列表接口
// @Description
// @Tags 基本接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} VideoListResponse
// @Router /publish/list/ [get]
func PublishList(c *gin.Context) {
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

	userVideoList := service.PublishList(usersLoginInfo[token].Id)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: userVideoList,
	})
}
