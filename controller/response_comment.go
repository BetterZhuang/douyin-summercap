package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseDataForComment struct {
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg,omitempty"`
	Comment    model.Comment `json:"comment,omitempty"`
}
type ResponseDataForCommentList struct {
	StatusCode  int32           `json:"status_code"`
	StatusMsg   string          `json:"status_msg,omitempty"`
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

func ResponseForComment(c *gin.Context, code int32) {
	c.JSON(http.StatusOK, &ResponseDataForComment{
		StatusCode: code,
		StatusMsg:  Code2Msg(code),
	})
}

func ResponseSuccessForComment(c *gin.Context, comment model.Comment) {
	c.JSON(http.StatusOK, &ResponseDataForComment{
		StatusCode: CodeSuccess,
		StatusMsg:  Code2Msg(CodeSuccess),
		Comment:    comment,
	})
}

func ResponseSuccessForCommentList(c *gin.Context, commentList []model.Comment) {
	c.JSON(http.StatusOK, &ResponseDataForCommentList{
		StatusCode:  CodeSuccess,
		StatusMsg:   Code2Msg(CodeSuccess),
		CommentList: commentList,
	})
}
