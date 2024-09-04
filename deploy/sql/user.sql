/*
 Navicat Premium Data Transfer

 Source Server         : docker
 Source Server Type    : MySQL
 Source Server Version : 80028 (8.0.28)
 Source Host           : localhost:3306
 Source Schema         : bodhi_v2

 Target Server Type    : MySQL
 Target Server Version : 80028 (8.0.28)
 File Encoding         : 65001

 Date: 13/05/2024 22:17:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_uuid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户UUID',
  `username` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Email',
  `name` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `avatar` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '头像',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '备注',
  `last_login_ip` varchar(16) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登陆IP',
  `last_login_time` int unsigned NOT NULL DEFAULT '0' COMMENT '最后登陆时间',
  `last_active_ip` varchar(16) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后活跃ip',
  `last_active_time` int unsigned NOT NULL DEFAULT '0' COMMENT '最后活跃时间（毫秒）',
  `custom_data` text COLLATE utf8mb4_general_ci  COMMENT '用户数据json',
  `mail_verified` tinyint unsigned DEFAULT '0' COMMENT '邮箱验证',
  `is_deleted` tinyint unsigned DEFAULT '0' COMMENT '是否删除',
  `is_enabled` tinyint unsigned DEFAULT '1' COMMENT '是否启用',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '添加时间(毫秒)',
  PRIMARY KEY (`user_uuid`),
  UNIQUE KEY `user_email` (`email`) USING BTREE,
  KEY `user_uuid` (`user_uuid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_uuid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户UUID',
  `role_uuid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色UUID',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '添加时间(毫秒)',
  PRIMARY KEY (`id`),
  KEY `user_role` (`user_uuid`,`role_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户角色关系表';

-- ----------------------------
-- Table structure for user_token
-- ----------------------------
DROP TABLE IF EXISTS `user_token`;
CREATE TABLE `user_token` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_uuid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户UUID',
  `refresh_token` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新token',
  `expires_time` int unsigned NOT NULL DEFAULT '0' COMMENT '过期时间（秒）',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '添加时间(毫秒)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_token` (`user_uuid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='token表';

SET FOREIGN_KEY_CHECKS = 1;
