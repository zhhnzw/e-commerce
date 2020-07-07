本demo旨在以电商为业务背景实践微服务

v0.1包含的子服务,进行中...: <br/>
① 前台web服务(不打算写)<br/>
② 前台api服务(golang gin)<br/>
③ 管理后台web服务(react实现)<br/>
④ 管理后台api服务(golang gin)<br/>
⑤ 商品服务(含库存)(golang grpc)<br/>
⑥ 订单服务(golang grpc)

已实现的技术要点:<br/>
① 插入了1000万条商品记录, 实践常见的sql优化<br/>
② redis缓存的实现，及其他利用redis实现的功能<br/>
  &ensp;(如记录商品访问量、商品访问榜单功能)<br/>
③ grpc服务的实现, api服务调用grpc子服务，核心服务由各个grpc子服务提供<br/>
④ 定时任务的实现(cron, 定时随机生成订单)<br/>
⑤ web前端技术的应用, react及相关类库antd、bizcharts等<br/>
⑥ docker-compose 一键部署

预备下一步更新的技术点:<br/>
① rpc连接池实现<br/>
② 脚本定时更新热门数据缓存(如本项目各模块的首页数据)<br/>
③ 服务注册与发现<br/>
④ 补充更多单元测试, 提高代码覆盖率

v1.0 预备更新的技术点:<br/>
① 下订单的操作修正为通过分布式事务实现<br/>
② 添加支付子服务
③ 部署引入k8s

#### 部署：
安装：docker && docker-compose<br/>
留意 各个服务conf文件夹下的config.yaml

进入build目录<br/>
① 一键部署启动命令：docker-compose up -d <br/>
② 移除命令：docker-compose down --rmi='all'

https://www.1024cx.top/ <br/>
账号密码:guest/guest12345678