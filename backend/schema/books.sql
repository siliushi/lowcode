-- CREATE DATABASE IF NOT EXISTS lowcodedb default charset utf8 COLLATE utf8_general_ci;
-- use lowcodedb;

DROP TABLE IF EXISTS `book`;

CREATE TABLE IF NOT EXISTS `book`  (
  `ysid` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  `price` DECIMAL(6,2) NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`ysid`),
  UNIQUE KEY `unq_book_name` (`name`),
  INDEX `idx_book_deleted_at`(`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user`;

CREATE TABLE IF NOT EXISTS `user` (
  `ysid` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  `real_name` varchar(128),
  `pwd_hash` varchar(128) ,
  `email` varchar(128),
  `phone` varchar(16),
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`ysid`),
  INDEX `idx_user_deleted_at`(`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `user` (`name`, `real_name`, `pwd_hash`, `email`, `phone`, `created_at`, `updated_at`) VALUES ('admin', '', '$2a$10$inPdnVt8Pg6HT0HQA/2teucS4BkaOoxbctm8LZgQHt.BT9Vs5L1za', '', '', '2020-06-12 17:29:34', '2020-06-12 17:29:34');