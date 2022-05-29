package model

// VideoDB 数据库Video数据映射模型
type VideoDB struct {
	Id            uint64 `gorm:"column:id;autoIncrement;primaryKey"`
	AuthorId      uint64 `gorm:"column:author_id"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount uint64 `gorm:"column:favorite_count"`
	CommentCount  uint64 `gorm:"column:comment_count"`
	CreateTime    int64  `gorm:"column:create_time"`
}

// Video 返回给用户的Video信息
type Video struct {
	Id            uint64 `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint64 `json:"favorite_count"`
	CommentCount  uint64 `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}
