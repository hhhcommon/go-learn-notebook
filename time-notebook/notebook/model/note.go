// CREATE TABLE `note` (
//   `note_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'note ID',
//   `category_id` bigint(20) NOT NULL COMMENT '分类 ID',
//   `content` longtext NOT NULL COMMENT 'note内容',
//   `img` varchar(1024) NOT NULL COMMENT 'note图片',
//   `view_count` int(255) NOT NULL COMMENT '阅读次数',
//   `comment_count` int(255) NOT NULL COMMENT '评论次数',
//   `user_id` bigint(20) NOT NULL  COMMENT '用户id',
//   `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1，删除2，和谐3',
//   `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
//   `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
//   PRIMARY KEY (`note_id`)
// ) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

package model

import (
	"time"
)

// Note  定义
type Note struct {
	NoteID       int64     `db:"note_id"`
	CategorID    int64     `db:"category_id"`
	Content      string    `db:"content"`
	Img          string    `db:"img"`
	ViewCount    uint32    `db:"view_count"`
	CommentCount uint32    `db:"comment_count"`
	UserID       int64     `db:"user_id"`
	Status       int16     `db:"status"`
	CreateTime   time.Time `db:"create_time"`
	UpdateTime   time.Time `db:"update_time"`
}
