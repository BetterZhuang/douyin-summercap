package dao

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/model"
	"go.uber.org/zap"
	"time"
)

func SearchLikeById(userId, videoId int64) int64 {
	temp := model.LikeDB{}
	config.Db.Where("user_id = ? and video_id = ?", userId, videoId).Find(&temp)
	return temp.Id
}

func InsertLike(userId, videoId int64) int64 {
	like := model.LikeDB{
		UserId:     userId,
		VideoId:    videoId,
		CreateTime: time.Now(),
	}
	tx := config.Db.Begin() //开启事务
	res := tx.Create(&like)
	if res.Error != nil {
		zap.L().Error("DB insertLike failed.", zap.String("userid", string(userId)), zap.Error(res.Error))
		tx.Rollback()
		return 0
	}
	tx.Exec("update `video_db` set `favorite_count`=`favorite_count`+1 where `id`=?", videoId) //视频点赞数+1
	tx.Commit()                                                                                //提交事务

	return res.RowsAffected
}

func DelLike(userId, videoId int64) bool {
	tx := config.Db.Begin() //开启事务
	res := tx.Where("user_id = ? and video_id = ?", userId, videoId).Delete(&model.LikeDB{})
	if res.Error != nil {
		zap.L().Error("DB DeleteLike failed.", zap.String("userid", string(userId)), zap.Error(res.Error))
		tx.Rollback()
		return false
	}
	tx.Exec("update `video_db` set `favorite_count`=`favorite_count`-1 where `id`=?", videoId) //视频点赞数-1
	tx.Commit()                                                                                //提交事务

	return true
}

func SelectLikeList(id int64) (res []int64) { // 用户id的所有点赞视频
	var Temp []model.LikeDB
	config.Db.Where("user_id = ?", id).Find(&Temp)
	for _, v := range Temp {
		res = append(res, v.VideoId)
	}
	return res
}
