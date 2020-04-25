##### 项目规范
1,每个目录 需要有独立的README.md  

2,以后每个业务或者基础组件维护自己的版本号

3,整个大仓库不再有tag，只有master 主干分支，所有mr发送前，一定要注意先merge master；

4,提供RPC内部服务放置在app/service中，任务队列放置在app/job中，对外网关服务放置在app/api，管理后台服务放置在app/admin

5,每个业务自建cmd文件夹,将main.go文件和test配置文件迁移进去

