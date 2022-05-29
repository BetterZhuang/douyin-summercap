# simple-demo

Mysql配置在config/config.go中修改

## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

工程无其他依赖，直接编译运行即可

```shell
go build && ./simple-demo
```

### 功能说明

接口功能不完善，仅作为示例

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试数据

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试

###5/29日完成扩展接口-II
0.关系操作、获取关注列表、获取粉丝列表\
1.获取关注或粉丝列表时  ->fix(初始化切片列表长度为0)\
2.修改关注数和粉丝数，正确对isfollow赋值，更正互关情况\
3.Response规范化，code和msg，规范代码\
4.使用zap日志记录请求及中间信息\
5.添加事务，执行关注或取关时，同步更新follow_count和follower_count值\
6.使用swagger生成全局接口文档

###文件介绍
/config/config.yaml配置文件，后续涉及到配置可全部整合到该文件
/controller/code.go response.go    response规范化
/dao 数据层操作
/logger zap日志相关
/model 定义的相关结构体
/settings viper读取config配置 

###使用swagger生成全局接口文档步骤
参考李文周博客 ```https://www.liwenzhou.com/posts/Go/golang-menu/```

1.按格式写接口注释\
2.安装swag工具 ```go install github.com/swaggo/swag/cmd/swag```\
3.```swag init```生成接口文档```./docs```目录下\
4.在项目代码中注册路由的地方引入\
5.运行项目后，打开浏览器访问```http://localhost:8080/swagger/index.html``` 就能看到Swagger 2.0 Api文档
