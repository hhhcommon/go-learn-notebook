package db

import "testing"

// 初始化方法里初始化数据库，建立连接。
func init() {
	//参数parseTime=true：将 mysql中的时间类型，自动解析为go结构体中的时间类型。
	dns := "root:123456@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {

		panic(err)
	}
}

// 根据ID查询分类
func TestGetCategoryByID(t *testing.T) {
	category, err := GetCategoryByID(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category: %#v", category)
}

// 根据ID切片查询分类切片
func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}

	for index, value := range categoryList {
		t.Logf("category_list %v ： %#v\n", index, value)
	}
	t.Logf("category_list ： %#v", categoryList)
}

func TestGetAllCategoryList(t *testing.T){
	categoryList, err := GetAllCategoryList()
	if err != nil {
		panic(err)
	}

	for index, value := range categoryList {
		t.Logf("category_list %v ： %#v\n", index, value)
	}
	t.Logf("category_list ： %#v", len(categoryList))
}
