package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CommentAction 评论操作
// @Summary 评论接口
// @Description 登录用户对视频进行评论
// @Tags 扩展接口-I
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} CommentListResponse
// @Router /comment/action/ [post]
func CommentAction(c *gin.Context) {

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

	videoId_ := c.Query("video_id")
	videoId, err := strconv.Atoi(videoId_)
	if err != nil {
		zap.L().Error("convert video_id from string to int64 failed", zap.String("videoId", videoId_), zap.Error(err))
		ResponseForComment(c, CodeInvalidParam)
		return
	}
	actionType_ := c.Query("action_type")
	actionType, err := strconv.Atoi(actionType_)
	if err != nil {
		zap.L().Error("convert action_type from string to int32 failed", zap.String("action_type", actionType_), zap.Error(err))
		ResponseForComment(c, CodeInvalidParam)
		return
	}

	if actionType == 1 { // 发布评论
		commentText := c.Query("comment_text")
		if commentText == "" {
			zap.L().Error("convert action_type from string to int32 failed", zap.String("action_type", actionType_), zap.Error(err))
			ResponseForComment(c, CodeCommentNotExist)
			return
		}

		request := &model.CommentRequest{
			UserId:      int64(user.Id),
			VideoId:     int64(videoId),
			CommentText: commentText,
		}

		comment, err := service.AddComment(request)
		if err != nil {
			ResponseForComment(c, CodeOperateFail)
			return
		}
		ResponseSuccessForComment(c, comment)
	} else {
		commentId_ := c.Query("comment_id")

		commentId, err := strconv.Atoi(commentId_)
		if err != nil {
			zap.L().Error("convert comment_id from string to int64 failed", zap.String("commentId", commentId_), zap.Error(err))
			ResponseForComment(c, CodeInvalidParam)
			return
		}
		res := service.DeleteComment(int64(commentId), int64(videoId))
		if res == false {
			ResponseForComment(c, CodeOperateFail)
			return
		}
		ResponseSuccessForComment(c, model.Comment{})
	}

}

// CommentList 视频评论列表
// @Summary 视频评论列表接口
// @Description 查看视频的所有评论
// @Tags 扩展接口-I
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} CommentListResponse
// @Router /comment/list/ [get]
func CommentList(c *gin.Context) {

	token := c.Query("token")
	if _, exist := usersLoginInfo[token]; !exist {
		ResponseError(c, CodeInvalidToken)
		return
	}
	//解析token获取user.id
	_, err := service.ParseToken(token)
	if err != nil {
		zap.L().Error("ParseToken failed", zap.Error(err))
		ResponseError(c, CodeParseTokenFail)
		return
	}

	videoId_ := c.Query("video_id")
	videoId, err := strconv.Atoi(videoId_)
	if err != nil {
		zap.L().Error("convert video_id from string to int64 failed", zap.String("videoId", videoId_), zap.Error(err))
		ResponseForComment(c, CodeInvalidParam)
		return
	}

	commentList, err := service.GetCommentList(int64(videoId))
	ResponseSuccessForCommentList(c, commentList)
}
