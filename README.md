# CloudDisk

> 轻量级云盘系统，基于go-zero、xorm实现
## 主要功能
   - 用户模块
     - 密码登录
     - 刷新Authorization
     - 邮箱注册
     - 用户详情
   - 存储模块
     - 中心存储资源管理
     - 文件上传
     - 文件删除
     - 文件移动
     - 文件重命名
     - 文件夹创建 
   - 个人存储资源管理模块
     - 文件关联存储
     - 文件列表
     - 文件名称修改
     - 文件夹创建
     - 文件删除
     - 文件移动
   - 文件分享模块
     - 创建分享记录
     - 分享文件资源详情
     - 分享文件资源存储

## 相关命令
```text
# 创建API服务
goctl api new core
# 启动服务
go run core.go -f etc/core-api.yaml
# 使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero
# minio docker 启动
docker run -dt  \
   -p 9000:9000 \
   -p 9091:9091 \
   --name minio \
   -v /Users/luke/Desktop/minio/data:/mnt/data \
   -e "MINIO_ROOT_USER=admin" \
   -e "MINIO_ROOT_PASSWORD=12345678"    \
   minio/minio server /data     \
   --console-address ":9091"

```