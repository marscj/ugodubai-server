/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : gfast-v32-github

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 19/01/2023 11:13:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', '1', '27', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '28', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '29', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '30', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '2', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '3', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '4', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '11', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '10', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '3', '12', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '3', '13', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '3', '14', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '15', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '19', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '20', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '21', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '24', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '26', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_3', '1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_3', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '31', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_31', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '35', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '33', 'All', '', '', '');

-- ----------------------------
-- Table structure for sys_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_auth_rule`;
CREATE TABLE `sys_auth_rule`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `icon` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '条件',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `menu_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型 0目录 1菜单 2按钮',
  `weigh` int(10) NOT NULL DEFAULT 0 COMMENT '权重',
  `is_hide` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '显示状态',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `is_link` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否外链 1是 0否',
  `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属模块',
  `model_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模型ID',
  `is_iframe` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否内嵌iframe',
  `is_cached` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否缓存',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由重定向地址',
  `is_affix` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否固定',
  `link_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '链接地址',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `weigh`(`weigh`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 59 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单节点表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_auth_rule
-- ----------------------------
INSERT INTO `sys_auth_rule` VALUES (1, 0, 'api/v1/system/auth', '权限管理', 'ele-Stamp', '', '', 0, 30, 0, '/system/auth', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-03-24 15:03:37', '2022-04-14 16:29:19');
INSERT INTO `sys_auth_rule` VALUES (2, 1, 'api/v1/system/menu/list', '菜单管理', 'ele-Calendar', '', '', 1, 0, 0, '/system/auth/menuList', 'system/menu/index', 0, '', 0, 0, 1, '', 0, '', '2022-03-24 17:24:13', '2022-03-29 10:54:49');
INSERT INTO `sys_auth_rule` VALUES (3, 2, 'api/v1/system/menu/add', '添加菜单', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 16:48:43', '2022-03-29 17:05:19');
INSERT INTO `sys_auth_rule` VALUES (4, 2, 'api/v1/system/menu/update', '修改菜单', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 17:04:25', '2022-03-29 18:11:36');
INSERT INTO `sys_auth_rule` VALUES (10, 1, 'api/v1/system/role/list', '角色管理', 'iconfont icon-juxingkaobei', '', '', 1, 0, 0, '/system/auth/roleList', 'system/role/index', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 18:15:03', '2022-03-30 10:25:34');
INSERT INTO `sys_auth_rule` VALUES (11, 2, 'api/v1/system/menu/delete', '删除菜单', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:49:10', '2022-04-06 14:49:17');
INSERT INTO `sys_auth_rule` VALUES (12, 10, 'api/v1/system/role/add', '添加角色', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:49:46', '2022-04-06 14:49:46');
INSERT INTO `sys_auth_rule` VALUES (13, 10, '/api/v1/system/role/edit', '修改角色', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:50:08', '2022-04-06 14:50:08');
INSERT INTO `sys_auth_rule` VALUES (14, 10, '/api/v1/system/role/delete', '删除角色', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:50:22', '2022-04-06 14:50:22');
-- INSERT INTO `sys_auth_rule` VALUES (15, 1, 'api/v1/system/dept/list', '部门管理', 'iconfont icon-siweidaotu', '', '', 1, 0, 0, '/system/auth/deptList', 'system/dept/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:52:23', '2022-04-07 22:59:20');
-- INSERT INTO `sys_auth_rule` VALUES (16, 17, 'aliyun', '阿里云-iframe', 'iconfont icon-diannao1', '', '', 1, 0, 0, '/demo/outLink/aliyun', 'layout/routerView/iframes', 1, '', 0, 1, 1, '', 0, 'https://www.aliyun.com/daily-act/ecs/activity_selection?spm=5176.8789780.J_3965641470.5.568845b58KHj51', '2022-04-06 17:26:29', '2022-04-07 15:27:17');
-- INSERT INTO `sys_auth_rule` VALUES (17, 0, 'outLink', '外链测试', 'iconfont icon-zhongduancanshu', '', '', 0, 20, 0, '/demo/outLink', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 15:20:51', '2022-04-14 16:29:07');
-- INSERT INTO `sys_auth_rule` VALUES (18, 17, 'tenyun', '腾讯云-外链', 'iconfont icon-shouye_dongtaihui', '', '', 1, 0, 0, '/demo/outLink/tenyun', 'layout/routerView/link', 1, '', 0, 0, 1, '', 0, 'https://cloud.tencent.com/act/new?cps_key=20b1c3842f74986b2894e2c5fcde7ea2&fromSource=gwzcw.3775555.3775555.3775555&utm_id=gwzcw.3775555.3775555.3775555&utm_medium=cpc', '2022-04-07 15:23:52', '2022-04-07 15:27:25');
-- INSERT INTO `sys_auth_rule` VALUES (19, 15, 'api/v1/system/dept/add', '添加部门', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:56:39', '2022-04-07 22:56:39');
-- INSERT INTO `sys_auth_rule` VALUES (20, 15, 'api/v1/system/dept/edit', '修改部门', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:57:00', '2022-04-07 22:57:00');
-- INSERT INTO `sys_auth_rule` VALUES (21, 15, 'api/v1/system/dept/delete', '删除部门', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:57:30', '2022-04-07 22:57:30');
-- INSERT INTO `sys_auth_rule` VALUES (22, 1, 'api/v1/system/post/list', '岗位管理', 'iconfont icon-neiqianshujuchucun', '', '', 1, 0, 0, '/system/auth/postList', 'system/post/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:58:46', '2022-04-09 14:26:15');
-- INSERT INTO `sys_auth_rule` VALUES (23, 22, 'api/v1/system/post/add', '添加岗位', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:14:49', '2022-04-09 14:14:49');
-- INSERT INTO `sys_auth_rule` VALUES (24, 22, 'api/v1/system/post/edit', '修改岗位', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:15:25', '2022-04-09 14:15:25');
-- INSERT INTO `sys_auth_rule` VALUES (25, 22, 'api/v1/system/post/delete', '删除岗位', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:15:47', '2022-04-09 14:15:47');
INSERT INTO `sys_auth_rule` VALUES (26, 1, 'api/v1/system/user/list', '用户管理', 'ele-User', '', '', 1, 0, 0, '/system/auth/user/list', 'system/user/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:19:10', '2022-04-09 14:19:58');
INSERT INTO `sys_auth_rule` VALUES (27, 0, 'api/v1/system/dict', '系统配置', 'iconfont icon-shuxingtu', '', '', 0, 40, 0, '/system/dict', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-04-14 16:28:51', '2022-04-18 14:40:56');
INSERT INTO `sys_auth_rule` VALUES (28, 27, 'api/v1/system/dict/type/list', '字典管理', 'iconfont icon-crew_feature', '', '', 1, 0, 0, '/system/dict/type/list', 'system/dict/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-14 16:32:10', '2022-04-16 17:02:50');
INSERT INTO `sys_auth_rule` VALUES (29, 27, 'api/v1/system/dict/dataList', '字典数据管理', 'iconfont icon-putong', '', '', 1, 0, 1, '/system/dict/data/list/:dictType', 'system/dict/dataList', 0, '', 0, 0, 1, '', 0, '', '2022-04-18 12:04:17', '2022-04-18 14:58:43');
INSERT INTO `sys_auth_rule` VALUES (30, 27, 'api/v1/system/config/list', '参数管理', 'ele-Cherry', '', '', 1, 0, 0, '/system/config/list', 'system/config/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-18 21:05:20', '2022-04-18 21:13:19');
INSERT INTO `sys_auth_rule` VALUES (31, 0, 'api/v1/system/monitor', '系统监控', 'iconfont icon-xuanzeqi', '', '', 0, 30, 0, '/system/monitor', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-04-19 10:40:19', '2022-04-19 10:44:38');
INSERT INTO `sys_auth_rule` VALUES (32, 31, 'api/v1/system/monitor/server', '服务监控', 'iconfont icon-shuju', '', '', 1, 0, 0, '/system/monitor/server', 'system/monitor/server/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-19 10:43:32', '2022-04-19 10:44:47');
INSERT INTO `sys_auth_rule` VALUES (33, 35, 'api/swagger', 'api文档', 'iconfont icon--chaifenlie', '', '', 1, 0, 0, '/system/swagger', 'layout/routerView/iframes', 1, '', 0, 1, 1, '', 0, 'http://localhost:8808/swagger', '2022-04-21 09:23:43', '2022-11-29 17:10:35');
INSERT INTO `sys_auth_rule` VALUES (34, 31, 'api/v1/system/loginLog/list', '登录日志', 'ele-Finished', '', '', 1, 0, 0, '/system/monitor/loginLog', 'system/monitor/loginLog/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-28 09:59:47', '2022-04-28 09:59:47');
INSERT INTO `sys_auth_rule` VALUES (35, 0, 'api/v1/system/tools', '系统工具', 'iconfont icon-zujian', '', '', 0, 25, 0, '/system/tools', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-10-26 09:29:08', '2022-10-26 10:11:25');
INSERT INTO `sys_auth_rule` VALUES (38, 31, 'api/v1/system/operLog/list', '操作日志', 'iconfont icon-bolangnengshiyanchang', '', '', 1, 0, 0, '/system/monitor/operLog', 'system/monitor/operLog/index', 0, '', 0, 0, 1, '', 0, '', '2022-12-23 16:19:05', '2022-12-23 16:21:50');
INSERT INTO `sys_auth_rule` VALUES (39, 31, 'api/v1/system/online/list', '在线用户', 'iconfont icon-skin', '', '', 1, 0, 0, '/system/monitor/userOnlie', 'system/monitor/userOnline/index', 0, '', 0, 0, 1, '', 0, '', '2023-01-11 15:48:06', '2023-01-11 17:02:39');

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `config_id` int(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键值',
  `config_type` tinyint(1) NULL DEFAULT 0 COMMENT '系统内置（Y是 N否）',
  `create_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`config_id`) USING BTREE,
  UNIQUE INDEX `uni_config_key`(`config_key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = COMPACT;
  
-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '文件上传-文件大小', 'sys.uploadFile.fileSize', '50M', 1, 31, 31, '文件上传大小限制', NULL, '2021-07-06 14:57:35');
INSERT INTO `sys_config` VALUES (2, '文件上传-文件类型', 'sys.uploadFile.fileType', 'doc,docx,zip,xls,xlsx,rar,jpg,jpeg,gif,npm,png,mp4', 1, 31, 31, '文件上传后缀类型限制', NULL, '2022-12-16 09:52:45');
INSERT INTO `sys_config` VALUES (3, '图片上传-图片类型', 'sys.uploadFile.imageType', 'jpg,jpeg,gif,npm,png', 1, 31, 0, '图片上传后缀类型限制', NULL, NULL);
INSERT INTO `sys_config` VALUES (4, '图片上传-图片大小', 'sys.uploadFile.imageSize', '50M', 1, 31, 31, '图片上传大小限制', NULL, NULL);
INSERT INTO `sys_config` VALUES (11, '静态资源', 'static.resource', '/', 1, 2, 0, '', NULL, NULL);
 
-- ---------------------------- 
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int(4) NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '部门状态（0正常 1停用）',
  `created_by` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int(4) NULL DEFAULT 0 COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表格回显样式',
  `is_default` tinyint(1) NULL DEFAULT 0 COMMENT '是否默认（1是 0否）',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `create_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 0, '男', '1', 'sys_user_sex', '', '', 0, 1, 2, 2, '备注信息', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (2, 0, '女', '2', 'sys_user_sex', '', '', 0, 1, 2, 2, '备注信息', NULL, NULL);

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `dict_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `status` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `create_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, '用户性别', 'sys_user_sex', 1, 2, 1, '用于选择用户性别', NULL, NULL);


-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `login_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统访问记录' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
INSERT INTO `sys_login_log` VALUES (1, 'demo', '::1', '内网IP', 'Chrome', 'Windows 10', 1, '登录成功', '2023-01-19 10:17:18', '系统后台');

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '模块标题',
  `business_type` int(2) NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求方式',
  `operator_type` int(1) NULL DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作人员',
  -- `dept_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作地点',
  `oper_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求参数',
  `error_msg` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime NULL DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志记录' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------


-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int(4) NOT NULL COMMENT '显示顺序',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '岗位信息表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_post
-- ----------------------------


-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态;0:禁用;1:正常',
  `list_order` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint(3) UNSIGNED NOT NULL DEFAULT 3 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, 1, 0, '超级管理员', '备注', 3, '2022-04-01 11:38:39', '2022-04-28 10:00:15');
INSERT INTO `sys_role` VALUES (2, 1, 0, '普通管理员', '备注', 3, '2022-04-01 11:38:39', '2022-04-28 10:01:34');
INSERT INTO `sys_role` VALUES (3, 1, 0, '站点管理员', '站点管理人员', 3, '2022-04-01 11:38:39', '2022-04-01 11:38:39');

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和部门关联表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `birthday` int(11) NOT NULL DEFAULT 0 COMMENT '生日',
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `user_salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密盐',
  `user_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `user_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `sex` tinyint(2) NOT NULL DEFAULT 0 COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `dept_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_admin` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否后台管理员 1 是  0   否',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系地址',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT ' 描述信息',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `agent_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '代理商ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_login`(`user_name`, `deleted_at`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile`, `deleted_at`) USING BTREE,
  INDEX `user_nickname`(`user_nickname`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '13578342363', '超级管理员', 0, 'c567ae329f9929b518759d3bea13f492', 'f9aZTAa8yz', 1, 'yxh669@qq.com', 1, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-19/ccwpeuqz1i2s769hua.jpeg', 101, '', 1, 'asdasfdsaf大发放打发士大夫发按时', '描述信息', '::1', '2022-10-26 03:01:52', '2021-06-22 17:58:00', '2022-11-03 15:44:38', NULL, 1);
INSERT INTO `sys_user` VALUES (2, 'yixiaohu', '13699885599', '奈斯', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'yxh@qq.com', 1, 'upload_file/2022-11-04/co3e5ljknns8jhlp8s.jpg', 102, '备注', 1, '', '', '::1', '2022-11-04 09:54:56', '2021-06-22 17:58:00', '2022-11-04 17:54:56', NULL, 1);
INSERT INTO `sys_user` VALUES (3, 'zs', '16399669855', '张三', 0, '41e3778c20338f4d7d6cc886fd3b2a52', 'redoHIj524', 1, 'zs@qq.com', 0, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-08-02/cd8nif79egjg9kbkgk.jpeg', 101, '', 1, '', '', '::1', '2022-04-28 10:01:47', '2021-06-22 17:58:00', '2022-04-28 10:01:47', NULL, 1);


-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户标识',
  `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
  `create_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `ip` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `explorer` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_token`(`token`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户在线状态表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
INSERT INTO `sys_user_online` VALUES (15, 'f2d1992a5bff07f46a70490451a225e8', '7ZUSfVIf2HyYjcv86SKPPs29v003ECPEScsdYsYYqO1xEIcOpHEu9FS4ZmjQsf1GCmQAky2EuUzyGJF53YyQWvdOP3vC5KeHSmJ1BX0mSAnnw7CD4fNQF4wbtkE4I78lTUjvovXRSC5oDkWPMe79iQ==', '2023-01-13 14:09:51', 'demo', '::1', 'Chrome', 'Windows 10');
INSERT INTO `sys_user_online` VALUES (16, 'c0ce4001700ef589195c41ef073daa62', '7ZUSfVIf2HyYjcv86SKPPs29v003ECPEScsdYsYYqO0y3Gdni2HPIbjTYvAE1/8jYVxUh0VVfhtbUzIENCClH8vlzKtsEfway1I2p8fkF9NRP0ycB7htjT0UJLDmhMUpMaTXSYnL2PPorrqaf4roHg==', '2023-01-19 10:17:18', 'demo', '::1', 'Chrome', 'Windows 10');

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Table structure for sys_agent
-- ----------------------------
-- 删除现有的sys_agent表
DROP TABLE IF EXISTS `sys_agent`;
CREATE TABLE `sys_agent` (
    `agent_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '代理商ID',
    `name` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
    `email` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
    `contact_name` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '联系人姓名',
    `contact_phone` VARCHAR(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '联系人电话',
    `address` VARCHAR(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地址',
    `nationality` VARCHAR(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '国籍',
    `agent_code` VARCHAR(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '代理商代码',
    `available_limit` DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '可用额度',
    `credit_limit` DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '信用额度',
    `outstanding_balance` DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '未结算额度',
    `account_blance` DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '账户余额',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 0.不可用 1.可用',
    `admin_id` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '管理员ID',
    `license_url` VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '许可证URL',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`agent_id`) USING BTREE,
    UNIQUE INDEX `name`(`name`) USING BTREE,
    UNIQUE INDEX `agent_code`(`agent_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代理商表' ROW_FORMAT = COMPACT;

INSERT INTO `sys_agent` VALUES (1, '第五季', 'agent1@example.com', 'Contact Name 1', '1234567890', 'Address 1', 'China', 'AGT001', 10000.14, 20000.14, 5000.14, 100.10, 1, 1, 'https://example.com/license1','2022-11-03 15:44:38','2022-11-03 15:44:38');
INSERT INTO `sys_agent` VALUES (2, '京鱼', 'agent2@example.com', 'Contact Name 2', '2345678901', 'Address 2', 'Japan', 'AGT002', 15000.14, 25000.14, 7500.14, 100.10, 0, 2, 'https://example.com/license2','2022-11-03 15:44:38','2022-11-03 15:44:38');
INSERT INTO `sys_agent` VALUES (3, '程品', 'agent3@example.com', 'Contact Name 3', '3456789012', 'Address 3', 'USA', 'AGT003', 20000.14, 30000.14, 10000.14, 100.10, 1, 3, 'https://example.com/license3','2022-11-03 15:44:38','2022-11-03 15:44:38');
INSERT INTO `sys_agent` VALUES (4, 'Sun Light', 'agent4@example.com', 'Contact Name 4', '4567890123', 'Address 4', 'Germany', 'AGT004', 25000.14, 35000.14, 100.10, 12500.14, 1, 4, 'https://example.com/license4','2022-11-03 15:44:38','2022-11-03 15:44:38');
INSERT INTO `sys_agent` VALUES (5, '任我行', 'agent5@example.com', 'Contact Name 5', '5678901234', 'Address 5', 'France', 'AGT005', 30000.14, 40000.14, 15000.14, 100.10, 1, 5, 'https://example.com/license5','2022-11-03 15:44:38','2022-11-03 15:44:38');
INSERT INTO `sys_agent` VALUES (6, '深圳自由行', 'agent6@example.com', 'Contact Name 6', '6789012345', 'Address 6', 'China', 'AGT006', 35000.14, 45000.14, 17500.14, 100.10, 1, 6, 'https://example.com/license6','2022-11-03 15:44:38','2022-11-03 15:44:38');
INSERT INTO `sys_agent` VALUES (7, '野图', 'agent7@example.com', 'Contact Name 7', '7890123456', 'Address 7', 'Australia', 'AGT007', 40000.14, 50000.14, 20000.14, 100.10, 1, 7, 'https://example.com/license7','2022-11-03 15:44:38','2022-11-03 15:44:38');
INSERT INTO `sys_agent` VALUES (8, 'Silk Way', 'agent8@example.com', 'Contact Name 8', '8901234567', 'Address 8', 'China', 'AGT008', 45000.14, 55000.14, 22500.14, 100.10, 1, 8, 'https://example.com/license8','2022-11-03 15:44:38','2022-11-03 15:44:38');
INSERT INTO `sys_agent` VALUES (9, '皇家旅途', 'agent9@example.com', 'Contact Name 9', '9012345678', 'Address 9', 'UK', 'AGT009', 50000.14, 60000.14, 25000.14, 100.10, 1, 9, 'https://example.com/license9','2022-11-03 15:44:38','2022-11-03 15:44:38');
INSERT INTO `sys_agent` VALUES (10, 'Starking', 'agent10@example.com', 'Contact Name 10', '0123456789', 'Address 10', 'China', 'AGT010', 55000.14, 65000.14, 27500.14, 100.10, 1, 10, 'https://example.com/license10','2022-11-03 15:44:38','2022-11-03 15:44:38');

-- ----------------------------
-- Table structure for sys_supplier
-- ----------------------------
DROP TABLE IF EXISTS `sys_supplier`;
CREATE TABLE `sys_supplier` (
  `supplier_id` bigint(20) UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '唯一标识符',
  `name` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `code` VARCHAR(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  UNIQUE INDEX `name`(`name`) USING BTREE
)ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '供应商' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Table structure for sys_booking
-- ----------------------------
-- 删除现有的sys_booking表
DROP TABLE IF EXISTS `sys_booking`;
CREATE TABLE `sys_booking` ( 
  `booking_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '代理商ID',
  `parent_id` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '父',
  `related_id` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT  '关联订单ID',
  `agent_id` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '代理商ID',
  `action_date` DATETIME NULL DEFAULT NULL COMMENT '执行日期',
  `product_id` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '产品ID',
  `variation_id` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '变体产品',
  `fit_number` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT  'FIT',
  `sku` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT  'SKU',
  `guest_name` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT  '客人姓名',
  `guest_contact` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '客人联系方式',
  `product_name` VARCHAR(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT  '产品名称',
  `unit_price` DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '单价',
  `quantity` INT NOT NULL DEFAULT 1 COMMENT '数量',
  `total_price` DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '总价',
  `tax` DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '税',
  `currency` VARCHAR(3)  NULL DEFAULT 'AED' COMMENT '货币',
  `booking_status` TINYINT NOT NULL DEFAULT 1 COMMENT '订单状态 0.待核单 1.已核单出票中 2.已出票 3.取消待确认 4.已取消',
  -- `payment_status` TINYINT NOT NULL DEFAULT 1 COMMENT '支付状态 0.等待支付 1.未支付 2.已支付 3.已退款',
  -- `payment_method` TINYINT NOT NULL DEFAULT 1 COMMENT '支付方式 0.余额 1.信用 2.支付宝 3.微信 4.公司转账',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_by` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `updated_by` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`booking_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Table structure for sys_booking_meta
-- ----------------------------
DROP TABLE IF EXISTS `sys_booking_meta`;
CREATE TABLE `sys_booking_meta` (
  `meta_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `booking_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `meta_key` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_value` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`meta_id`),
  KEY `booking_id` (`booking_id`),
  KEY `meta_key` (`meta_key`(32))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '订单meta表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Table structure for sys_gateway
-- ----------------------------
DROP TABLE IF EXISTS `sys_gateway`;
CREATE TABLE `sys_gateway` (
  `gateway_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name_en` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `name_cn` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `content` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`gateway_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '支付表' ROW_FORMAT = COMPACT;
INSERT INTO `sys_gateway` (`name_en`, `name_cn`, `content`) VALUES
('Banlance', '余额', ''),
('Credit', '信用度', ''),
('Bank Transfer', '银行转账', ''),
('Wechat', '微信', ''),
('Alipay', '支付宝', '');

-- ----------------------------
-- Table structure for sys_payment_tokens
-- ----------------------------
DROP TABLE IF EXISTS `sys_payment_tokens`;
CREATE TABLE `sys_payment_tokens` (
  `token_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `gateway_id` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `agent_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `type` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_default` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`token_id`),
  KEY `user_id` (`user_id`),
  KEY `agent_id` (`agent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '支付表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Table structure for sys_payment_tokenmeta
-- ----------------------------
DROP TABLE IF EXISTS `sys_payment_tokenmeta`;
CREATE TABLE `sys_payment_tokenmeta` (
  `meta_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `payment_token_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `meta_key` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_value` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`meta_id`),
  KEY `payment_token_id` (`payment_token_id`),
  KEY `meta_key` (`meta_key`(32))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '支付meta表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Table structure for sys_product
-- ----------------------------
DROP TABLE IF EXISTS `sys_product`;
CREATE TABLE `sys_product` (
  `product_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name_en` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `name_cn` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `description_en` longtext COLLATE utf8mb4_unicode_ci  COMMENT '产品简介',
  `description_cn` longtext COLLATE utf8mb4_unicode_ci  COMMENT '产品简介',
  `content_en` longtext COLLATE utf8mb4_unicode_ci COMMENT '产品内容',
  `content_cn` longtext COLLATE utf8mb4_unicode_ci COMMENT '产品内容',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态 0.下线 1.上线',
  `image` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '缩略图',
  `order` int(4) NULL DEFAULT 0 COMMENT '显示顺序',
  `is_deleted` TINYINT NOT NULL DEFAULT 0 COMMENT '状态 0.未删除 1.已删除',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`product_id`),
  KEY `name_en` (`name_en`),
  KEY `name_cn` (`name_cn`),
  KEY `status` (`status`),
  KEY `order` (`order`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品表' ROW_FORMAT = COMPACT;

INSERT INTO `sys_product` (`product_id`, `name_cn`, `image`, `description_cn`, `content_cn`) VALUES
(1, '哈利法塔', 'https://www.burjkhalifa.ae/images/gallery/burj-khalifa_02.jpg', '迪拜塔，也被称为哈利法塔，是阿联酋迪拜市的标志性建筑物。它是世界上最高的人造结构物，高828米，拥有163层。迪拜塔于2010年完工，并成为一个多功能建筑，包括豪华酒店、住宅、观光景点和办公空间。塔内设有观景台，游客可以欣赏到壮观的城市景色。迪拜塔成为迪拜市的象征之一，吸引着全球游客前来参观。', '<div><p>参观迪拜塔时需要注意以下事项：</p>
<ol>
<li>提前预订门票：迪拜塔是一座著名的旅游景点，门票需提前预订以确保入场。</li>
<li>遵守规定：参观者需遵守迪拜塔的参观规定和指示，包括禁止吸烟、禁止携带食物等。</li>
<li>注意穿着：建议穿着轻便舒适的服装，因为在迪拜塔室外环境可能较炎热。</li>
<li>爬塔禁止：迪拜塔的顶部并不对公众开放，不允许任何人爬塔或试图进入无授权区域。</li>
<li>安全意识：参观时请保持警惕，注意个人财物安全，并遵守塔内的安全提示和指示。</li>
<li>摄影规定：摄影是允许的，但需要遵守规定，尤其是在观景台上注意不要阻碍其他游客。</li>
<li>注意时间安排：预计参观迪拜塔可能需要一些时间，包括排队等待和观赏城市景色，建议提前做好时间规划。</li>
<li>注意天气状况：由于迪拜气候炎热，建议参观者带上防晒霜、帽子和水，并在炎热期间随时寻找遮阴处。
请记住这些注意事项以确保您的参观体验顺利和愉快。</li>
</ol>
</div>');

-- ----------------------------
-- Table structure for sys_terms
-- ----------------------------
DROP TABLE IF EXISTS `sys_product_terms`;
CREATE TABLE `sys_product_terms` (
  `term_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `taxonomy` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `parent` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `name_en` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `name_cn` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`term_id`),
  KEY `taxonomy` (`taxonomy`),
  KEY `parent` (`parent`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '产品术语表' ROW_FORMAT = COMPACT;

INSERT INTO `sys_product_terms` (`term_id`, `taxonomy`, `name_cn`, `name_en`) VALUES
(1, 'category', '门票', 'Ticket'),
(2, 'category', '一日游', 'Tours'),
(3, 'category', '餐饮', 'Food'),
(4, 'category', '租车', 'Car Rental'),
(5, 'tag', '特色推荐', 'Featured'),
(6, 'tag', '必玩', 'Must Play');

-- ----------------------------
-- Table structure for sys_product_terms_lookup
-- ----------------------------
DROP TABLE IF EXISTS `sys_product_terms_lookup`;
CREATE TABLE `sys_product_terms_lookup` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `term_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `product_id` (`product_id`),
  KEY `term_id` (`term_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '产品关联表' ROW_FORMAT = COMPACT;

INSERT INTO `sys_product_terms_lookup` (`product_id`, `term_id`) VALUES
( 1, 1),
( 1, 5),
( 1, 6);

-- ----------------------------
-- Table structure for sys_product_price_lookup
-- ----------------------------
DROP TABLE IF EXISTS `sys_product_price_lookup`;
CREATE TABLE `sys_product_price_lookup` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `min_price`  DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '最低价格',
  `max_price`  DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '最高价格',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id`(`id`, `product_id`) USING BTREE,
  KEY `product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '产品价格关联' ROW_FORMAT = COMPACT;
INSERT INTO `sys_product_price_lookup` (`product_id`, `min_price`, `max_price`) VALUES
( 1, 147.00, 255),
( 1, 85.00, 289),
( 1, 85.00, 289);

-- ----------------------------
-- Table structure for sys_product_meta
-- ----------------------------
DROP TABLE IF EXISTS `sys_product_meta`;
CREATE TABLE `sys_product_meta` (
  `meta_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `meta_key` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `meta_value` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`meta_id`),
  KEY `product_id` (`product_id`),
  KEY `meta_key` (`meta_key`(32))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '产品meta表' ROW_FORMAT = COMPACT;
INSERT INTO `sys_product_meta` (`meta_id`, `product_id`, `meta_key`, `meta_value`) VALUES
(1, 1, 'gallery_url', 'google.com'),
(2, 1, 'gallery_url', 'google.com');

-- ----------------------------
-- Table structure for sys_product_keywords
-- ----------------------------
DROP TABLE IF EXISTS `sys_product_keywords`;
CREATE TABLE `sys_product_keywords` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `keyword` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `product_id` (`product_id`),
  KEY `keyword` (`keyword`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='产品关键字表' ROW_FORMAT = COMPACT;
INSERT INTO `sys_product_keywords` (`product_id`, `keyword`) VALUES
(1, '哈利法塔'),
(1, 'Burjkhalifa'),
(1, 'hlft');

-- ----------------------------
-- Table structure for sys_variation
-- ----------------------------
DROP TABLE IF EXISTS `sys_variation`;
CREATE TABLE `sys_variation` (
  `variation_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `name_en` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '默认名称',
  `name_cn` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '默认名称',
  `sku` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT  'SKU',
  `status` TINYINT NOT NULL DEFAULT '0' COMMENT '状态 0.下线 1.上线',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`variation_id`),
  KEY `name_en` (`name_en`),
  KEY `name_cn` (`name_cn`),
  KEY `sku` (`sku`),
  KEY `status` (`status`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品变体表' ROW_FORMAT = COMPACT;
INSERT INTO `sys_variation` (`variation_id`, `product_id`, `sku`, `name_cn`, `name_en`) VALUES
(1, 1, 'ATT-124+125-NonPrime', '124+125层普通票', '124+125 Non Prime'),
(2, 1, 'ATT-124+125-Prime', '124+125层黄金票', '124+125 Prime');

-- ----------------------------
-- Table structure for sys_variation_attribute
-- ----------------------------
DROP TABLE IF EXISTS `sys_variation_attribute`;
CREATE TABLE `sys_variation_attribute` (
  `attribute_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `attribute_key` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `attribute_value_en` longtext COLLATE utf8mb4_unicode_ci,
  `attribute_value_cn` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`attribute_id`),
  KEY `attribute_key` (`attribute_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '产品变体参数表' ROW_FORMAT = COMPACT;
INSERT INTO `sys_variation_attribute` (`attribute_id`, `attribute_key`, `attribute_value_cn`, `attribute_value_en`) VALUES
(1, 'age', '成人', 'Aldult'),
(2, 'age', '小孩', 'Child');

-- ----------------------------
-- Table structure for sys_variation_price
-- ----------------------------
DROP TABLE IF EXISTS `sys_variation_price`;
CREATE TABLE `sys_variation_price` (
  `variation_price_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `variation_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `attribute_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `agent_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `start_date` DATE DEFAULT NULL,
  `end_date` DATE DEFAULT NULL,
  `cost_price` DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '成本价',
  `special_price`  DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '预付价格',
  `selling_price`  DECIMAL(10, 2) NOT NULL DEFAULT 0.0 COMMENT '销售价',
  `currency` VARCHAR(3) NOT NULL DEFAULT 'AED' COMMENT '货币',
  `stock` int(11) NOT NULL DEFAULT '0',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`variation_price_id`),
  KEY `end_date` (`end_date`),
  KEY `start_date` (`start_date`),
  KEY `special_price` (`special_price`),
  KEY `selling_price` (`selling_price`),
  KEY `currency` (`currency`),
  KEY `stock` (`stock`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品变体价格表' ROW_FORMAT = COMPACT;
INSERT INTO `sys_variation_price` (`variation_id`, `attribute_id`, `agent_id`, `start_date`, `end_date`, `cost_price`,`special_price`,`selling_price`) VALUES
(1, 1, 1, '2023-07-24', '2023-07-30', '147.00', '150.00', '155.00'),
(1, 2, 1, '2023-07-24', '2023-07-30', '120.00', '124.00', '147.00'),
(2, 1, 1, '2023-07-24', '2023-07-30', '247.00', '250.00', '255.00'),
(2, 2, 1,'2023-07-24', '2023-07-30', '220.00', '224.00', '247.00'),
(1, 1, 2, '2023-07-24', '2023-07-30', '147.00', '150.00', '155.00'),
(1, 2, 2, '2023-07-24', '2023-07-30', '120.00', '124.00', '147.00'),
(2, 1, 2, '2023-07-24', '2023-07-30', '247.00', '250.00', '255.00'),
(2, 2, 2,'2023-07-24', '2023-07-30', '220.00', '224.00', '247.00'),
(1, 1, 0, '2023-07-24', '2023-07-30', '947.00', '950.00', '955.00'),
(1, 2, 0, '2023-07-24', '2023-07-30', '920.00', '924.00', '947.00'),
(2, 1, 0, '2023-07-24', '2023-07-30', '947.00', '950.00', '955.00'),
(2, 2, 0,'2023-07-24', '2023-07-30', '920.00', '924.00', '947.00');

-- ----------------------------
-- Table structure for sys_product_lookup
-- ----------------------------
DROP TABLE IF EXISTS `sys_product_lookup`;
CREATE TABLE `sys_product_lookup` (
  `variation_lookup_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `variation_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `attribute_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `variation_price_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `agent_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`variation_lookup_id`),
  KEY `product_id` (`product_id`),
  KEY `variation_id` (`variation_id`),
  KEY `attribute_id` (`attribute_id`),
  KEY `variation_price_id` (`variation_price_id`),
  KEY `agent_id` (`agent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '产品变体关联表' ROW_FORMAT = COMPACT;
INSERT INTO `sys_product_lookup` (`product_id`, `variation_id`, `attribute_id`,`variation_price_id`, `agent_id`) VALUES
( 1, 1, 1, 1, 1),
( 1, 1, 2, 2, 1),
( 1, 2, 1, 3, 1),
( 1, 2, 2, 4, 0);

SET FOREIGN_KEY_CHECKS = 1;