SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for zta_client
-- ----------------------------
CREATE TABLE IF NO EXISTS `zta_client` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'name',
  `server_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `uuid` varchar(40) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'uuid',
  `port` int(10) unsigned NOT NULL DEFAULT '0',
  `expire` int(11) unsigned NOT NULL DEFAULT '0',
  `relay` json DEFAULT NULL COMMENT 'relay',
  `server` json DEFAULT NULL COMMENT 'server',
  `target` json DEFAULT NULL COMMENT 'target',
  `ca_pem` varchar(4000) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `cert_pem` varchar(3000) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_uuid` (`uuid`) USING BTREE,
  KEY `idx_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='client';

-- ----------------------------
-- Table structure for zta_oauth2
-- ----------------------------
CREATE TABLE IF NO EXISTS `zta_oauth2` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `company` varchar(100) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'company',
  `client_id` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'client_id',
  `client_secret` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'client_secret',
  `redirect_url` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'redirect_url',
  `scopes` json DEFAULT NULL COMMENT '域',
  `auth_url` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'auth_url',
  `token_url` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'token_url',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_company` (`company`) USING BTREE,
  KEY `idx_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='oauth2';

-- ----------------------------
-- Table structure for zta_relay
-- ----------------------------
CREATE TABLE IF NO EXISTS `zta_relay` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'name',
  `uuid` varchar(40) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'uuid',
  `host` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `port` int(10) unsigned NOT NULL DEFAULT '0',
  `out_port` int(10) unsigned NOT NULL DEFAULT '0',
  `ca_pem` varchar(4000) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `cert_pem` varchar(3000) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_uuid` (`uuid`) USING BTREE,
  KEY `idx_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='relay';

-- ----------------------------
-- Table structure for zta_resource
-- ----------------------------
CREATE TABLE IF NO EXISTS `zta_resource` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'name',
  `uuid` varchar(40) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'uuid',
  `type` varchar(40) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'dns,cidr',
  `host` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `port` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_uuid` (`uuid`) USING BTREE,
  KEY `idx_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='resource';

-- ----------------------------
-- Table structure for zta_server
-- ----------------------------
CREATE TABLE IF NO EXISTS `zta_server` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `resource_id` varchar(100) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'resource id , split by ","',
  `name` varchar(100) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'name',
  `uuid` varchar(40) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'uuid',
  `host` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `port` int(10) unsigned NOT NULL DEFAULT '0',
  `out_port` int(10) unsigned NOT NULL DEFAULT '0',
  `ca_pem` varchar(4000) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `cert_pem` varchar(3000) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_uuid` (`uuid`) USING BTREE,
  KEY `idx_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='server';

-- ----------------------------
-- Table structure for zta_user
-- ----------------------------
CREATE TABLE IF NO EXISTS `zta_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(100) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'email',
  `uuid` varchar(40) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'uuid',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_uuid` (`uuid`) USING BTREE,
  KEY `idx_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='user';

SET FOREIGN_KEY_CHECKS = 1;
