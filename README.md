微服务模式重构 shihuo 中...

已经重构好商品服务goods和前台api服务层front-service

#### 部署：

安装：docker && docker-compose

留意 各个服务conf文件夹下的config.yaml

进入build目录

① 一键部署启动命令：docker-compose up -d

② 移除命令：docker-compose down --rmi='all'