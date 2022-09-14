create schema if not exists demo;
use demo;
CREATE TABLE `user`
(
    `user_id`          int unsigned NOT NULL AUTO_INCREMENT,
    `guid`             varchar(100) NOT NULL DEFAULT '' COMMENT '用户唯一标志 即token',
    `forbidden_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0正常1禁用',
    `created_at`       datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `udx_guid` (`guid`),
    KEY                `idx_created_at` (`created_at`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户';


