# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 8.0.18)
# Database: ncov2
# Generation Time: 2020-03-22 06:22:59 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table goadmin_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_menu`;

CREATE TABLE `goadmin_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0',
  `type` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `order` int(11) unsigned NOT NULL DEFAULT '0',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `header` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_menu` DISABLE KEYS */;

INSERT INTO `goadmin_menu` (`id`, `parent_id`, `type`, `order`, `title`, `icon`, `uri`, `header`, `created_at`, `updated_at`)
VALUES
	(1,0,1,2,'Admin','fa-tasks','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,1,1,2,'Users','fa-users','/info/manager',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(3,1,1,3,'Roles','fa-user','/info/roles',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(4,1,1,4,'Permission','fa-ban','/info/permission',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(5,1,1,5,'Menu','fa-bars','/menu',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(6,1,1,6,'Operation log','fa-history','/info/op',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(7,0,1,1,'Dashboard','fa-bar-chart','/',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(8,0,0,2,'白名单','fa-bars','/info/org_whitelist','','2020-03-20 22:17:50','2020-03-20 22:17:50'),
	(9,0,0,2,'今日未提交名单','fa-bars','/info/org_whitelist','','2020-03-20 22:18:32','2020-03-20 22:20:58'),
	(10,0,0,2,'部门','fa-users','','','2020-03-20 22:20:00','2020-03-20 22:20:53'),
	(11,10,0,2,'部门管理','fa-align-left','/info/org_dep','','2020-03-20 22:20:39','2020-03-20 22:20:46'),
	(12,10,0,2,'部门管理员','fa-users','/info/admin_user','','2020-03-20 22:21:35','2020-03-21 12:01:06'),
	(13,0,0,2,'小程序绑定用户','fa-user','/info/wx_mp_bind_info','','2020-03-20 22:23:03','2020-03-20 22:23:03'),
	(14,0,0,2,'上报数据','fa-database','/info/report_record_default','','2020-03-20 22:35:24','2020-03-20 22:35:24'),
	(15,0,0,2,'机构管理','fa-bank','/info/organization','','2020-03-21 12:18:16','2020-03-21 12:18:40');

/*!40000 ALTER TABLE `goadmin_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_operation_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_operation_log`;

CREATE TABLE `goadmin_operation_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `input` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `admin_operation_log_user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_operation_log` WRITE;
/*!40000 ALTER TABLE `goadmin_operation_log` DISABLE KEYS */;

INSERT INTO `goadmin_operation_log` (`id`, `user_id`, `path`, `method`, `ip`, `input`, `created_at`, `updated_at`)
VALUES
	(1,1,'/admin/info/manager','GET','127.0.0.1','','2020-03-20 22:16:57','2020-03-20 22:16:57'),
	(2,1,'/admin/menu','GET','127.0.0.1','','2020-03-20 22:17:00','2020-03-20 22:17:00'),
	(3,1,'/admin/menu','GET','127.0.0.1','','2020-03-20 22:17:04','2020-03-20 22:17:04'),
	(4,1,'/admin/menu/new','POST','127.0.0.1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"7f37c8f1-f72d-4ba5-9e43-9f47bedd75b2\"],\"header\":[\"\"],\"icon\":[\"fa-bars\"],\"parent_id\":[\"\"],\"roles[]\":[\"2\"],\"title\":[\"白名单\"],\"uri\":[\"/info/org_whitelist\"]}','2020-03-20 22:17:50','2020-03-20 22:17:50'),
	(5,1,'/admin/menu/new','POST','127.0.0.1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"e3b46c3d-82c3-4737-b8bb-74b74bc81389\"],\"header\":[\"\"],\"icon\":[\"fa-bars\"],\"parent_id\":[\"\"],\"roles[]\":[\"2\"],\"title\":[\"今日未提交名单\"],\"uri\":[\"/info/org_whitelist\"]}','2020-03-20 22:18:32','2020-03-20 22:18:32'),
	(6,1,'/admin/menu/new','POST','127.0.0.1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"297efed1-d40f-4ed5-95ee-f9ef5c733ab3\"],\"header\":[\"\"],\"icon\":[\"fa-users\"],\"parent_id\":[\"\"],\"title\":[\"部门\"],\"uri\":[\"\"]}','2020-03-20 22:20:00','2020-03-20 22:20:00'),
	(7,1,'/admin/menu/new','POST','127.0.0.1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"9b995bee-b9f2-42d5-9ab1-22cb5fc43844\"],\"header\":[\"\"],\"icon\":[\"fa-align-left\"],\"parent_id\":[\"10\"],\"title\":[\"部门管理\"],\"uri\":[\"/info/org_dep\"]}','2020-03-20 22:20:39','2020-03-20 22:20:39'),
	(8,1,'/admin/menu/edit/show','GET','127.0.0.1','','2020-03-20 22:20:42','2020-03-20 22:20:42'),
	(9,1,'/admin/menu/edit','POST','127.0.0.1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"7f073712-791a-4096-b673-1195e468d980\"],\"header\":[\"\"],\"icon\":[\"fa-align-left\"],\"id\":[\"11\"],\"parent_id\":[\"10\"],\"roles[]\":[\"2\"],\"title\":[\"部门管理\"],\"uri\":[\"/info/org_dep\"]}','2020-03-20 22:20:46','2020-03-20 22:20:46'),
	(10,1,'/admin/menu/edit/show','GET','127.0.0.1','','2020-03-20 22:20:50','2020-03-20 22:20:50'),
	(11,1,'/admin/menu/edit','POST','127.0.0.1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"30ed2498-21c0-4123-9d20-296e21640661\"],\"header\":[\"\"],\"icon\":[\"fa-users\"],\"id\":[\"10\"],\"parent_id\":[\"0\"],\"roles[]\":[\"2\"],\"title\":[\"部门\"],\"uri\":[\"\"]}','2020-03-20 22:20:53','2020-03-20 22:20:53'),
	(12,1,'/admin/menu/edit/show','GET','127.0.0.1','','2020-03-20 22:20:56','2020-03-20 22:20:56'),
	(13,1,'/admin/menu/edit','POST','127.0.0.1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"d1b4e32a-999a-453c-ab1a-b688a576caf2\"],\"header\":[\"\"],\"icon\":[\"fa-bars\"],\"id\":[\"9\"],\"parent_id\":[\"0\"],\"roles[]\":[\"2\"],\"title\":[\"今日未提交名单\"],\"uri\":[\"/info/org_whitelist\"]}','2020-03-20 22:20:58','2020-03-20 22:20:58'),
	(14,1,'/admin/menu/new','POST','127.0.0.1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"c07dcc2f-4a78-48a3-ad66-1aeb00071887\"],\"header\":[\"\"],\"icon\":[\"fa-users\"],\"parent_id\":[\"\"],\"title\":[\"部门\"],\"uri\":[\"/info/admin\"]}','2020-03-20 22:21:35','2020-03-20 22:21:35'),
	(15,1,'/admin/menu/edit/show','GET','127.0.0.1','','2020-03-20 22:21:39','2020-03-20 22:21:39'),
	(16,1,'/admin/menu/edit','POST','127.0.0.1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"b2e7571b-e1be-4dd5-a7ab-8c67a54e39f2\"],\"header\":[\"\"],\"icon\":[\"fa-users\"],\"id\":[\"12\"],\"parent_id\":[\"0\"],\"title\":[\"部门\"],\"uri\":[\"/info/admin_user\"]}','2020-03-20 22:21:47','2020-03-20 22:21:47'),
	(17,1,'/admin/menu/edit/show','GET','127.0.0.1','','2020-03-20 22:21:53','2020-03-20 22:21:53'),
	(18,1,'/admin/menu/edit','POST','127.0.0.1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"fc6c6079-0d57-4164-98e6-6dfb80fcd034\"],\"header\":[\"\"],\"icon\":[\"fa-users\"],\"id\":[\"12\"],\"parent_id\":[\"10\"],\"title\":[\"部门管理员\"],\"uri\":[\"/info/admin_user\"]}','2020-03-20 22:22:19','2020-03-20 22:22:19'),
	(19,1,'/admin/menu/new','POST','127.0.0.1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"fe09ef9e-4404-4231-b0c3-0d4157d59469\"],\"header\":[\"\"],\"icon\":[\"fa-user\"],\"parent_id\":[\"\"],\"roles[]\":[\"2\"],\"title\":[\"小程序绑定用户\"],\"uri\":[\"/info/wx_mp_bind_info\"]}','2020-03-20 22:23:03','2020-03-20 22:23:03'),
	(20,1,'/admin/menu','GET','127.0.0.1','','2020-03-20 22:23:30','2020-03-20 22:23:30'),
	(21,1,'/admin/info/org_whitelist','GET','127.0.0.1','','2020-03-20 22:23:32','2020-03-20 22:23:32'),
	(22,1,'/admin/info/org_whitelist','GET','127.0.0.1','','2020-03-20 22:23:34','2020-03-20 22:23:34'),
	(23,1,'/admin/info/org_dep','GET','127.0.0.1','','2020-03-20 22:23:38','2020-03-20 22:23:38'),
	(24,1,'/admin/info/admin_user','GET','127.0.0.1','','2020-03-20 22:23:39','2020-03-20 22:23:39'),
	(25,1,'/admin/info/org_dep','GET','127.0.0.1','','2020-03-20 22:23:43','2020-03-20 22:23:43'),
	(26,1,'/admin/info/admin_user','GET','127.0.0.1','','2020-03-20 22:23:44','2020-03-20 22:23:44'),
	(27,1,'/admin/info/org_dep','GET','127.0.0.1','','2020-03-20 22:23:45','2020-03-20 22:23:45'),
	(28,1,'/admin/info/admin_user','GET','127.0.0.1','','2020-03-20 22:23:45','2020-03-20 22:23:45'),
	(29,1,'/admin/info/wx_mp_bind_info','GET','127.0.0.1','','2020-03-20 22:23:47','2020-03-20 22:23:47'),
	(30,1,'/admin/info/org_dep','GET','127.0.0.1','','2020-03-20 22:27:48','2020-03-20 22:27:48'),
	(31,1,'/admin/info/admin_user','GET','127.0.0.1','','2020-03-20 22:27:50','2020-03-20 22:27:50'),
	(32,1,'/admin/info/wx_mp_bind_info','GET','127.0.0.1','','2020-03-20 22:27:51','2020-03-20 22:27:51'),
	(33,1,'/admin/info/org_whitelist','GET','127.0.0.1','','2020-03-20 22:27:53','2020-03-20 22:27:53'),
	(34,1,'/admin/info/org_whitelist','GET','127.0.0.1','','2020-03-20 22:27:54','2020-03-20 22:27:54'),
	(35,1,'/admin/info/manager','GET','127.0.0.1','','2020-03-20 22:27:56','2020-03-20 22:27:56'),
	(36,1,'/admin/info/roles','GET','127.0.0.1','','2020-03-20 22:27:57','2020-03-20 22:27:57'),
	(37,1,'/admin/info/permission','GET','127.0.0.1','','2020-03-20 22:27:58','2020-03-20 22:27:58'),
	(38,1,'/admin/menu','GET','127.0.0.1','','2020-03-20 22:27:58','2020-03-20 22:27:58'),
	(39,1,'/admin/info/op','GET','127.0.0.1','','2020-03-20 22:27:59','2020-03-20 22:27:59'),
	(40,1,'/admin/info/org_whitelist','GET','127.0.0.1','','2020-03-20 22:28:00','2020-03-20 22:28:00'),
	(41,1,'/admin/menu','GET','127.0.0.1','','2020-03-20 22:32:05','2020-03-20 22:32:05'),
	(42,1,'/admin/info/org_whitelist','GET','127.0.0.1','','2020-03-20 22:32:43','2020-03-20 22:32:43'),
	(43,1,'/admin/info/org_whitelist','GET','127.0.0.1','','2020-03-20 22:32:44','2020-03-20 22:32:44'),
	(44,1,'/admin/info/org_whitelist','GET','127.0.0.1','','2020-03-20 22:34:15','2020-03-20 22:34:15'),
	(45,1,'/admin/menu','GET','127.0.0.1','','2020-03-20 22:34:19','2020-03-20 22:34:19');

/*!40000 ALTER TABLE `goadmin_operation_log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_permissions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_permissions`;

CREATE TABLE `goadmin_permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `http_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permissions_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_permissions` DISABLE KEYS */;

INSERT INTO `goadmin_permissions` (`id`, `name`, `slug`, `http_method`, `http_path`, `created_at`, `updated_at`)
VALUES
	(1,'All permission','*','','*','2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,'Dashboard','dashboard','GET,PUT,POST,DELETE','/','2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(3,'Basic','basic','','/info/org_whitelist','2020-03-21 12:09:46','2020-03-21 12:11:37');

/*!40000 ALTER TABLE `goadmin_permissions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_role_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_role_menu`;

CREATE TABLE `goadmin_role_menu` (
  `role_id` int(11) unsigned NOT NULL,
  `menu_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_role_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_role_menu` DISABLE KEYS */;

INSERT INTO `goadmin_role_menu` (`role_id`, `menu_id`, `created_at`, `updated_at`)
VALUES
	(1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(1,7,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,7,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(1,8,'2019-09-11 10:20:55','2019-09-11 10:20:55'),
	(2,8,'2019-09-11 10:20:55','2019-09-11 10:20:55'),
	(2,11,'2020-03-20 22:20:46','2020-03-20 22:20:46'),
	(2,10,'2020-03-20 22:20:53','2020-03-20 22:20:53'),
	(2,9,'2020-03-20 22:20:58','2020-03-20 22:20:58'),
	(2,13,'2020-03-20 22:23:03','2020-03-20 22:23:03'),
	(2,14,'2020-03-20 22:35:24','2020-03-20 22:35:24'),
	(2,12,'2020-03-21 12:01:06','2020-03-21 12:01:06'),
	(1,15,'2020-03-21 12:18:40','2020-03-21 12:18:40');

/*!40000 ALTER TABLE `goadmin_role_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_role_permissions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_role_permissions`;

CREATE TABLE `goadmin_role_permissions` (
  `role_id` int(11) unsigned NOT NULL,
  `permission_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_role_permissions` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_role_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_role_permissions` DISABLE KEYS */;

INSERT INTO `goadmin_role_permissions` (`role_id`, `permission_id`, `created_at`, `updated_at`)
VALUES
	(1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(1,2,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,2,'2020-03-21 12:09:57','2020-03-21 12:09:57'),
	(2,3,'2020-03-21 12:09:57','2020-03-21 12:09:57');

/*!40000 ALTER TABLE `goadmin_role_permissions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_role_users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_role_users`;

CREATE TABLE `goadmin_role_users` (
  `role_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_roles` (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_role_users` WRITE;
/*!40000 ALTER TABLE `goadmin_role_users` DISABLE KEYS */;

INSERT INTO `goadmin_role_users` (`role_id`, `user_id`, `created_at`, `updated_at`)
VALUES
	(1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,2,'2020-03-21 12:10:40','2020-03-21 12:10:40');

/*!40000 ALTER TABLE `goadmin_role_users` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_roles
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_roles`;

CREATE TABLE `goadmin_roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_roles_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_roles` WRITE;
/*!40000 ALTER TABLE `goadmin_roles` DISABLE KEYS */;

INSERT INTO `goadmin_roles` (`id`, `name`, `slug`, `created_at`, `updated_at`)
VALUES
	(1,'Administrator','administrator','2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,'Operator','operator','2019-09-10 00:00:00','2020-03-21 12:09:57');

/*!40000 ALTER TABLE `goadmin_roles` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_session
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_session`;

CREATE TABLE `goadmin_session` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `sid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `values` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `goadmin_session` WRITE;
/*!40000 ALTER TABLE `goadmin_session` DISABLE KEYS */;

INSERT INTO `goadmin_session` (`id`, `sid`, `values`, `created_at`, `updated_at`)
VALUES
	(24,'7b2ab90f-05a9-4880-b05e-ff564642003d','{\"user_id\":1}','2020-03-22 14:10:36','2020-03-22 14:10:36');

/*!40000 ALTER TABLE `goadmin_session` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_user_permissions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_user_permissions`;

CREATE TABLE `goadmin_user_permissions` (
  `user_id` int(11) unsigned NOT NULL,
  `permission_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_permissions` (`user_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_user_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_user_permissions` DISABLE KEYS */;

INSERT INTO `goadmin_user_permissions` (`user_id`, `permission_id`, `created_at`, `updated_at`)
VALUES
	(1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,2,'2020-03-21 12:10:40','2020-03-21 12:10:40'),
	(2,3,'2020-03-21 12:10:40','2020-03-21 12:10:40');

/*!40000 ALTER TABLE `goadmin_user_permissions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_users`;

CREATE TABLE `goadmin_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_users_username_unique` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_users` WRITE;
/*!40000 ALTER TABLE `goadmin_users` DISABLE KEYS */;

INSERT INTO `goadmin_users` (`id`, `username`, `password`, `name`, `avatar`, `remember_token`, `created_at`, `updated_at`)
VALUES
	(1,'admin','$2a$10$wj4S9VKgkjM6K2WUdg7I8O8zD1BASfcqyLfK3nSygQUXkwwYW31Ni','admin','','tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh','2019-09-10 00:00:00','2019-09-10 00:00:00'),
	(2,'operator','$2a$10$QiqHSbZX.Ti0i42xQZIr3O3/3ZRXpMloVQ9eNF.G.mxm7vhHwgEya','Operator','',NULL,'2019-09-10 00:00:00','2020-03-21 12:10:40');

/*!40000 ALTER TABLE `goadmin_users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
