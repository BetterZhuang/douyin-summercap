/**
    @author: zzg
    @date: 2022/5/28 21:45
    @dir_path: controller
    @note:
**/

package controller

const (
	CodeSuccess int32 = iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
	CodeParseTokenFail
	CodeOperateFail

	CodeCommentNotExist
)

var codeMsgmap = map[int32]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeNeedLogin:      "需要登录",
	CodeInvalidToken:   "无效的Token",
	CodeParseTokenFail: "解析Token失败",
	CodeOperateFail:    "操作失败",

	CodeCommentNotExist: "评论内容为空",
}

// Code2Msg :返回code-msg
func Code2Msg(c int32) string {
	msg, ok := codeMsgmap[c]
	if !ok {
		msg = codeMsgmap[CodeServerBusy]
	}
	return msg
}
