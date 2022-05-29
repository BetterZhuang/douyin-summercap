package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 10001, //程序中的错误码
	"msg" : xx,   //提示信息
	"data" : {},  //数据
}
*/

type ResponseData struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserList []model.User `json:"user_list"`
	//Video []model.Video `json:"video_list"`
}

//定义方法

func ResponseError(c *gin.Context, code int32) {
	c.JSON(http.StatusOK, &ResponseData{
		StatusCode: code,
		StatusMsg:  Code2Msg(code),
		UserList: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code int32, msg string) {
	c.JSON(http.StatusOK, &ResponseData{
		StatusCode: code,
		StatusMsg:  msg,
		UserList: nil,
	})
}

func ResponseSuccess(c *gin.Context, data []model.User) {
	c.JSON(http.StatusOK, &ResponseData{
		StatusCode: CodeSuccess,
		StatusMsg:  Code2Msg(CodeSuccess),
		UserList: data,
	})
}
