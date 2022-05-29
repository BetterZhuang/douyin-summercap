package service

import (
	"log"
	"os/exec"
	"time"

	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/model"
)

func Publish(saveFile string, fileName string, user model.User) (bool, string) {
	// Get Video cover
	var count int64
	config.Db.Model(&model.VideoDB{}).Where("author_id = ? AND play_url = ?", user.Id, config.SavePath+fileName).Count(&count)
	if count > 0 {
		return false, "File already uploaded!"
	}

	jpgName := fileName + ".jpg"
	saveJPG := saveFile + ".jpg"
	cmd := exec.Command("ffmpeg", "-i", saveFile, "-vframes", "1", saveJPG)
	if err := cmd.Run(); err != nil {
		log.Println("run false")
		return false, err.Error()
	}
	log.Println("run false after")
	postVideo := model.VideoDB{
		AuthorId:      user.Id,
		PlayUrl:       config.SavePath + fileName,
		CoverUrl:      config.SavePath + jpgName,
		FavoriteCount: 0,
		CommentCount:  0,
		CreateTime:    time.Now().Unix(),
	}
	row := config.Db.Create(&postVideo).RowsAffected
	if row == 0 {
		return false, "Fail to upload video to database!"
	}
	return true, ""
}

func PublishList(id uint64) []model.Video {
	var videos []model.VideoDB
	config.Db.Model(&model.VideoDB{}).Where("author_id = ?", id).Find(&videos)
	videoInfos := make([]model.Video, 0)
	for _, video := range videos {
		var author model.User
		config.Db.Model(&model.UserDB{}).Where("id = ?", video.AuthorId).First(&author)

		videoInfos = append(videoInfos, model.Video{
			Id:            video.Id,
			Author:        author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
		})
	}
	return videoInfos
}
