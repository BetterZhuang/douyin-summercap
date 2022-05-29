/**
    @author: zzg
    @date: 2022/5/26 15:22
    @dir_path: service
    @note:
**/

package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
	"log"
)

// RelationFollow 执行关注操作
func RelationFollow(request *model.ActionRequest) (bool, error) {
	//先检查是否已关注
	followed := dao.SearchRelationById(request.UserId, request.ToUserId)
	if followed>0{//已关注
		return false, errors.New("you have followed")
	}
	//follow操作
	inserted := dao.InsertRelation(request.UserId,request.ToUserId)
	if inserted<1{//follow失败
		return false, errors.New("insert this relation failed")
	}
	//操作成功返回true，nil
	return true, nil
}

// RelationUnfollow 执行取关操作
func RelationUnfollow(request *model.ActionRequest) (bool, error)  {
	//检查是否已关注
	followed := dao.SearchRelationById(request.UserId, request.ToUserId)
	if followed<0{//已关注
		return false, errors.New("you have not followed, needn't to unfollow")
	}
	//unfollow操作
	deleted := dao.DelRelation(request.UserId, request.ToUserId)
	//操作成功返回true，nil
	if !deleted{//unfollow 失败
		return false, errors.New("delete this relation failed")
	}
	//操作成功返回true，nil
	return true, nil
}

// GetFollowList 获取userid的关注列表
func GetFollowList(userid int64) (userList []model.User, err error) {
	//获取关注的id_List
	followList := dao.SelectFollowList(userid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	userList = make([]model.User, 0, len(followList))

	for i:=0;i<len(followList);i++{
		//根据userid获取具体信息
		user := dao.SelectUserInfoById(followList[i])
		userList = append(userList, model.User{
			Id: user.Id,
			Name: user.Name,
			FollowCount: int64(user.FollowCount),
			FollowerCount: int64(user.FollowerCount),
			IsFollow: true, //关注列表中的必然是当前id已关注的，恒为true
		})
	}
	return
}

// GetFollowerList 获取userid的粉丝列表
func GetFollowerList(userid int64) (userList []model.User, err error) {
	//获取粉丝的id_List
	followList := dao.SelectFollowerList(userid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	userList = make([]model.User, 0, len(followList))
	for i:=0;i<len(followList);i++{
		//根据userid获取具体信息
		user := dao.SelectUserInfoById(followList[i])
		//判断是否互关
		followed := dao.SearchRelationById(userid, int64(user.Id))
		var mutualFollow bool
		if followed>0{
			mutualFollow = true
		}else{
			mutualFollow = false
		}
		userList = append(userList, model.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   int64(user.FollowCount),
			FollowerCount: int64(user.FollowerCount),
			IsFollow:      mutualFollow,
		})
	}
	return
}


/*
done:
	1.获取关注或粉丝列表时，会有空，之后才是真正关注或粉丝  ->fix(初始化列表长度为0)
	2.完善isfollow情况，正确对isfollow赋值，互关情况
	3.response规范化，code和msg，规范代码
	4.zap日志
	5.添加事务，执行关注或取关时，同步更新follow_count和follower_count值
 */