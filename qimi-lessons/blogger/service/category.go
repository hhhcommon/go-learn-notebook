package service

import (
	"blogger/dao/db"
	"blogger/model"
	"log"
)

// GetAllCategoryList 获取所有分类的列表
func GetAllCategoryList() (list []*model.Category, err error)  {
	list, err = db.GetAllCategoryList()
	if err != nil {
		log.Fatal("get  all  category List error！")
	}
	return
}
