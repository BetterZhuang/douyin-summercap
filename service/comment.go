package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

func AddComment(request *model.CommentRequest) (model.Comment, error) {
	commentId := dao.InsertComment(request) // 插入一条评论记录

	userDb := model.UserDB{}
	config.Db.Model(&model.UserDB{}).Where("id = ?", request.UserId).First(&userDb)
	user := model.User{
		Id:            userDb.Id,
		Name:          userDb.Name,
		FollowCount:   int64(userDb.FollowCount),
		FollowerCount: int64(userDb.FollowCount),
		IsFollow:      userDb.IsFollow,
	}

	commentDb := dao.SearchCommentById(commentId)
	createTime := fmt.Sprintf("%02d-%02d", commentDb.CreateTime.Month(), commentDb.CreateTime.Day())

	comment := model.Comment{
		Id:          commentId,
		User:        user,
		CommentText: request.CommentText,
		CreateDate:  createTime,
	}

	return comment, nil
}

func DeleteComment(commentId, videoId int64) bool {
	res := dao.DeleteComment(commentId, videoId)
	return res
}

func GetCommentList(videoId int64) ([]model.Comment, error) {
	commentIdList := dao.GetCommentList(videoId)

	commentList := make([]model.Comment, 0, len(commentIdList))
	for i := 0; i < len(commentIdList); i++ {
		commentDb := model.CommentDB{} // 一条评论
		config.Db.Model(&model.CommentDB{}).Where("id = ?", commentIdList[i]).Find(&commentDb)

		userId := commentDb.UserId
		userDb := model.UserDB{}
		config.Db.Model(&model.UserDB{}).Where("id = ?", userId).First(&userDb)
		user := model.User{
			Id:            userDb.Id,
			Name:          userDb.Name,
			FollowCount:   int64(userDb.FollowCount),
			FollowerCount: int64(userDb.FollowCount),
			IsFollow:      userDb.IsFollow,
		}

		createTime := fmt.Sprintf("%02d-%02d", commentDb.CreateTime.Month(), commentDb.CreateTime.Day())

		comment := model.Comment{
			Id:          commentDb.Id,
			User:        user,
			CommentText: commentDb.CommentText,
			CreateDate:  createTime,
		}
		commentList = append(commentList, comment)
	}

	return commentList, nil
}
