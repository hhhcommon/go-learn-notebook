package model

import (
	"time"
)

// 定义文章结构体
// 需要匹配数据库的结构
// 并制定数据库的字段标签

// `id`,`category_id`,`content`,`title`,`view_count`,`comment_count`,`username`,`status`,`summary`,`create_time`,`update_time`

// ArticleInfo 文章信息的结构体
type ArticleInfo struct {
	ID           int64     `db:"id"`
	CategoryID   int64     `db:"category_id"`
	Title        string    `db:"title"`
	Summary      string    `db:"summary"`    // 文章摘要
	ViewCount    uint32    `db:"view_count"` // 查看次数
	CommentCount uint32    `db:"comment_count"`
	UserName     string    `db:"username"`
	CreateTime   time.Time `db:"create_time"` // 创建时间

}

// ArticleDetail 结构体，包含某article的全部内容
type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}

// ArticleRecord 整个文章记录
type ArticleRecord struct {
	ArticleInfo
	Category
}
