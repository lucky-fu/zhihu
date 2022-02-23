DROP DATABASE IF EXISTS zhihu;

CREATE DATABASE zhihu;
use zhihu


DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    phone int(11) UNIQUE,
    nickname varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    avatar_url varchar(100) NOT NULL DEFAULT '/static/images/default.jpg',
   `status` int DEFAULT '0',
    marked_count int(11) unsigned DEFAULT 0, 
    question_count int(11) unsigned DEFAULT 0,
    answer_count int(11) unsigned DEFAULT 0,
    create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modify_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS questions;

CREATE TABLE questions (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    user_id int(11) unsigned NOT NULL,
    title varchar(100) NOT NULL DEFAULT '',
     `status` int DEFAULT '0',
    detail longtext,
    answer_count int(11) unsigned NOT NULL DEFAULT 0,
    comment_count int(11) unsigned NOT NULL DEFAULT 0,
    create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modify_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8;

CREATE TABLE `answers` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `question_id` int unsigned NOT NULL,
  `user_id` int unsigned NOT NULL,
  `content` longtext,
  `marked_count` int unsigned DEFAULT '0',
  `comment_count` int unsigned DEFAULT '0',
  `status` int DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `question_id` (`question_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


CREATE TABLE answer_comments (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    user_id int(11) unsigned NOT NULL,
    answer_id int(11) unsigned NOT NULL,
    create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modify_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    content longtext,
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8;

CREATE TABLE `member_followers` (
  `member_id` int unsigned NOT NULL,
  `follower_id` int unsigned NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`member_id`,`follower_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci