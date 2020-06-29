CREATE TABLE IF NOT EXISTS `sys_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(32) NOT NULL COMMENT '用户名,账号',
  `nick_name` varchar(32) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` char(60) NOT NULL COMMENT '加密后的密码',
  `mobile` char(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(32) NOT NULL DEFAULT '' COMMENT '邮箱',
  `is_valid` boolean NOT NULL DEFAULT 1 COMMENT '是否可用',
  `is_super` boolean NOT NULL DEFAULT 0 COMMENT '是否是超级管理员',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像地址',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE `user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;