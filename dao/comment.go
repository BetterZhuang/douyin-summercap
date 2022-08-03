package dao

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/model"
	"go.uber.org/zap"
	"time"
)

func SearchCommentById(commentId int64) model.CommentDB {
	temp := model.CommentDB{}
	config.Db.Where("id = ?", commentId).Find(&temp)
	return temp
}

// InsertComment 添加评论
func InsertComment(request *model.CommentRequest) int64 {
	temp := model.CommentDB{
		UserId:      request.UserId,
		VideoId:     request.VideoId,
		CommentText: request.CommentText,
		CreateTime:  time.Now(),
	}
	tx := config.Db.Begin() //开启事务
	res := tx.Create(&temp)
	if res.Error != nil {
		zap.L().Error("DB insertComment failed.", zap.String("userid", string(request.UserId)), zap.Error(res.Error))
		tx.Rollback()
		return 0
	}
	tx.Exec("update `video_db` set `comment_count`=`comment_count`+1 where `id`=?", request.VideoId) //视频评论数+1
	tx.Commit()

	return temp.Id // returns inserted records count, via comment_id
}

// DeleteComment 删除评论
func DeleteComment(commentID, videoId int64) bool {
	temp := model.CommentDB{}
	tx := config.Db.Begin() //开启事务
	res := config.Db.Where("id = ?", commentID).Delete(&temp)
	if res.Error != nil {
		zap.L().Error("DB deleteComment failed.", zap.String("comment_id", string(commentID)), zap.Error(res.Error))
		tx.Rollback()
		return false
	}
	tx.Exec("update `video_db` set `comment_count`=`comment_count`-1 where `id`=?", videoId) //视频评论数-1
	tx.Commit()

	return true
}

func GetCommentList(videoId int64) (res []int64) { // 视频下的所有评论
	var temp []model.CommentDB
	config.Db.Where("video_id = ?", videoId).Order("create_time desc").Find(&temp)
	for _, v := range temp {
		res = append(res, v.Id)
	}
	return res
}
