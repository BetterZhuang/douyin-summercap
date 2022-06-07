## 8888组作品-抖音项目服务端

### 功能说明
1. 视频feed流、视频投稿、个人信息
2. 点赞列表、用户评论
3. 关注列表、粉丝列表

详细接口说明文档见
```https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145```

### 模拟器
- Android Studio: 注意本地IP要设置为 10.0.2.2

### 文件介绍
/config/config.yaml 配置文件，后续涉及到配置可全部整合到该文件\
/controller 业务模块定义\
/service 逻辑接口定义\
/dao 数据访问层操作\
/public 视频保存目录\
/logger zap日志相关\
/model 定义的相关结构体\
/settings viper读取config配置\
/docs Swagger接口文档目录

### 使用swagger生成全局接口文档步骤
1.按格式写接口注释\
2.安装swag工具 ```go install github.com/swaggo/swag/cmd/swag```\
3.```swag init```生成接口文档```./docs```目录下\
4.在项目代码中注册路由的地方引入\
5.运行项目后，打开浏览器访问```http://localhost:8080/swagger/index.html``` 就能看到Swagger 2.0 Api文档
