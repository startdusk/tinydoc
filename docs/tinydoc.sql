SET NAMES utf8mb4;

SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS tinydoc;

--文档表

DROP TABLE IF EXISTS `doc`;

CREATE TABLE
    `doc` (
        `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
        `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文档标题',
        `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文档简介',
        `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '文档状态 1:未发布,2:已发布,3:已禁用',
        `version` decimal(10, 2) DEFAULT '0.01' COMMENT '文档版本号',
        `internal_version` int unsigned DEFAULT 0 COMMENT '内部文档版本号, 每当里面有文章更新, 就往上加1',
        `index` int(10) unsigned DEFAULT NULL COMMENT '排序字段',
        `created_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
        `updated_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
        `deleted_time` datetime(3) DEFAULT NULL COMMENT '删除时间',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

--文档细节表

DROP TABLE IF EXISTS `detail`;

CREATE TABLE
    `detail` (
        `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
        `doc_id` int(10) unsigned NOT NULL COMMENT '关联文档id, 文档一对多文章' `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
        `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
        `content` LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文章内容 放Markdown文本',
        `type` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '文章类型 1:文章, 2:文件夹',
        `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '文章状态 1:启用, 2:禁用',
        `index` int(10) unsigned DEFAULT NULL COMMENT '排序字段: 0:默认首页',
        `created_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
        `updated_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
        `deleted_time` datetime(3) DEFAULT NULL COMMENT '删除时间',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

--文档细节历史版本表(每更新一次文档细节就写一条记录保存到文档细节历史版本)

DROP TABLE IF EXISTS `history`;

CREATE TABLE
    `history` (
        `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
        `doc_id` int(10) unsigned NOT NULL COMMENT '关联文档id, 文档一对多文章' `detail_id` int(10) unsigned NOT NULL COMMENT '关联文档细节id' `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
        `content` LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文章内容 放Markdown文本',
        `internal_version` int unsigned DEFAULT 0 COMMENT '内部文档版本号, 关联doc内部版本号',
        `created_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
        `updated_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
        `deleted_time` datetime(3) DEFAULT NULL COMMENT '删除时间',
        PRIMARY KEY (`id`) UNIQUE KEY `uni_history` (`doc_id`, `detail_id`) USING BTREE
    ) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;