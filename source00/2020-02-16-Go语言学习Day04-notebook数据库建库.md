---
title: "Go语言学习Day04-Notebook数据库建库"
date: 2020-02-16T16:42:43+08:00
draft: true
---

# Notebook数据库建库

### 1. 资料

* [MySQL UPDATE 更新 | 菜鸟教程](https://www.runoob.com/mysql/mysql-update-query.html)
* [MySQL 列属性修改操作_数据库_u012149181的博客-CSDN博客](https://blog.csdn.net/u012149181/article/details/80336793)
* [胶囊日记](http://timepill.net/)
* [GORM入门指南 | 李文周的博客](https://www.liwenzhou.com/posts/Go/gorm/)

### 2. 建立数据库

#### 2.1 数据库整体规划

![](https://tva1.sinaimg.cn/large/0082zybpgy1gbydkwmd72j31l40k4tb4.jpg)

**notebook.sql:**


```sql
DROP DATABASE IF EXISTS `notebook`;

CREATE DATABASE `notebook` DEFAULT CHARSET utf8;

USE `notebook`;

DROP TABLE IF EXISTS `user`;

CREATE TABLE  IF NOT EXISTS `user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `email` varchar(128) NOT NULL COMMENT '邮箱',
  `password` varchar(256) NOT NULL COMMENT '密码',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
   `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1，会员为2，试用3，过期4',
    PRIMARY KEY ( `user_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `userinfo`;

CREATE TABLE `userinfo` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户 ID',
  `username` varchar(256) NOT NULL COMMENT '用户昵称',
  `header_img` varchar(256) NOT NULL COMMENT '用户头像',
  `bio` longtext NOT NULL COMMENT 'bio内容',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
   PRIMARY KEY ( `user_id` )
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `note`;

CREATE TABLE `note` (
  `note_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'note ID',
  `category_id` bigint(20) NOT NULL COMMENT '分类 ID',
  `content` longtext NOT NULL COMMENT 'note内容',
  `img` varchar(1024) NOT NULL COMMENT 'note图片',
  `view_count` int(255) NOT NULL COMMENT '阅读次数',
  `comment_count` int(255) NOT NULL COMMENT '评论次数',
  `user_id` bigint(20) NOT NULL  COMMENT '用户id',
  `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1，删除2，和谐3',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`note_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `comment_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论 ID',
  `note_id` bigint(20) NOT NULL  COMMENT 'note ID',
  `user_id` bigint(20) NOT NULL  COMMENT '作者 用户 ID',
  `feed_user_id` bigint(20) NOT NULL  COMMENT '被艾特的 用户 ID',
  `content` text NOT NULL COMMENT '评论内容',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `status` int(255) NOT NULL DEFAULT '1' COMMENT '评论状态: 正常1, 删除2',
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

1. user 表，可以做登录注册了。
2. note 表，可以做 note的 新建 修改 删除了。
3. comment 表， 可以做评论了。
4. userinfo 表，可以个人主页了。

### 3. 功能规划-注册功能

user表增加一条就是注册



#### 3.1 网站首页：

http://timepill.panda8z.com/

进入首页 index.html

路由为 `/`



#### 3.2 注册功能

路由：`http://timepill.panda8z.com/join?source=header-home`  附带了注册来源参数。

请求方式：注册页面 `GET` ；注册功能 `POST`

请求参数：

* `email` 用户电子邮箱；
* `password` 用户填写的密码；

加密方式： password sha256加密后 post给后端



**以下是github注册页面，供参考使用**

![image-20200216205300784](https://tva1.sinaimg.cn/large/0082zybpgy1gbyjan060tj311n0u0tdz.jpg)

### 4. 功能规划-登录功能

查询user表是否存在一条记录，存在即登录成功。

路由： `https://timepill.panda8z.com/login`

请求方式：登录页面 `GET` ；登录功能 `POST`

请求参数：

* `email` 用户电子邮箱；
* `password` 用户填写的密码；

加密方式： password sha256加密后 post给后端



**github登录页面供参考使用：**

![image-20200216210131784](https://tva1.sinaimg.cn/large/0082zybpgy1gbyjjgsewxj30hu0ny0tw.jpg)

### 5. 功能规划-用户信息展示

user 联合 userinfo 表 可以做出用户信息展示。

### 6. 功能规划-新建note

user 联合 note 表可以新增 note

note展示界面