/**
    @author: zzg
    @date: 2022/5/26 15:16
    @dir_path: model
    @note:
**/

package model

import "time"


type RelationDB struct {
	Id         int64     `gorm:"column:id;autoIncrement;primaryKey"`
	UserId     int64       `gorm:"column:user_id"`     //关注者
	ToUserId   int64       `gorm:"column:to_user_id"`  //被关注
	CreateTime time.Time `gorm:"column:create_time"`
}

type Relation struct {
	Id         int64     `json:"id"`
	UserId     int64       `json:"user_id"`
	ToUserId   int64       `json:"to_user_id"`
	CreateTime time.Time `json:"create_time"`
}

type ActionRequest struct {
	UserId int64          `json:"user_id"`
	Token string          `json:"token"`
	ToUserId int64        `json:"to_user_id"`
	ActionType string     `json:"action_type"`
}