package model

import "time"

type CommentDB struct {
	Id          int64     `gorm:"column:id;autoIncrement;primaryKey"` // 评论id
	UserId      int64     `gorm:"column:user_id"`                     // 用户id
	VideoId     int64     `gorm:"column:video_id"`                    // 视频id
	CommentText string    `gorm:"column:content"`                     // 评论内容
	CreateTime  time.Time `gorm:"column:create_time"`                 // 评论时间
}

type Comment struct {
	Id          int64  `json:"id,omitempty"`          // 评论id
	User        User   `json:"user"`                  // 评论用户信息
	CommentText string `json:"content,omitempty"`     // 评论内容
	CreateDate  string `json:"create_date,omitempty"` // 评论时间
}

type CommentRequest struct {
	UserId      int64  `json:"user_id"`
	VideoId     int64  `json:"video_id"`
	CommentText string `json:"comment_text"`
}
