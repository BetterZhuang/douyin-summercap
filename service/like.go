package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

func VideoLike(userId, videoId int64, actionType int32) (bool, error) {

	// 先检查是否已经点赞
	liked := dao.SearchLikeById(userId, videoId)
	if actionType == 1 {
		if liked > 0 { // 已经点赞
			return true, nil
		}
		inserted := dao.InsertLike(userId, videoId)
		if inserted < 1 { // 点赞失败
			return false, errors.New("insert this like failed")
		}
	} else {
		if liked <= 0 { // 没有点赞过不
			return true, nil
		}
		deleted := dao.DelLike(userId, videoId)
		if !deleted { // 取消点赞失败
			return false, errors.New("delete this like failed")
		}
	}
	return true, nil
}

func GetLikeList(userid int64) ([]model.Video, error) {
	// 获取点赞视频的列表video_id list
	likeList := dao.SelectLikeList(userid)

	videoList := make([]model.Video, 0, len(likeList))
	for i := 0; i < len(likeList); i++ {
		videodb := model.VideoDB{} // 被点赞的视频
		config.Db.Model(&model.VideoDB{}).Where("id = ?", likeList[i]).Find(&videodb)

		video_userid := videodb.AuthorId
		userdb := model.UserDB{} // 被点赞视频的作者
		config.Db.Model(&model.UserDB{}).Where("id = ?", video_userid).First(&userdb)

		is_follow := false                         // 当前用户 是否关注 被点赞用户
		followList := dao.SelectFollowList(userid) // 获取当前用户的关注列表
		for i := 0; i < len(followList); i++ {
			if followList[i] == int64(video_userid) {
				is_follow = true
				break
			}
		}

		user := model.User{
			Id:            userdb.Id,
			Name:          userdb.Name,
			FollowCount:   int64(userdb.FollowCount),
			FollowerCount: int64(userdb.FollowerCount),
			IsFollow:      is_follow,
		}

		videoList = append(videoList, model.Video{
			Id:            videodb.Id,
			Author:        user,
			PlayUrl:       videodb.PlayUrl,
			CoverUrl:      videodb.CoverUrl,
			FavoriteCount: videodb.FavoriteCount,
			CommentCount:  videodb.CommentCount,
			IsFavorite:    true,
		})
	}
	return videoList, nil
}
