package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction 点赞或取消点赞
// @Summary 点赞接口
// @Description 登录用户对视频的点赞或取消点赞操作
// @Tags 扩展接口-I
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /favorite/action/ [post]
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList 获取所有点赞视频
// @Summary 获取所有点赞视频接口
// @Description 所有登录用户的点赞视频
// @Tags 扩展接口-I
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} VideoListResponse
// @Router /favorite/list/ [get]
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
