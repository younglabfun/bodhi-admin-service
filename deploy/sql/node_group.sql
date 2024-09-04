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

 Date: 19/04/2024 21:24:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for node_group
-- ----------------------------
DROP TABLE IF EXISTS `node_group`;
CREATE TABLE `node_group` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '功能分组ID',
  `title` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分组名称',
  `name` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分组id',
  `sort` int unsigned DEFAULT '0' COMMENT '排序',
  `is_deleted` tinyint unsigned DEFAULT '0' COMMENT '是否删除',
  `created_at` int unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='功能分组表';

SET FOREIGN_KEY_CHECKS = 1;
