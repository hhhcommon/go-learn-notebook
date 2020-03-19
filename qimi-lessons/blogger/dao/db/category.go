package db

import (
	"blogger/model"

	"github.com/jmoiron/sqlx"
)

// InsertCategory 添加分类。使用整个结构体对象映射数据库的一个条目
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlstr := "insert into category(category_name, category_no) value(?,?)"
	result, err := DB.Exec(sqlstr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = result.LastInsertId()
	return
}

// GetCategoryByID 根据分类id获取单个分类对象
func GetCategoryByID(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlstr := "select id, category_name, category_no from category where id = ?"
	err = DB.Get(category, sqlstr, id)
	return
}

// GetCategoryList 根据多个id组成的切片获取多个分类
func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	sqlstr, args, err := sqlx.In("select id, category_name, category_no from category where id in(?) order by category_no asc", categoryIds)
	err = DB.Select(&categoryList, sqlstr, args...)
	return
}

// GetAllCategoryList 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sqlstr := "select id, category_name, category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlstr)
	return
}
