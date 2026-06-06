/*
 Navicat Premium Dump SQL

 Source Server         : wsl-ubuntu-mysql8
 Source Server Type    : MySQL
 Source Server Version : 80409 (8.4.9)
 Source Host           : 127.0.0.1:3306
 Source Schema         : gin-admin

 Target Server Type    : MySQL
 Target Server Version : 80409 (8.4.9)
 File Encoding         : 65001

 Date: 06/06/2026 21:09:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_logs
-- ----------------------------
DROP TABLE IF EXISTS `admin_logs`;
CREATE TABLE `admin_logs`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `ip` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `input` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_admin_operation_logs_path`(`path` ASC) USING BTREE,
  INDEX `idx_admin_operation_logs_method`(`method` ASC) USING BTREE,
  INDEX `idx_admin_operation_logs_ip`(`ip` ASC) USING BTREE,
  INDEX `idx_admin_operation_logs_created_at`(`created_at` ASC) USING BTREE,
  INDEX `idx_admin_operation_logs_updated_at`(`updated_at` ASC) USING BTREE,
  INDEX `fk_admin_operation_logs_admin_user`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 52 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_logs
-- ----------------------------
INSERT INTO `admin_logs` VALUES (50, 1, '/admin/config/:id', '/admin/config/1', 'PUT', '127.0.0.1', '{\"config_value\":\"管理后台\",\"config_key\":\"name\",\"config_label\":\"网站名称\",\"type\":0,\"is_can_front\":0,\"is_required\":0,\"order\":1,\"group_id\":0,\"state\":1}', '2026-06-06 21:09:06.654', '2026-06-06 21:09:06.654');
INSERT INTO `admin_logs` VALUES (51, 1, '/admin/config/:id', '/admin/config/1', 'PUT', '127.0.0.1', '{\"config_value\":\"管理后台\",\"config_key\":\"name\",\"config_label\":\"网站名称\",\"type\":0,\"describe\":\"管理后台配置式例\",\"is_can_front\":0,\"is_required\":0,\"order\":1,\"group_id\":0,\"state\":1}', '2026-06-06 21:09:37.300', '2026-06-06 21:09:37.300');

-- ----------------------------
-- Table structure for admin_menus
-- ----------------------------
DROP TABLE IF EXISTS `admin_menus`;
CREATE TABLE `admin_menus`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` bigint UNSIGNED NOT NULL DEFAULT 0,
  `order` bigint UNSIGNED NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_show` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否展示在菜单中',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_admin_menus_created_at`(`created_at` ASC) USING BTREE,
  INDEX `idx_admin_menus_updated_at`(`updated_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_menus
-- ----------------------------
INSERT INTO `admin_menus` VALUES (11, 0, 1, '管理后台', 'lucide:settings', 'system', '#', 1, '2026-04-27 12:26:35.285', '2026-04-27 12:26:35.285');
INSERT INTO `admin_menus` VALUES (12, 11, 2, '管理员', 'lucide:user-round-pen', 'admin/users/index', 'admin/users/index', 1, '2026-04-27 12:26:35.285', '2026-04-27 12:26:35.285');
INSERT INTO `admin_menus` VALUES (13, 11, 3, '角色', 'lucide:shield-check', 'admin/roles/index', 'admin/roles/index', 1, '2026-04-27 12:26:35.285', '2026-04-27 12:26:35.285');
INSERT INTO `admin_menus` VALUES (14, 11, 4, '权限', 'lucide:lock-keyhole', 'admin/permissions/index', 'admin/permissions/index', 1, '2026-04-27 12:26:35.285', '2026-04-27 12:26:35.285');
INSERT INTO `admin_menus` VALUES (15, 11, 5, '菜单', 'lucide:menu', 'admin/menus/index', 'admin/menus/index', 1, '2026-04-27 12:26:35.285', '2026-04-27 12:26:35.285');
INSERT INTO `admin_menus` VALUES (16, 11, 6, '日志', 'lucide:scroll-text', 'admin/logs/index', 'admin/logs/index', 1, '2026-04-27 12:26:35.285', '2026-04-27 12:26:35.285');
INSERT INTO `admin_menus` VALUES (17, 11, 7, '文件', 'lucide:folder-open', 'admin/files/index', 'admin/files/index', 1, '2026-04-27 12:26:35.285', '2026-04-27 12:26:35.285');
INSERT INTO `admin_menus` VALUES (18, 11, 8, '配置', 'lucide:wrench', 'admin/configs/index', 'admin/configs/index', 1, '2026-04-27 12:26:35.285', '2026-04-27 12:26:35.285');

-- ----------------------------
-- Table structure for admin_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_permissions`;
CREATE TABLE `admin_permissions`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `slug` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `http_method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `http_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `order` bigint UNSIGNED NULL DEFAULT NULL,
  `parent_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_admin_permissions_created_at`(`created_at` ASC) USING BTREE,
  INDEX `idx_admin_permissions_updated_at`(`updated_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 76 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_permissions
-- ----------------------------
INSERT INTO `admin_permissions` VALUES (10, '管理员管理', 'admin:users', 'any', '/admin/users/index', 10, 0, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (11, '管理员列表', 'admin:users:list', 'get', '/admin/users', 11, 10, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (12, '管理员创建', 'admin:users:create', 'post', '/admin/user', 12, 10, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (13, '管理员修改', 'admin:users:update', 'put', 'admin/user/:id', 13, 10, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (14, '管理员删除', 'admin:users:delete', 'delete', 'admin/user/:id', 14, 10, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (15, '单个管理员信息', 'admin:users:get', 'get', 'admin/user/:id', 15, 10, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (20, '角色管理', 'admin:roles', 'any', '/admin/roles/index', 20, 0, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (21, '角色列表', 'admin:roles:list', 'get', '/admin/roles', 21, 20, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (22, '角色创建', 'admin:roles:ctreate', 'post', '/admin/role', 22, 20, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (23, '角色修改', 'admin:role:update', 'put', 'admin/role/:id', 23, 20, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (24, '角色删除', 'admin:role:delete', 'delete', 'admin/role/:id', 24, 20, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (25, '单个角色信息', 'admin:roles:get', 'get', 'admin/role/:id', 25, 20, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (26, '角色添加菜单', 'admin:roles:menus', 'post', 'admin/role/:id/menus', 26, 20, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (27, '角色添加权限', 'admin:roles:permissions', 'post', 'admin/role/:id/permissions', 27, 20, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (30, '菜单管理', 'admin:menus', 'any', '/admin/menus/index', 30, 0, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (31, '菜单列表', 'admin:menus:list', 'get', '/admin/menus', 31, 30, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (32, '菜单创建', 'admin:menus:create', 'post', '/admin/menu', 32, 30, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (33, '菜单修改', 'admin:menus:update', 'put', '/admin/menu/:id', 33, 30, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (34, '菜单删除', 'admin:menus:delete', 'delete', '/admin/menu/:id', 34, 30, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (35, '单个菜单信息', 'admin:menus:get', 'get', '/admin/menu/:id', 35, 30, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (40, '权限管理', 'admin:perimissions', 'any', '/admin/perimissions/index', 40, 0, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (41, '权限列表', 'admin:perimissions:list', 'get', '/admin/perimissions', 41, 40, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (42, '权限创建', 'admin:perimissions:create', 'post', '/admin/perimission', 42, 40, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (43, '权限修改', 'admin:perimissions:update', 'put', '/admin/perimission/:id', 43, 40, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (44, '权限删除', 'admin:perimissions:delete', 'delete', '/admin/perimission/:id', 44, 40, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (45, '单个权限信息', 'admin:perimissions:get', 'get', '/admin/perimission/:id', 45, 40, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (50, '配置管理', 'admin:configs', 'any', '/admin/configs/index', 50, 0, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (51, '配置列表', 'admin:configs:list', 'get', '/admin/configs', 51, 50, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (52, '配置创建', 'admin:configs:create', 'post', '/admin/config', 52, 50, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (53, '配置修改', 'admin:configs:update', 'put', '/admin/config/:id', 53, 50, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (54, '配置删除', 'admin:configs:delete', 'delete', '/admin/config/:id', 54, 50, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (55, '单个配置信息', 'admin:configs:get', 'get', '/admin/config/:id', 55, 50, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (60, '文件管理', 'admin:files', 'any', '/admin/files/index', 60, 0, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (61, '文件列表', 'admin:files:list', 'get', '/admin/files', 61, 60, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (62, '文件创建', 'admin:files:cteate', 'post', '/admin/file', 62, 60, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (63, '文件修改', 'admin:files:update', 'put', '/admin/file/:id', 63, 60, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (64, '文件删除', 'admin:files:delete', 'delete', '/admin/file/:id', 64, 60, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (65, '文件单个信息', 'admin:files:get', 'get', '/admin/file/:id', 65, 60, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (66, '文件上传', 'admin:files:update', 'post', '/admin/upload', 66, 60, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (70, '日志管理', 'admin:logs', 'any', '/admin/logs/index', 70, 0, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (71, '日志列表', 'admin:logs:list', 'get', '/admin/logs', 71, 70, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_permissions` VALUES (72, '日志单个信息', 'admin:logs:get', 'get', '/admin/log/:id', 72, 70, '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');

-- ----------------------------
-- Table structure for admin_role_menus
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_menus`;
CREATE TABLE `admin_role_menus`  (
  `role_id` bigint UNSIGNED NOT NULL,
  `menu_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE,
  INDEX `fk_admin_role_menus_admin_menu`(`menu_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_role_menus
-- ----------------------------

-- ----------------------------
-- Table structure for admin_role_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_permissions`;
CREATE TABLE `admin_role_permissions`  (
  `role_id` bigint UNSIGNED NOT NULL,
  `permission_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`role_id`, `permission_id`) USING BTREE,
  INDEX `fk_admin_role_permissions_admin_permission`(`permission_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_role_permissions
-- ----------------------------
INSERT INTO `admin_role_permissions` VALUES (2, 20);
INSERT INTO `admin_role_permissions` VALUES (2, 21);
INSERT INTO `admin_role_permissions` VALUES (2, 22);
INSERT INTO `admin_role_permissions` VALUES (2, 23);
INSERT INTO `admin_role_permissions` VALUES (2, 24);
INSERT INTO `admin_role_permissions` VALUES (2, 25);

-- ----------------------------
-- Table structure for admin_role_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_users`;
CREATE TABLE `admin_role_users`  (
  `user_id` bigint UNSIGNED NOT NULL,
  `role_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE,
  INDEX `fk_admin_role_users_admin_role`(`role_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_role_users
-- ----------------------------
INSERT INTO `admin_role_users` VALUES (1, 1);

-- ----------------------------
-- Table structure for admin_roles
-- ----------------------------
DROP TABLE IF EXISTS `admin_roles`;
CREATE TABLE `admin_roles`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `slug` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_admin_roles_created_at`(`created_at` ASC) USING BTREE,
  INDEX `idx_admin_roles_updated_at`(`updated_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_roles
-- ----------------------------
INSERT INTO `admin_roles` VALUES (1, '超级管理员', 'admin', '2026-06-04 21:57:11.360', '2026-06-04 21:57:11.360');
INSERT INTO `admin_roles` VALUES (2, '测试', 'test', '2026-06-04 21:57:11.360', '2026-06-06 21:04:54.388');

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `avatar_id` bigint UNSIGNED NULL DEFAULT NULL,
  `is_super` tinyint UNSIGNED NOT NULL DEFAULT 0,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `u_username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `u_email`(`email` ASC) USING BTREE,
  INDEX `idx_admin_users_created_at`(`created_at` ASC) USING BTREE,
  INDEX `idx_admin_users_updated_at`(`updated_at` ASC) USING BTREE,
  INDEX `fk_admin_users_avatar_file`(`avatar_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_users
-- ----------------------------
INSERT INTO `admin_users` VALUES (1, 'admin', '$2a$14$UPDOeuhOq6k6o2jnp3rCnudpcogjfSImV9hsHjKSEuMsPdoWY9Pk6', NULL, 'Administrator', NULL, 1, '2026-04-27 14:34:45.205', '2026-04-27 14:34:45.205', NULL);

-- ----------------------------
-- Table structure for configs
-- ----------------------------
DROP TABLE IF EXISTS `configs`;
CREATE TABLE `configs`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `config_value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `config_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `type` tinyint NULL DEFAULT NULL,
  `options` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `describe` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `is_can_front` bigint NULL DEFAULT NULL,
  `is_required` bigint UNSIGNED NULL DEFAULT NULL,
  `order` bigint UNSIGNED NULL DEFAULT NULL,
  `group_id` bigint UNSIGNED NULL DEFAULT NULL,
  `state` tinyint NULL DEFAULT NULL,
  `show_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `placeholder` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `u_configs_config_key`(`config_key` ASC) USING BTREE,
  INDEX `idx_configs_config_label`(`config_label` ASC) USING BTREE,
  INDEX `idx_configs_created_at`(`created_at` ASC) USING BTREE,
  INDEX `idx_configs_updated_at`(`updated_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of configs
-- ----------------------------
INSERT INTO `configs` VALUES (1, 'name', '管理后台', '网站名称', 0, '', '管理后台配置式例', 0, 0, 1, 0, 1, '', '', '2026-06-04 21:38:38.252', '2026-06-06 21:09:37.295');

-- ----------------------------
-- Table structure for demo_table
-- ----------------------------
DROP TABLE IF EXISTS `demo_table`;
CREATE TABLE `demo_table`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` datetime(3) NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `i_created_at`(`created_at` ASC) USING BTREE,
  INDEX `i_updated_at`(`updated_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of demo_table
-- ----------------------------

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `origin_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `group_id` bigint NULL DEFAULT NULL,
  `size` bigint NULL DEFAULT NULL,
  `storage` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `type` tinyint NULL DEFAULT NULL,
  `ext` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `content_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `e_tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `bucket` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `last_modified` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_files_group_id`(`group_id` ASC) USING BTREE,
  INDEX `idx_files_storage`(`storage` ASC) USING BTREE,
  INDEX `idx_files_created_at`(`created_at` ASC) USING BTREE,
  INDEX `idx_files_updated_at`(`updated_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of files
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
