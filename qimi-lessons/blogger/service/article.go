package service

import (
	"blogger/dao/db"
	"blogger/model"
	"log"
)

// GetArticleRecordList 获取文章对应的分类
// @param pageNumber 页码  PageSize 每页的数量
func GetArticleRecordListByPageInfo(pageNumber, pageSize int) (recordList []*model.ArticleRecord, err error) {
	// 1. 获取文章列表
	articleList, err := db.GetArticleList(pageNumber, pageSize)
	if err != nil {
		log.Fatal("获取文章列表出错")
		return
	}

	if len(articleList) < 0 {
		log.Fatal("获取文章列表数量出错")
		return
	}

	// 2. 获取分类
	categoryIDs := GetCategoryIDsByArticleInfoList(articleList)

	categoryList, err := db.GetCategoryList(categoryIDs)
	if err != nil {
		log.Fatal("获取分类列表出错")
		return
	}

	// 3. 返回页面 聚合数据

	for _, articleInfo := range articleList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *articleInfo,
		}

		for _, category := range categoryList {
			if category.ID == articleInfo.CategoryID {
				articleRecord.Category = *category
				continue
			}
		}
		recordList = append(recordList, articleRecord)
	}
	return
}

// GetCategoryIDsByArticleInfoList 根据文章信息列表获取分类列表
func GetCategoryIDsByArticleInfoList(articleInfoList []*model.ArticleInfo) (categoryIDs []int64) {
	// 遍历文章列表拿出获得所有分类ID
	for _, articleInfo := range articleInfoList {
		categoryID := articleInfo.CategoryID

		for _, id := range categoryIDs {
			if id != categoryID {
				categoryIDs = append(categoryIDs, categoryID)
			}
		}
	}
	return
}

func GetArticleRecordListByCategoryID(categoryID, pageNumber, pageSize int) (articleRecodeList []*model.ArticleRecord, err error) {
	// 1. 获取文章列表
	articleList, err := db.GetArticleListByCategoryId(categoryID, pageNumber, pageSize)
	if err != nil {
		log.Fatal("获取文章列表出错")
		return
	}

	if len(articleList) < 0 {
		log.Fatal("获取文章列表数量出错")
		return
	}

	// 2. 获取分类
	categoryIDs := GetCategoryIDsByArticleInfoList(articleList)
	categoryList, err := db.GetCategoryList(categoryIDs)
	if err != nil {
		log.Fatal("获取分类列表出错")
		return
	}

	// 3. 返回页面 聚合数据

	for _, articleInfo := range articleList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *articleInfo,
		}
		// 找出词条文章记录对应的分类
		for _, category := range categoryList {
			if articleInfo.CategoryID == category.ID {
				articleRecord.Category = *category
			}
		}

		articleRecodeList = append(articleRecodeList, articleRecord)

	}
	return
}
