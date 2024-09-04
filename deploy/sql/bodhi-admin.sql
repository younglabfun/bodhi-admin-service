/*
 Navicat Premium Data Transfer

 Source Server         : docker-mysql8
 Source Server Type    : MySQL
 Source Server Version : 80028 (8.0.28)
 Source Host           : localhost:33060
 Source Schema         : bodhi-admin

 Target Server Type    : MySQL
 Target Server Version : 80028 (8.0.28)
 File Encoding         : 65001

 Date: 01/09/2024 20:05:15
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
  `is_enabled` tinyint DEFAULT '0' COMMENT '是否启用',
  `is_deleted` tinyint DEFAULT '0' COMMENT '是否删除',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '添加时间(毫秒)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code_index` (`func_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单表';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` (`id`, `pid`, `type`, `title`, `func_code`, `route`, `component`, `icon`, `href`, `sort`, `is_show`, `is_enabled`, `is_deleted`, `created_at`) VALUES (1, 0, 0, '工作台', 'Dashboard', 'Dashboard', 'dashboard/index', 'fa fa-dashboard', '/', 0, 1, 1, 0, 1724316970);
INSERT INTO `menu` (`id`, `pid`, `type`, `title`, `func_code`, `route`, `component`, `icon`, `href`, `sort`, `is_show`, `is_enabled`, `is_deleted`, `created_at`) VALUES (2, 0, 0, '系统配置', 'Setting', 'Setting', 'Layout', 'fa fa-cog', '/setting/menu-list', 9, 1, 1, 0, 1724379998);
INSERT INTO `menu` (`id`, `pid`, `type`, `title`, `func_code`, `route`, `component`, `icon`, `href`, `sort`, `is_show`, `is_enabled`, `is_deleted`, `created_at`) VALUES (3, 2, 0, '菜单', 'menu:list', 'MenuList', 'setting/menu-list', 'fa fa-list', 'menu-list', 1, 1, 1, 0, 1724380122);
INSERT INTO `menu` (`id`, `pid`, `type`, `title`, `func_code`, `route`, `component`, `icon`, `href`, `sort`, `is_show`, `is_enabled`, `is_deleted`, `created_at`) VALUES (4, 2, 0, '功能', 'node:list', 'NodeList', 'setting/node-list', 'fa fa-table', 'node-list', 2, 1, 1, 0, 1724404586);
INSERT INTO `menu` (`id`, `pid`, `type`, `title`, `func_code`, `route`, `component`, `icon`, `href`, `sort`, `is_show`, `is_enabled`, `is_deleted`, `created_at`) VALUES (5, 0, 0, 'Tree', 'Tree', 'Tree', 'tree/index', 'fa fa-tree', 'tree', 3, 0, 0, 0, 1724404626);
INSERT INTO `menu` (`id`, `pid`, `type`, `title`, `func_code`, `route`, `component`, `icon`, `href`, `sort`, `is_show`, `is_enabled`, `is_deleted`, `created_at`) VALUES (6, 0, 0, 'Form', 'Form', 'Form', 'form/index', 'fa fa-calendar', 'index', 5, 0, 0, 0, 1724404804);
INSERT INTO `menu` (`id`, `pid`, `type`, `title`, `func_code`, `route`, `component`, `icon`, `href`, `sort`, `is_show`, `is_enabled`, `is_deleted`, `created_at`) VALUES (7, 0, 0, '用户管理', 'User', 'User', 'Layout', 'fa fa-users', '/user/user-list', 7, 1, 1, 0, 1724404966);
INSERT INTO `menu` (`id`, `pid`, `type`, `title`, `func_code`, `route`, `component`, `icon`, `href`, `sort`, `is_show`, `is_enabled`, `is_deleted`, `created_at`) VALUES (8, 7, 0, '用户', 'user:list', 'UserList', 'user/user-list', 'fa fa-user', 'user-list', 0, 1, 1, 0, 1724405103);
INSERT INTO `menu` (`id`, `pid`, `type`, `title`, `func_code`, `route`, `component`, `icon`, `href`, `sort`, `is_show`, `is_enabled`, `is_deleted`, `created_at`) VALUES (9, 7, 0, '角色', 'role:list', 'RoleList', 'user/role-list', 'fa fa-user-secret', 'role-list', 1, 1, 1, 0, 1724405192);
INSERT INTO `menu` (`id`, `pid`, `type`, `title`, `func_code`, `route`, `component`, `icon`, `href`, `sort`, `is_show`, `is_enabled`, `is_deleted`, `created_at`) VALUES (10, 0, 0, 'external-link', 'External-link', 'External-link', 'layout', '', 'https://panjiachen.github.io/vue-element-admin-site/#/', 7, 0, 0, 0, 1724405259);
COMMIT;

-- ----------------------------
-- Table structure for node
-- ----------------------------
DROP TABLE IF EXISTS `node`;
CREATE TABLE `node` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '功能ID',
  `group_id` int unsigned NOT NULL DEFAULT '0' COMMENT '应用功能分组ID',
  `func_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '功能标识',
  `name` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '功能名称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '功能描述',
  `is_enabled` tinyint unsigned DEFAULT '1' COMMENT '是否启用',
  `is_deleted` tinyint unsigned DEFAULT '0' COMMENT '是否删除',
  `created_at` int unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `func_code` (`func_code`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='应用功能表';

-- ----------------------------
-- Records of node
-- ----------------------------
BEGIN;
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (1, 1, 'user:remove', '删除用户', '删除用户权限', 1, 0, 1724576757);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (2, 1, 'user:edit', '编辑用户', '编辑用户权限', 1, 0, 1724580800);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (3, 1, 'user:create', '添加用户', '添加用户权限\n', 1, 0, 1724580811);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (4, 4, 'T:action', 'taction', 'test action', 0, 0, 1724580823);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (5, 0, 'Dashboard', '工作台', '工作台页面', 1, 0, 1725158526);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (6, 1, 'user:list', '用户管理', '管理用户权限', 1, 0, 1725186018);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (7, 1, 'role:list', '角色管理', '角色管理', 1, 0, 1725186061);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (8, 5, 'node:list', 'node', '', 1, 0, 1725186616);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (9, 5, 'menu:list', '菜单列表', '', 1, 0, 1725191009);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (10, 5, 'menu:create', '添加菜单', '', 1, 0, 1725191066);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (11, 1, 'role:remove', '删除角色', '', 1, 0, 1725191083);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (12, 5, 'menu:remove', '删除菜单', '', 1, 0, 1725191122);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (13, 5, 'menu:edit', '编辑菜单', '', 1, 0, 1725191183);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (14, 1, 'role:create', '添加角色', '', 1, 0, 1725191721);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (15, 1, 'role:edit', '编辑角色', '', 1, 0, 1725191743);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (16, 5, 'node:create', '添加功能', '', 1, 0, 1725191807);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (17, 5, 'node:edit', '编辑功能', '', 1, 0, 1725191828);
INSERT INTO `node` (`id`, `group_id`, `func_code`, `name`, `description`, `is_enabled`, `is_deleted`, `created_at`) VALUES (18, 5, 'node:remove', '删除功能', '', 1, 0, 1725191882);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='功能分组表';

-- ----------------------------
-- Records of node_group
-- ----------------------------
BEGIN;
INSERT INTO `node_group` (`id`, `title`, `name`, `sort`, `is_deleted`, `created_at`) VALUES (1, '用户及角色管理', 'user', 0, 0, 1724573894);
INSERT INTO `node_group` (`id`, `title`, `name`, `sort`, `is_deleted`, `created_at`) VALUES (3, 't222', 'ttt', 0, 0, 1724574570);
INSERT INTO `node_group` (`id`, `title`, `name`, `sort`, `is_deleted`, `created_at`) VALUES (4, 'new', 'nnn', 1, 0, 1724727084);
INSERT INTO `node_group` (`id`, `title`, `name`, `sort`, `is_deleted`, `created_at`) VALUES (5, '系统配置', 'setting', 99, 0, 1725190953);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `role_uuid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色UUID',
  `name` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色说明',
  `authorize_json` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '应用功能授权',
  `is_default` tinyint DEFAULT '0' COMMENT '是否默认角色模版',
  `is_enabled` tinyint unsigned DEFAULT '1' COMMENT '是否启用',
  `is_deleted` tinyint unsigned DEFAULT '0' COMMENT '是否删除',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '添加时间(毫秒)',
  PRIMARY KEY (`role_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色表';

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`role_uuid`, `name`, `description`, `authorize_json`, `is_default`, `is_enabled`, `is_deleted`, `created_at`) VALUES ('b071e6f0-6843-11ef-a2d3-46a1f660a16d', '111', '有一个', '[5]', 0, 1, 0, 1725182540);
INSERT INTO `role` (`role_uuid`, `name`, `description`, `authorize_json`, `is_default`, `is_enabled`, `is_deleted`, `created_at`) VALUES ('eec1f8c6-6831-11ef-a2d3-46a1f660a16d', '超级管理员', 'it\'s test role\n', '[1,2,5,6]', 0, 1, 0, 1725174914);
COMMIT;

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
  `last_login_ip` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登陆IP',
  `last_login_time` int unsigned NOT NULL DEFAULT '0' COMMENT '最后登陆时间',
  `last_active_ip` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后活跃ip',
  `last_active_time` int unsigned NOT NULL DEFAULT '0' COMMENT '最后活跃时间（毫秒）',
  `custom_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '用户数据json',
  `mail_verified` tinyint unsigned DEFAULT '0' COMMENT '邮箱验证',
  `is_deleted` tinyint unsigned DEFAULT '0' COMMENT '是否删除',
  `is_enabled` tinyint unsigned DEFAULT '1' COMMENT '是否启用',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '添加时间(毫秒)',
  PRIMARY KEY (`user_uuid`),
  UNIQUE KEY `user_email` (`email`) USING BTREE,
  KEY `user_uuid` (`user_uuid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`user_uuid`, `username`, `password`, `email`, `name`, `avatar`, `remark`, `last_login_ip`, `last_login_time`, `last_active_ip`, `last_active_time`, `custom_data`, `mail_verified`, `is_deleted`, `is_enabled`, `created_at`) VALUES ('101af57c-6607-11ef-a3df-46a1f660a16e', 'new123', '0ec431f146fdb2a3281092895d436577', '2@qq.com', '12', '', '', '', 0, '', 0, '', 0, 1, 0, 1724936599);
INSERT INTO `user` (`user_uuid`, `username`, `password`, `email`, `name`, `avatar`, `remark`, `last_login_ip`, `last_login_time`, `last_active_ip`, `last_active_time`, `custom_data`, `mail_verified`, `is_deleted`, `is_enabled`, `created_at`) VALUES ('9324eeca-4ca6-11ef-b8ab-46a1f660a16e', 'admin', '376095d3bc790b5661434fd7aa603021', '55811858@qq.com', 'Admin', '', '', '192.168.0.100', 1725191963, '192.168.0.100', 1725191963, '', 0, 0, 1, 1722146378);
INSERT INTO `user` (`user_uuid`, `username`, `password`, `email`, `name`, `avatar`, `remark`, `last_login_ip`, `last_login_time`, `last_active_ip`, `last_active_time`, `custom_data`, `mail_verified`, `is_deleted`, `is_enabled`, `created_at`) VALUES ('c44fc128-660f-11ef-b835-46a1f660a16e', 'WWWWWWWWWWW', 'f0e9d83998ee50ccac75eda1692e684a', 'w12@ww.com', 'Mr12W', '', 'Q12qwer1234', '192.168.0.100', 1724945250, '192.168.0.100', 1724945250, '', 0, 0, 1, 1724940337);
INSERT INTO `user` (`user_uuid`, `username`, `password`, `email`, `name`, `avatar`, `remark`, `last_login_ip`, `last_login_time`, `last_active_ip`, `last_active_time`, `custom_data`, `mail_verified`, `is_deleted`, `is_enabled`, `created_at`) VALUES ('e9b7e4c6-6610-11ef-a866-46a1f660a16e', '12geo', '04a7cf3d12880756113ebc7642e458dc', 'o12@oo.com', '12o', '', 'O12qwer1234', '192.168.0.100', 1725185956, '192.168.0.100', 1725185956, '', 0, 0, 1, 1724940829);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户角色关系表';

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO `user_role` (`id`, `user_uuid`, `role_uuid`, `created_at`) VALUES (5, 'e9b7e4c6-6610-11ef-a866-46a1f660a16e', 'eec1f8c6-6831-11ef-a2d3-46a1f660a16d', 1725182545);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='token表';

-- ----------------------------
-- Records of user_token
-- ----------------------------
BEGIN;
INSERT INTO `user_token` (`id`, `user_uuid`, `refresh_token`, `expires_time`, `created_at`) VALUES (1, '9324eeca-4ca6-11ef-b8ab-46a1f660a16e', 'bcf98605c4b970d5eb830f615eeeb215', 1725278363, 1722146940);
INSERT INTO `user_token` (`id`, `user_uuid`, `refresh_token`, `expires_time`, `created_at`) VALUES (2, 'c44fc128-660f-11ef-b835-46a1f660a16e', '5098d1b518c409c2318e10a6d156cc6b', 1725031650, 1724945250);
INSERT INTO `user_token` (`id`, `user_uuid`, `refresh_token`, `expires_time`, `created_at`) VALUES (3, 'e9b7e4c6-6610-11ef-a866-46a1f660a16e', 'e8f2cbd59e6a7954ebfc3f4c7fb56423', 1725272356, 1725182760);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
