/**
    @author: zzg
    @date: 2022/5/26 15:32
    @dir_path: dao
    @note:
**/

package dao

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/model"
	"go.uber.org/zap"
	"time"
)

func SearchRelationById(userId, toUserId int64) int64 {
	temp := model.RelationDB{}
	config.Db.Where("user_id = ? and to_user_id = ?", userId, toUserId).Find(&temp)
	return temp.Id
}

func InsertRelation(userId, toUserId int64) int64 {
	relation := model.RelationDB{
		UserId:     userId,
		ToUserId:   toUserId,
		CreateTime: time.Now(),
	}
	tx := config.Db.Begin() //开启事务
	res := tx.Create(&relation)
	if res.Error != nil {
		zap.L().Error("DB insertRelation failed.", zap.String("userid", string(userId)), zap.Error(res.Error))
		tx.Rollback()
		return 0
	}
	tx.Exec("update `user_db` set `follow_count`=`follow_count`+1 where `id`=?", userId)       //关注+1
	tx.Exec("update `user_db` set `follower_count`=`follower_count`+1 where `id`=?", toUserId) //粉丝+1
	tx.Commit()                                                                                //提交事务

	return res.RowsAffected
}

func DelRelation(userId, toUserId int64) bool {
	tx := config.Db.Begin() //开启事务
	res := tx.Where("user_id = ? and to_user_id = ?", userId, toUserId).Delete(&model.RelationDB{})
	if res.Error != nil {
		zap.L().Error("DB DeleteRelation failed.", zap.String("userid", string(userId)), zap.Error(res.Error))
		tx.Rollback()
		return false
	}
	tx.Exec("update `user_db` set `follow_count`=`follow_count`-1 where `id`=?", userId)       //关注-1
	tx.Exec("update `user_db` set `follower_count`=`follower_count`-1 where `id`=?", toUserId) //粉丝-1
	tx.Commit()                                                                                //提交事务

	return true
}

func SelectFollowList(id int64) (res []int64) {
	var Temp []model.RelationDB
	config.Db.Where("user_id = ?", id).Find(&Temp)
	for _, v := range Temp {
		res = append(res, v.ToUserId)
	}
	return res
}

func SelectFollowerList(id int64) (res []int64) {
	var Temp []model.RelationDB
	config.Db.Where("to_user_id = ?", id).Find(&Temp)
	for _, v := range Temp {
		res = append(res, v.UserId)
	}
	return res
}

func SelectUserInfoById(id int64) model.UserDB {
	var res model.UserDB
	config.Db.Where("id = ?", id).Find(&res)
	return res
}
