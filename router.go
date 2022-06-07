package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"

	_ "github.com/RaymondCode/simple-demo/docs" // 千万不要忘了导入上一步生成的docs
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	v2 := r.Group("/douyin/v2")

	// basic apis
	v2.GET("/feed/", controller.Feed)
	v2.GET("/user/", controller.UserInfo)
	v2.POST("/user/register/", controller.Register)
	v2.POST("/user/login/", controller.Login)
	v2.POST("/publish/action/", controller.Publish)
	v2.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	v2.POST("/favorite/action/", controller.FavoriteAction, middleware.RateLimitMiddleware(2*time.Second, 1))
	v2.GET("/favorite/list/", controller.FavoriteList)
	v2.POST("/comment/action/", controller.CommentAction)
	v2.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	v2.POST("/relation/action/", controller.RelationAction, middleware.RateLimitMiddleware(2*time.Second, 1))
	v2.GET("/relation/follow/list/", controller.FollowList)
	v2.GET("/relation/follower/list/", controller.FollowerList)
}
