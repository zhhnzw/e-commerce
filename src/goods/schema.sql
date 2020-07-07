CREATE TABLE IF NOT EXISTS `tb_goods` (
  `id` int NOT NULL AUTO_INCREMENT,
  `goods_uuid` char(36) NOT NULL COMMENT '商品id',
  `goods_type_id` int NOT NULL DEFAULT 0 COMMENT '商品类型id',
  `primary_type` enum('unknown','clothes', 'pants', 'shoes') DEFAULT 'unknown' COMMENT '一级商品类型,如: clothes(衣服)、pants(裤子)、shoes(鞋子)',
  `secondary_type` enum('unknown','shirt','jacket','casual_pants', 'sports_pants', 'basketball_shoes','casual_shoes') DEFAULT 'unknown' COMMENT '二级商品类型,如: shirt(衬衫)、jacket(夹克)、casual_pants(休闲裤)、sports_pants(运动裤)、basketball_shoes(篮球鞋)、casual_shoes(休闲鞋)',
  `img` varchar(256) NOT NULL DEFAULT '' COMMENT '商品图片',
  `imgs` text NOT NULL COMMENT '商品详情图片',
  `publish_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '标题',
  `subtitle` varchar(256) NOT NULL DEFAULT '' COMMENT '子标题',
  `price` int NOT NULL DEFAULT 0 COMMENT '单位:厘. 存int类型精确到厘方便计算',
  `stock` int NOT NULL DEFAULT 0 COMMENT '库存',
  `is_valid` boolean NOT NULL DEFAULT 1 COMMENT '是否可用,用于下线和删除操作',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '删除时间, 逻辑删除',
  PRIMARY KEY (`id`),
  UNIQUE `unique_goods_uuid` (`goods_uuid`),
  KEY `goods_type_id` (`goods_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

CREATE TABLE IF NOT EXISTS `tb_goods_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `primary_type` enum('unknown','clothes', 'pants', 'shoes') DEFAULT 'unknown' COMMENT '一级商品类型,如: clothes(衣服)、pants(裤子)、shoes(鞋子)',
  `secondary_type` enum('unknown','shirt','jacket','casual_pants', 'sports_pants', 'basketball_shoes','casual_shoes') DEFAULT 'unknown' COMMENT '二级商品类型,如: shirt(衬衫)、jacket(夹克)、casual_pants(休闲裤)、sports_pants(运动裤)、basketball_shoes(篮球鞋)、casual_shoes(休闲鞋)',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '删除时间, 逻辑删除',
  PRIMARY KEY (`id`),
  UNIQUE `unique_primary_secondary` (`primary_type`, `secondary_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品类型表';

INSERT INTO `tb_goods_type` (`primary_type`, `secondary_type`) VALUES ('clothes', 'shirt'),('clothes', 'jacket'),('pants','casual_pants'),('pants','sports_pants'),('shoes','basketball_shoes'),('shoes','casual_shoes');

CREATE TABLE IF NOT EXISTS `tb_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_uuid` char(36) NOT NULL COMMENT '用户id',
  `nick_name` varchar(32) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` char(60) NOT NULL COMMENT '加密后的密码',
  `mobile` char(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(32) NOT NULL DEFAULT '' COMMENT '邮箱',
  `is_valid` boolean NOT NULL DEFAULT 1 COMMENT '是否可用',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像地址',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '删除时间, 逻辑删除',
  PRIMARY KEY (`id`),
  UNIQUE `unique_user_uuid` (`user_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE IF NOT EXISTS `tb_user_collection` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_uuid` char(36) NOT NULL COMMENT '用户id',
  `goods_uuid` char(36) NOT NULL COMMENT '商品id',
  PRIMARY KEY (`id`),
  UNIQUE `unique_user_goods_uuid` (`user_uuid`,`goods_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户收藏商品表';

CREATE TABLE IF NOT EXISTS `tb_article` (
  `id` int NOT NULL AUTO_INCREMENT,
  `author_uuid` char(36) NOT NULL COMMENT '作者id',
  `avatar` varchar(256) NOT NULL DEFAULT '' COMMENT '用户头像',
  `article_uuid` char(36) NOT NULL COMMENT '文章id',
  `content` text COMMENT '文章内容',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE `unique_author_uuid` (`author_uuid`),
  UNIQUE `unique_article_author_uuid` (`article_uuid`,`author_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

CREATE TABLE IF NOT EXISTS `tb_article_likes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_uuid` char(36) NOT NULL COMMENT '用户id',
  `article_uuid` char(36) NOT NULL COMMENT '文章id',
  PRIMARY KEY (`id`),
  UNIQUE `unique_user_article_uuid` (`user_uuid`,`article_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章点赞表';

CREATE TABLE IF NOT EXISTS `tb_comment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `comment_uuid` char(36) NOT NULL COMMENT '评论id',
  `article_uuid` char(36) NOT NULL COMMENT '文章id',
  `parent_uuid` char(36) NOT NULL DEFAULT '0' COMMENT '0代表为文章的直接评论，否则此条评论为回复评论',
  `comment_text` varchar(512) NOT NULL DEFAULT '' COMMENT '评论内容',
  `comment_user_id` int NOT NULL COMMENT '评论人的id',
  `comment_user_avatar` varchar(256) NOT NULL DEFAULT '' COMMENT '评论人的头像',
  `comment_user_name` varchar(32) NOT NULL DEFAULT '' COMMENT '评论人的昵称',
  `comment_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论时间',
  `to_user_id` int DEFAULT 0 COMMENT '回复目标人的id。如果不是回复类型的评论, 则为0',
  `to_user_name` varchar(32) NOT NULL DEFAULT '' COMMENT '回复目标人的昵称',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE `unique_comment_uuid` (`comment_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';