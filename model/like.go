package model

import "time"

type LikeDB struct {
	Id         int64     `gorm:"column:id;autoIncrement;primaryKey"`
	UserId     int64     `gorm:"column:user_id"`  // 用户id
	VideoId    int64     `gorm:"column:video_id"` // 视频id
	CreateTime time.Time `gorm:"column:create_time"`
}

type Like struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	VideoId    int64     `json:"video-id"`
	CreateTime time.Time `json:"create_time"`
}

type LikeActionRequest struct {
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
	VideoId    int64  `json:"video_id"`
	ActionType int32  `json:"action_type"`
}
