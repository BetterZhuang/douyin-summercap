package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// RelationAction 关系操作
// @Summary 关系操作接口
// @Description 登录用户对其他用户进行关注或取关
// @Tags 扩展接口-II
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /relation/action [post]
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	if _, exist := usersLoginInfo[token]; !exist {
		ResponseError(c, CodeInvalidToken)
		return
	}
	//解析token获取user.id
	user, err := service.ParseToken(token)
	if err!=nil {
		zap.L().Error("ParseToken failed", zap.Error(err))
		ResponseError(c, CodeParseTokenFail)
		return
	}
	//获取relation对象即操作类型
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	toUserIdnum, err := strconv.Atoi(toUserId)
	//避免对自己的操作
	if int64(user.Id) == int64(toUserIdnum){
		ResponseErrorWithMsg(c,CodeOperateFail,"自己不要为难自己o")
		return
	}
	if err != nil {
		zap.L().Error("convert touserid from string to int64 failed", zap.String("touserid",  toUserId), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//构建关系操作结构体
	requestBody := &model.ActionRequest{
		UserId: int64(user.Id),
		Token: token,
		ToUserId: int64(toUserIdnum),
		ActionType: actionType,
	}
	operationFlag := false
	//type=1关注；type=2取关
	if actionType == "1"{
		//follow
		follow, err := service.RelationFollow(requestBody)
		if err != nil {//关注失败
			zap.L().Error("follow failed", zap.String("userid", string(requestBody.UserId)), zap.Error(err))
			ResponseError(c, CodeOperateFail)
			return
		}
		operationFlag = follow
	}else if actionType == "2"{
		//unfollow
		unfollow, err := service.RelationUnfollow(requestBody)
		if err != nil {//取关失败
			zap.L().Error("unfollow failed", zap.String("userid", string(requestBody.UserId)), zap.Error(err))
			ResponseError(c, CodeOperateFail)
			return
		}
		operationFlag = unfollow
	}
	if !operationFlag{
		//操作失败
		ResponseError(c, CodeOperateFail)
		return
	}
	//成功返回
	ResponseSuccess(c, nil)
}


// FollowList 获取关注列表
// @Summary 获取关注列表接口
// @Description 登录用户关注的所有用户列表
// @Tags 扩展接口-II
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /relation/follow/list/ [get]
func FollowList(c *gin.Context) {
	//token为空或不存在，返回err信息
	token := c.Query("token")
	if _, exist := usersLoginInfo[token]; !exist {
		ResponseError(c,CodeInvalidToken)
		return
	}
	userId := c.Query("user_id")
	userIdnum, err := strconv.Atoi(userId)
	if err != nil {
		zap.L().Error("convert touserid from string to int64 failed", zap.String("touserid",  userId), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//根据userid获取关注列表
	list, err := service.GetFollowList(int64(userIdnum))
	if err != nil {
		zap.L().Error("get follow list failed", zap.String("touserid",  userId), zap.Error(err))
		ResponseError(c, CodeOperateFail)
		return
	}
	//成功返回
	ResponseSuccess(c, list)
}


// FollowerList 获取粉丝列表
// @Summary 获取粉丝列表接口
// @Description 所有关注登录用户的粉丝列表
// @Tags 扩展接口-II
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /relation/follower/list/ [get]
func FollowerList(c *gin.Context) {
	//token为空或不存在，返回err信息
	token := c.Query("token")
	if _, exist := usersLoginInfo[token]; !exist {
		ResponseError(c, CodeInvalidToken)
		return
	}
	userId := c.Query("user_id")
	userIdnum, err := strconv.Atoi(userId)
	if err != nil {
		zap.L().Error("convert touserid from string to int64 failed", zap.String("touserid",  userId), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//根据userid获取粉丝列表
	list, err := service.GetFollowerList(int64(userIdnum))
	if err != nil {
		zap.L().Error("get fans list failed", zap.String("touserid",  userId), zap.Error(err))
		ResponseError(c, CodeOperateFail)
		return
	}
	//成功返回
	ResponseSuccess(c, list)
}
