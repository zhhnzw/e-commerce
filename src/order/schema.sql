CREATE TABLE IF NOT EXISTS `tb_order` (
  `id` int NOT NULL AUTO_INCREMENT,
  `goods_uuid` char(36) NOT NULL COMMENT '商品id',
  `goods_type_id` int NOT NULL DEFAULT 0 COMMENT '商品类型id',
  `primary_type` enum('unknown','clothes', 'pants', 'shoes') DEFAULT 'unknown' COMMENT '一级商品类型,如: clothes(衣服)、pants(裤子)、shoes(鞋子)',
  `secondary_type` enum('unknown','shirt','jacket','casual_pants', 'sports_pants', 'basketball_shoes','casual_shoes') DEFAULT 'unknown' COMMENT '二级商品类型,如: shirt(衬衫)、jacket(夹克)、casual_pants(休闲裤)、sports_pants(运动裤)、basketball_shoes(篮球鞋)、casual_shoes(休闲鞋)',
  `img` varchar(256) NOT NULL DEFAULT '' COMMENT '商品图片',
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '标题',
  `subtitle` varchar(256) NOT NULL DEFAULT '' COMMENT '子标题',
  `price` int NOT NULL DEFAULT 0 COMMENT '单位:厘. 存int类型精确到厘方便计算',
  `order_status` enum('new','closed','paid','completed') DEFAULT 'new' COMMENT '订单状态: new(新建)、closed(已关闭)、paid(已支付)、completed(已完成)',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '删除时间, 逻辑删除',
  PRIMARY KEY (`id`),
  KEY `goods_type_id` (`goods_type_id`),
  KEY `primary_secondary_type` (`secondary_type`, `primary_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';