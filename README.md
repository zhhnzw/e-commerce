https://1024cx.top/ <br/>
账号密码: guest/guest12345678

### 部署方式一（kubernetes，也是我当前的部署方式）：

nginx 直接在宿主机上启动了，其他服务都部署在 k8s里，本项目共设计了3个Pod，mysql、redis、e-commerce，其中 e-commerce 包含了4个container，即 web-backend、backend-service、goods、order，本部署方式只需要操作yaml文件即可，配置文件见build目录。

① 把4个本地镜像（web-backend、backend-service、goods、order）推送到[阿里云镜像仓库](https://cr.console.aliyun.com/cn-hangzhou/instances/repositories)

② 把3个Pod的yaml文件传到服务器上并启动服务，`kubectl apply -f mysql.yaml`，`kubectl apply -f redis.yaml`，`kubectl apply -f e-commerce.yaml`

③ 检查服务状态`kubectl get pods,svc`，e-commerce 的 `STATUS`为`Running`，`READY`为`4/4`即为正常

④ 更新服务时，本地推送更新的镜像到阿里云镜像仓库（image打上新的版本tag），在服务器上编辑e-commerce.yaml更新对应container的镜像tag版本，然后执行`kubectl apply -f e-commerce.yaml`即可升级版本

### 部署方式二（docker-compose）：

本部署方式可用来本地测试，需要把编译后的可执行文件移动到build目录，golang后端服务把可执行文件放到build内对应服务文件夹下，前端服务编译后的build文件夹，放到 build/web-backend 下，创建网络: docker network create outside，然后进入本项目的build目录进行部署操作。

安装：docker && docker-compose<br/>
留意 各个服务conf文件夹下的config.yaml

进入build目录<br/>
① 一键部署启动命令：docker-compose up -d <br/>
② 移除命令：docker-compose down --rmi='all'

### 开发进度笔记

本demo旨在以电商为业务背景实践常用技术

v0.1包含的子服务: <br/>
① 管理后台api服务(golang gin)<br/>
② 商品服务(含库存)(golang grpc)<br/>
③ 订单服务(golang grpc)<br/>
④ 管理后台web服务(react实现)

已实现的技术要点:<br/>① gin提供api服务，zap日志库，gorm，redis-go，viper配置文件，swagger生成接口文档<br/>
② 插入了1000万条商品记录, 常见的sql优化<br/>
③ redis缓存的实现，及其他利用redis实现的功能<br/>
  &ensp;(如记录商品访问量、商品访问榜单功能)<br/>
④ grpc服务的实现, api服务调用grpc子服务，核心服务由各个grpc子服务提供<br/>
⑤ 定时任务的实现(cron, 定时随机生成订单)<br/>
⑥ web前端技术的应用<br/>
⑦ docker-compose <br/>
⑧ kubernetes

预备下一步更新的技术点:<br/>
① gin项目优化，引入swagger<br/>
② 脚本定时更新热门数据缓存(如本项目各模块的首页数据)<br/>
③ 补充更多单元测试, 提高代码覆盖率<br/>
④ 下订单的操作修正为通过分布式事务实现<br/>
⑤ gRPC负载均衡<br/>
⑥ 添加支付子服务