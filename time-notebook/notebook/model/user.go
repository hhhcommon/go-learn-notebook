// CREATE TABLE  IF NOT EXISTS `user` (
// 	`user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户id',
// 	`email` varchar(128) NOT NULL COMMENT '邮箱',
// 	`password` varchar(256) NOT NULL COMMENT '密码',
// 	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
// 	`update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
// 	 `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1，会员为2，试用3，过期4',
// 	  PRIMARY KEY ( `user_id` )
//   )ENGINE=InnoDB DEFAULT CHARSET=utf8;

package model

import (
	"time"
)

// User 定义
type User struct {
	UserID     int64  `db:"user_id"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	CreateTime time.Time `db:"create_time`
	UpdateTime time.Time `db:"update_time"`
	Status     int16  `db:"status"`
}
