drop
database if exists `little_blog`;
drop
user if exists 'az'@'localhost';
create
database `little_blog` default character set utf8mb4;
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
    `topic`         varchar(200) NOT NULL COMMENT '文章主题',
    `thumbnail`     varchar(200) DEFAULT NULL COMMENT '缩略图url',
    `content`     text NOT NULL COMMENT '文章内容',
    `publish_state` smallint     DEFAULT NULL COMMENT '发布状态',
    `publish_time`  timestamp    DEFAULT NULL COMMENT '发布时间',
    `create_user`   varchar(32)  DEFAULT NULL COMMENT '创建人',
    `create_time`   timestamp    DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_user`   varchar(32)  DEFAULT NULL,
    `update_time`   timestamp    DEFAULT CURRENT_TIMESTAMP,
    `is_delete`     smallint     DEFAULT '0' COMMENT '是否已删除(0-未删除,1-已删除)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `tb_article_topic_uindex` (`topic`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='博客文章表';


/**
  参数表
 */
CREATE TABLE `tb_parameter`
(
    `id`          varchar(32)  NOT NULL COMMENT '主键id',
    `parent`      varchar(30)  NOT NULL COMMENT '父key',
    `key`         varchar(30)  NOT NULL COMMENT 'key',
    `value`       varchar(200) NOT NULL COMMENT 'value',
    `state`       smallint(1)       DEFAULT 0 COMMENT '状态',
    `create_user` varchar(32)       DEFAULT NULL COMMENT '创建人',
    `create_time` timestamp    NULL DEFAULT NULL COMMENT '创建时间',
    `update_user` varchar(32)       DEFAULT NULL COMMENT '更新人',
    `update_time` timestamp    NULL DEFAULT NULL COMMENT '更新时间',
    `is_delete`   smallint(1)       DEFAULT 0 COMMENT '是否已删除(0-未删除,1-已删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='参数表';

/**
  用户表
 */
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`
(
    `id`              int(11)      NOT NULL AUTO_INCREMENT,
    `username`        varchar(32)  NOT NULL COMMENT '用户名称',
    `password`        varchar(64)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `gender`          tinyint(3)   NOT NULL DEFAULT '1' COMMENT '性别：1 男,2 女, 3 未知',
    `birthday`        timestamp             DEFAULT NULL COMMENT '生日',
    `last_login_time` timestamp             DEFAULT NULL COMMENT '最近一次登录时间',
    `last_login_ip`   varchar(64)  NOT NULL DEFAULT '' COMMENT '最近一次登录IP地址',
    `nickname`        varchar(64)  NOT NULL DEFAULT '' COMMENT '用户昵称或网络名称',
    `mobile`          varchar(20)  NULL     DEFAULT '' COMMENT '用户手机号码',
    `avatar`          varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像图片',
    `state`           tinyint(3)   NOT NULL DEFAULT '1' COMMENT '1 活跃, 2 冻结, 3 废弃, 4 注销',
    `create_user`     varchar(32)           DEFAULT NULL COMMENT '创建人',
    `create_time`     timestamp    NULL     DEFAULT NULL COMMENT '创建时间',
    `update_user`     varchar(32)           DEFAULT NULL COMMENT '更新人',
    `update_time`     timestamp    NULL     DEFAULT NULL COMMENT '更新时间',
    `is_delete`       smallint(1)           DEFAULT 0 COMMENT '是否已删除(0-未删除,1-已删除)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_name` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

DROP TABLE IF EXISTS `tb_permission`;
CREATE TABLE `tb_permission`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `pid`         bigint(20)      DEFAULT NULL COMMENT '父级权限id',
    `name`        varchar(100)    DEFAULT NULL COMMENT '名称',
    `value`       varchar(200)    DEFAULT NULL COMMENT '权限值',
    `icon`        varchar(500)    DEFAULT NULL COMMENT '图标',
    `type`        int(1)          DEFAULT NULL COMMENT '权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）',
    `uri`         varchar(200)    DEFAULT NULL COMMENT '前端资源路径',
    `state`       int(1)          DEFAULT NULL COMMENT '启用状态；0->禁用；1->启用',
    `create_user` varchar(32)     DEFAULT NULL COMMENT '创建人',
    `create_time` timestamp  NULL DEFAULT NULL COMMENT '创建时间',
    `update_user` varchar(32)     DEFAULT NULL COMMENT '更新人',
    `update_time` timestamp  NULL DEFAULT NULL COMMENT '更新时间',
    `is_delete`   smallint(1)     DEFAULT 0 COMMENT '是否已删除(0-未删除,1-已删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='后台用户权限表';

DROP TABLE IF EXISTS `tb_role`;
CREATE TABLE `tb_role`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `name`        varchar(100)    DEFAULT NULL COMMENT '名称',
    `description` varchar(500)    DEFAULT NULL COMMENT '描述',
    `state`       int(1)          DEFAULT '1' COMMENT '启用状态：0->禁用；1->启用',
    `create_user` varchar(32)     DEFAULT NULL COMMENT '创建人',
    `create_time` timestamp  NULL DEFAULT NULL COMMENT '创建时间',
    `update_user` varchar(32)     DEFAULT NULL COMMENT '更新人',
    `update_time` timestamp  NULL DEFAULT NULL COMMENT '更新时间',
    `is_delete`   smallint(1)     DEFAULT 0 COMMENT '是否已删除(0-未删除,1-已删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='后台用户角色表';

DROP TABLE IF EXISTS `tb_role_permission_relation`;
CREATE TABLE `tb_role_permission_relation`
(
    `id`            bigint(20) NOT NULL AUTO_INCREMENT,
    `role_id`       varchar(32) DEFAULT NULL,
    `permission_id` varchar(32) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='后台用户角色和权限关系表';

DROP TABLE IF EXISTS `tb_user_role_relation`;
CREATE TABLE `tb_user_role_relation`
(
    `id`      bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` varchar(32) DEFAULT NULL,
    `role_id` varchar(32) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 17
  DEFAULT CHARSET = utf8 COMMENT ='后台用户和角色关系表';