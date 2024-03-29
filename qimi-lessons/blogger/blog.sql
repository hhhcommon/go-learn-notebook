/*
SQLyog Ultimate v10.00 Beta1
MySQL - 5.5.58 : Database - blogger
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`blogger` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `blogger`;

/*Table structure for table `article` */

DROP TABLE IF EXISTS `article`;

CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `category_id` bigint(20) NOT NULL COMMENT '分类id',
  `content` longtext NOT NULL COMMENT '文章内容',
  `title` varchar(1024) NOT NULL COMMENT '文章标题',
  `view_count` int(255) NOT NULL COMMENT '阅读次数',
  `comment_count` int(255) NOT NULL COMMENT '评论次数',
  `username` varchar(128) NOT NULL COMMENT '作者',
  `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1',
  `summary` varchar(256) NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

/*Data for the table `article` */

insert  into `article`(`id`,`category_id`,`content`,`title`,`view_count`,`comment_count`,`username`,`status`,`summary`,`create_time`,`update_time`) values (1,1,'this a test ak dkdkdkddkddkd111','我是标题1',1,0,'Mr.Sun1',1,'我是\n									很多的\n									内容1','2019-10-04 23:34:06',NULL),(2,2,'this a test ak dkdkdkddkddkd222','我是标题2',1,0,'Mr.Sun2',1,'我是\n									很多的\n									内容2','2019-10-04 23:34:39',NULL),(3,2,'this a test ak dkdkdkddkddkd333','我是标题3',1,1,'Mr.Sun3',1,'我是\n									很多的\n									内容3','2019-10-04 23:34:55',NULL);

/*Table structure for table `category` */

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `category_name` varchar(255) NOT NULL COMMENT '分类名字',
  `category_no` int(10) NOT NULL COMMENT '分类排序',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

/*Data for the table `category` */

insert  into `category`(`id`,`category_name`,`category_no`,`create_time`,`update_time`) values (1,'css/html',1,'2019-08-12 10:55:45','2019-08-12 10:59:00'),(2,'后端开发',2,'2019-08-12 10:56:07','2019-08-12 10:59:03'),(3,'Java开发',3,'2019-08-12 10:56:16','2019-08-12 10:59:05'),(4,'C++开发',4,'2019-08-12 10:56:24','2019-08-12 10:59:08'),(5,'架构剖析',5,'2019-08-12 10:56:36','2019-08-12 10:59:10'),(6,'Golang开发',6,'2019-08-12 10:56:45','2019-08-12 10:59:14');

/*Table structure for table `comment` */

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `content` text NOT NULL COMMENT '评论内容',
  `username` varchar(64) NOT NULL COMMENT '评论作者',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `status` int(255) NOT NULL DEFAULT '1' COMMENT '评论状态: 0, 删除；1， 正常',
  `article_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `comment` */

/*Table structure for table `leave` */

DROP TABLE IF EXISTS `leave`;

CREATE TABLE `leave` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `leave` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
