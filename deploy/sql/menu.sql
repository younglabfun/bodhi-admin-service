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

 Date: 24/04/2024 21:53:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `pid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父id',
  `type` tinyint DEFAULT '0' COMMENT '菜单类型，0管理后台，1前台',
  `title` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单名',
  `func_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '功能标识，用于控制权限',
  `route` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由名',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `href` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '链接',
  `sort` int unsigned DEFAULT '0' COMMENT '菜单排序',
  `is_show` tinyint DEFAULT '1' COMMENT '是否显示',
  `is_enabled` tinyint DEFAULT '1' COMMENT '是否启用',
  `is_deleted` tinyint DEFAULT '0' COMMENT '是否删除',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '添加时间(毫秒)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code_index` (`func_code`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单表';

SET FOREIGN_KEY_CHECKS = 1;
