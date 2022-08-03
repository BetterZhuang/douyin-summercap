package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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
	//token := c.Query("token")
	//
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}

	token := c.Query("token")
	if _, exist := usersLoginInfo[token]; !exist {
		ResponseError(c, CodeInvalidToken)
		return
	}
	//解析token获取user.id
	user, err := service.ParseToken(token)
	if err != nil {
		zap.L().Error("ParseToken failed", zap.Error(err))
		ResponseError(c, CodeParseTokenFail)
		return
	}
	//获取relation对象即操作类型
	//user_id := c.Query("user_id")  // 应该用token中的比较安全？
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	video_id_num, err1 := strconv.Atoi(video_id)
	if err1 != nil {
		zap.L().Error("convert video_id from string to int64 failed", zap.String("touserid", video_id), zap.Error(err))
		ResponseErrorForLike(c, CodeInvalidParam)
		return
	}
	action_type_num, err2 := strconv.Atoi(action_type)
	if err2 != nil {
		zap.L().Error("convert touserid from string to int64 failed", zap.String("touserid", action_type), zap.Error(err))
		ResponseErrorForLike(c, CodeInvalidParam)
		return
	}

	liked, err := service.VideoLike(int64(user.Id), int64(video_id_num), int32(action_type_num))
	if err != nil { //点赞失败
		zap.L().Error("like failed", zap.String("userid", string(user.Id)), zap.Error(err))
		ResponseError(c, CodeOperateFail)
		return
	}
	if !liked { // err!=nil 而 liked=false，说明出错了
		//操作失败
		ResponseError(c, CodeOperateFail)
		return
	}
	//成功返回
	ResponseSuccessForLike(c, CodeSuccess)
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
	//c.JSON(http.StatusOK, VideoListResponse{
	//	Response: model.Response{
	//		StatusCode: 0,
	//	},
	//	VideoList: DemoVideos,
	//})

	token := c.Query("token")
	if _, exist := usersLoginInfo[token]; !exist {
		ResponseError(c, CodeInvalidToken)
		return
	}
	//解析token获取user.id
	user, err := service.ParseToken(token)
	if err != nil {
		zap.L().Error("ParseToken failed", zap.Error(err))
		ResponseError(c, CodeParseTokenFail)
		return
	}
	//获取relation对象即操作类型
	//user_id := c.Query("user_id")  // 应该用token中的比较安全？
	//根据userid获取点赞视频列表
	videoList, err := service.GetLikeList(int64(user.Id))
	ResponseSuccessForLikeList(c, videoList)
}
