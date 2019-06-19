<h3 align="center">Pandora - 基于k8s的自动化发布部署工具</h3>

[预览](http://116.62.121.87:8878/)
- username： admin
- password： adminadmin

![](https://github.com/ielepro/pandora/blob/master/assert/dashboard.png)

![](https://github.com/ielepro/pandora/blob/master/assert/deploy.png)

## 特性

- Go语言开发，编译简单、运行高效
- Web界面访问，交互友好
- 权限模型灵活自由
- docker image 构建, 支持自定义构建
- k8s deployment 版本升级部署
- 支持Git仓库
- 支持分支、Tag上线
- 部署Hook支持，可扩展性强
- 完善的上线工作流
- 邮件通知机制

## 使用

1. 下载源码包，编译安装
2. 安装机器必须装有`git`，并配置免密登录
3. 将k8s集群上master的kube-config文件拷贝到当前目录下，命名为admin.conf
4. 登录web页面 创建项目，配置git仓库地址等，创建上线单，编写打包脚本。
5. 上线操作：打包成docker镜像，上传到仓库，然后部署，更新集群里符合的deployment镜像版本
 

## 特别鸣谢

   本项目基于以下项目开发：
   - [gorm](https://github.com/jinzhu/gorm)
   - [gin](https://github.com/gin-gonic/gin)
   - [syncd](https://syncd.cc/docs/#donate.md)
   - [client-go](https://github.com/kubernetes/client-go)
