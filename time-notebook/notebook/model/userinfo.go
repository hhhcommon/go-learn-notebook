// CREATE TABLE `userinfo` (
// 	`user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户 ID',
// 	`username` varchar(256) NOT NULL COMMENT '用户昵称',
// 	`header_img` varchar(256) NOT NULL COMMENT '用户头像',
// 	`bio` longtext NOT NULL COMMENT 'bio内容',
// 	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
// 	`update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
// 	 PRIMARY KEY ( `user_id` )
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

package model

import (
	"time"
)

// UserInfo 用户信息定义
type UserInfo struct {
	UserID     int64     `db:"user_id"`
	UserName   string    `db:"username"`
	HeaderImg  string    `db:"header_img"`
	Bio        string    `db:"bio"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}
