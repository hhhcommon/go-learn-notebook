package model

//`id`,`category_name`,`category_no`,`create_time`,`update_time`

// Category 分类结构体
type Category struct {
	ID           int64  `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo   int64  `db:"category_no"`
}
