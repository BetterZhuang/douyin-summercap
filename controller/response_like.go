package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseDataForLike struct {
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg,omitempty"`
	VideoList  []model.Video `json:"video_list"`
}

func ResponseErrorForLike(c *gin.Context, code int32) {
	c.JSON(http.StatusOK, &ResponseDataForLike{
		StatusCode: code,
		StatusMsg:  Code2Msg(code),
	})
}

func ResponseSuccessForLike(c *gin.Context, code int32) {
	c.JSON(http.StatusOK, &ResponseDataForLike{
		StatusCode: code,
		StatusMsg:  Code2Msg(code),
	})
}

func ResponseSuccessForLikeList(c *gin.Context, data []model.Video) {
	c.JSON(http.StatusOK, &ResponseDataForLike{
		StatusCode: CodeSuccess,
		StatusMsg:  Code2Msg(CodeSuccess),
		VideoList:  data,
	})
}
