数据库商品表里现有1千万条生成的随机值记录,功能还在更新中

#### 访问链接示例：

[查询数据库记录数(目前只查商品表和商品类型表的记录数)](https://www.1024cx.top/count)

[查询商品列表(根据类别查询)](https://www.1024cx.top/v1/commodity?primaryType=clothes&secondaryType=shirt&pageSize=20&pageIndex=1)

[查询指定商品](https://www.1024cx.top/v1/commodity/2971f6ae-fc9a-48ad-9f0c-ace2aa75f9e2)

[查询商品榜单](https://www.1024cx.top/v1/hot?primaryType=clothes&pageSize=10&pageIndex=1)

#### 部署：

安装：docker && docker-compose

留意 conf文件夹下的config.yaml

① 一键部署启动命令：docker-compose up -d

② 移除命令：docker-compose down --rmi='all'