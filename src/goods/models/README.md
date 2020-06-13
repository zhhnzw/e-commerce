### 数据层(通过gorm访问mysql)

#### ① 商品：

考虑到商品表是张大表，需要做查询优化

延迟关联，索引覆盖的应用，见方法：func (model *Commodity) GetCommoditiesByType()

其他单张大表的查询场景都可使用这个优化方案，比如用户表
