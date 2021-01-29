drop
database if exists `little_blog`;
drop
user if exists 'az'@'localhost';
create
database `little_blog` default character set UTF8MB4;
use
`little_blog`;
create
user 'az'@'%' identified by 'azusa520';
grant all privileges on `little_blog`.* to
'az'@'%';
flush
privileges;

/**
  文章表
 */
CREATE TABLE `tb_article`
(
    `id`            varchar(32)  NOT NULL COMMENT '主键id',
    `topic`         varchar(100) NOT NULL COMMENT '文章主题',
    `thumbnail`     varchar(200) DEFAULT NULL COMMENT '缩略图url',
    `content`       text         NOT NULL COMMENT '文章内容',
    `publish_state` smallint(1) DEFAULT NULL COMMENT '发布状态',
    `publish_time`  timestamp    DEFAULT NULL COMMENT '发布时间',
    `create_user`   varchar(32)  DEFAULT NULL COMMENT '创建人',
    `create_time`   timestamp    DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_user`   varchar(32)  DEFAULT NULL,
    `update_time`   timestamp    DEFAULT CURRENT_TIMESTAMP,
    `is_delete`     smallint(1) DEFAULT '0' COMMENT '是否已删除(0-未删除,1-已删除)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `tb_article_topic_uindex` (`topic`)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4 COMMENT ='文章表';


/**
  参数信息表
 */
CREATE TABLE `tb_parameter`
(
    `id`          varchar(32)  NOT NULL COMMENT '主键id',
    `parent`      varchar(30)  NOT NULL COMMENT '父key',
    `key`         varchar(30)  NOT NULL COMMENT 'key',
    `value`       varchar(200) NOT NULL COMMENT 'value',
    `state`       smallint(1) DEFAULT 0 COMMENT '状态',
    `create_user` varchar(32) DEFAULT NULL COMMENT '创建人',
    `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
    `update_user` varchar(32) DEFAULT NULL COMMENT '更新人',
    `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    `is_delete`   smallint(1) DEFAULT 0 COMMENT '是否已删除(0-未删除,1-已删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4 COMMENT ='参数信息表';

/**
  用户信息表
 */
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`
(
    `id`              varchar(32)  NOT NULL,
    `username`        varchar(50)  NOT NULL COMMENT '用户名称',
    `password`        varchar(64)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `gender`          smallint(1) NOT NULL DEFAULT 1 COMMENT '性别：1 男,2 女, 3 未知',
    `birthday`        timestamp             DEFAULT NULL COMMENT '生日',
    `last_login_time` timestamp             DEFAULT NULL COMMENT '最近一次登录时间',
    `last_login_ip`   varchar(64)  NOT NULL DEFAULT '' COMMENT '最近一次登录IP地址',
    `nickname`        varchar(64)  NOT NULL DEFAULT '' COMMENT '用户昵称或网络名称',
    `mobile`          varchar(20) NULL DEFAULT '' COMMENT '用户手机号码',
    `avatar`          varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像图片',
    `state`           smallint(1) NOT NULL DEFAULT 1 COMMENT '1 活跃, 2 冻结, 3 废弃, 4 注销',
    `create_user`     varchar(32)           DEFAULT NULL COMMENT '创建人',
    `create_time`     timestamp NULL DEFAULT NULL COMMENT '创建时间',
    `update_user`     varchar(32)           DEFAULT NULL COMMENT '更新人',
    `update_time`     timestamp NULL DEFAULT NULL COMMENT '更新时间',
    `is_delete`       smallint(1) DEFAULT 0 COMMENT '是否已删除(0-未删除,1-已删除)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_name` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4 COMMENT ='用户信息表';

/**
  权限信息表
 */
DROP TABLE IF EXISTS `tb_permission`;
CREATE TABLE `tb_permission`
(
    `id`          varchar(32) NOT NULL,
    `pid`         varchar(32)  DEFAULT NULL COMMENT '父级权限id',
    `name`        varchar(100) DEFAULT NULL COMMENT '名称',
    `value`       varchar(200) DEFAULT NULL COMMENT '权限值',
    `icon`        varchar(500) DEFAULT NULL COMMENT '图标',
    `type`        smallint(1) DEFAULT NULL COMMENT '权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）',
    `uri`         varchar(200) DEFAULT NULL COMMENT '前端资源路径',
    `state`       smallint(1) DEFAULT NULL COMMENT '启用状态；0->禁用；1->启用',
    `create_user` varchar(32)  DEFAULT NULL COMMENT '创建人',
    `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
    `update_user` varchar(32)  DEFAULT NULL COMMENT '更新人',
    `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    `is_delete`   smallint(1) DEFAULT 0 COMMENT '是否已删除(0-未删除,1-已删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4 COMMENT ='权限信息表';

/**
  角色信息表
 */
DROP TABLE IF EXISTS `tb_role`;
CREATE TABLE `tb_role`
(
    `id`          varchar(32) NOT NULL,
    `name`        varchar(100) DEFAULT NULL COMMENT '名称',
    `description` varchar(500) DEFAULT NULL COMMENT '描述',
    `state`       smallint(1) DEFAULT 0 COMMENT '启用状态：0->启用；1->禁用',
    `create_user` varchar(32)  DEFAULT NULL COMMENT '创建人',
    `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
    `update_user` varchar(32)  DEFAULT NULL COMMENT '更新人',
    `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    `is_delete`   smallint(1) DEFAULT 0 COMMENT '是否已删除(0-未删除,1-已删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4 COMMENT ='角色信息表';

/**
  角色权限关联表
 */
DROP TABLE IF EXISTS `tb_role_permission_relation`;
CREATE TABLE `tb_role_permission_relation`
(
    `id`            varchar(32) NOT NULL,
    `role_id`       varchar(32) DEFAULT NULL,
    `permission_id` varchar(32) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4 COMMENT ='角色权限关联表';

/**
  用户角色关联表
 */
DROP TABLE IF EXISTS `tb_user_role_relation`;
CREATE TABLE `tb_user_role_relation`
(
    `id`      varchar(32) NOT NULL,
    `user_id` varchar(32) DEFAULT NULL,
    `role_id` varchar(32) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4 COMMENT ='用户角色关联表';

/**
  文章评论表
 */
DROP TABLE IF EXISTS `tb_comment`;
CREATE TABLE `tb_comment`
(
    `id`          varchar(32) NOT NULL,
    `article_id`  varchar(32) NOT NULL COMMENT '文章id',
    `pid`         varchar(32) NULL,
    `nickname`    varchar(100) DEFAULT NULL COMMENT '昵称',
    `email`       varchar(100) DEFAULT NULL COMMENT '联系email',
    `content`     text         DEFAULT NULL COMMENT '评论内容',
    `create_user` varchar(32)  DEFAULT NULL COMMENT '创建人',
    `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
    `update_user` varchar(32)  DEFAULT NULL COMMENT '更新人',
    `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    `is_delete`   smallint(1) DEFAULT 0 COMMENT '是否已删除(0-未删除,1-已删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4 COMMENT ='文章评论表';
