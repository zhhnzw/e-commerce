CREATE TABLE IF NOT EXISTS `tb_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(32) NOT NULL COMMENT '用户名,账号',
  `nick_name` varchar(32) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` char(60) NOT NULL COMMENT '加密后的密码',
  `mobile` char(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(32) NOT NULL DEFAULT '' COMMENT '邮箱',
  `is_valid` boolean NOT NULL DEFAULT 1 COMMENT '是否可用',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像地址',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE `user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

BEGIN;
INSERT INTO `tb_user` VALUES (1, '26394826', '卖萌小公举', '$2a$10$UXi2k0blmu3ImJvBv6KcTuqw.2bEEJq/qGz5TLb0Y94SY.qEwJLxK', '13543434343', '2131@11.com', 1, '', '2020-07-05 18:47:20', '2020-07-07 20:25:47', '1970-01-01 08:00:01');
INSERT INTO `tb_user` VALUES (2, '127308924', '发单小公举', '$2a$10$8YKsoHkz62isUI3yBIfshO0/g0rZfMMJGF.7szasR.HCJ20lUxPyu', '13543434343', '2131@11.com', 1, '', '2020-07-05 18:49:28', '2020-07-07 20:25:46', '1970-01-01 08:00:01');
INSERT INTO `tb_user` VALUES (3, '381936289', '接单小公举', '$2a$10$/87bIAON2Gs0HA/9/fZsX.NMKcPDwtSLLi1GXSaS4dyGz.VzYKDZ6', '13543434343', '2131@11.com', 1, '', '2020-07-05 18:50:22', '2020-07-07 20:24:30', '1970-01-01 08:00:01');
INSERT INTO `tb_user` VALUES (4, '434247532', '外卖小公举', '$2a$10$/87bIAON2Gs0HA/9/fZsX.NMKcPDwtSLLi1GXSaS4dyGz.VzYKDZ6', '13543434343', '2131@11.com', 1, '', '2020-07-07 20:25:26', '2020-07-07 20:25:42', '1970-01-01 08:00:01');
INSERT INTO `tb_user` VALUES (5, '534545642', '电击小公举', '$2a$10$/87bIAON2Gs0HA/9/fZsX.NMKcPDwtSLLi1GXSaS4dyGz.VzYKDZ6', '13543434343', '2131@11.com', 1, '', '2020-07-07 20:26:51', '2020-07-07 20:27:04', '1970-01-01 08:00:01');
COMMIT;