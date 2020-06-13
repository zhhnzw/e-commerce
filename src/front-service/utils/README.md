### 工具层

#### ① 缓存热门数据(redis string类型)：

缓存逻辑：缓存最近1天内访问次数大于1次的链接

具体实现：在中间件middleware_visit中记录访问量，
在中间件middleware_cache中根据访问量判断是否需要操作缓存

操作缓存逻辑(缓存逻辑只操作get请求)：

数据流首先从middleware_cache进来，若访问量少于等于1次，此中间件处理结束

若访问量大于1次做缓存操作，若缓存存在，则直接返回缓存，若没有缓存就走缓存数据的操作

缓存数据操作：在controller中流程处理完后把response set到context，在中间件middleware_cache中做后续操作，从context get到数据，写入redis

#### ② 历史访问量的实现(redis string类型)：

在中间件middleware_visit中对string类型做incr操作