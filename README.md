本demo旨在以电商为业务背景实践微服务

v0.1包含的子服务,进行中...: <br/>
①前台服务(不打算写)<br/>
②管理后台web服务(react实现)<br/>
③管理后台api服务(golang gin)<br/>
④商品服务(含库存)(golang grpc)<br/>
⑤订单服务(golang grpc)<br/>

#### 部署：
安装：docker && docker-compose<br/>
留意 各个服务conf文件夹下的config.yaml

进入build目录<br/>
① 一键部署启动命令：docker-compose up -d <br/>
② 移除命令：docker-compose down --rmi='all'