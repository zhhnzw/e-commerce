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