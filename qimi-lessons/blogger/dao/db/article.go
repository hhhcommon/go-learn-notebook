package db

import (
	"blogger/model"
	"fmt"
)

// InsertArticle 插入文章： 使用文章结构体对象映射数据库条目
func InsertArticle(article *model.ArticleDetail) (articleID int64, err error) {
	if article == nil {
		return
	}

	sqlstr := `insert into
	 article(content, summary, title, username, category_id, view_count, comment_count) values (?,?,?,?,?,?,?)`

	result, err := DB.Exec(sqlstr,
		article.Content,
		article.ArticleInfo.Summary,
		article.ArticleInfo.Title,
		article.ArticleInfo.UserName,
		article.ArticleInfo.CategoryID,
		article.ArticleInfo.ViewCount,
		article.ArticleInfo.CommentCount,
	)

	if err != nil {
		return
	}

	articleID, err = result.LastInsertId()
	return
}

// GetArticleList 获取文章信息列表： 根据页面数量和页码获取相应数据。

func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum <= 0 || pageSize <= 0 {
		return
	}

	sqlstr := `select 
					id, summary, title, view_count, create_time, comment_count, username, category_id
				from 
					article
				where
					status = 1
				order by create_time desc
				limit ?,?`
	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	return

}

// GetArticleDetail 获取文章详情
// `id`,`category_id`,`content`,`title`,`view_count`,`comment_count`,`username`,`status`,`summary`,`create_time`,`update_time`
func GetArticleDetail(articleID int64) (articleDetail *model.ArticleDetail, err error) {
	if articleID < 0 {
		err = fmt.Errorf("invalid parameter,article_id:%d", articleID)
		return
	}
	// 声明性返回值竟然在这里不能用  
	// 需要提前声明一块内存才能取址
	articleDetail = &model.ArticleDetail{}

	sqlstr := ` select  
		id, category_id, content, title, view_count, comment_count, summary, create_time, username 
	from
		article
	where 
		id = ?
	and 
		status = 1
	`
	err = DB.Get(articleDetail, sqlstr, articleID)
	return
}

func GetArticleListByCategoryId(categoryId,
	pageNum,
	pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlstr := `select 
						id, summary, title, view_count,
						 create_time, comment_count, username, category_id
					from 
						article 
					where 
						status = 1
						and
						category_id = ?
					order by create_time desc
					limit ?, ?`
	err = DB.Select(&articleList, sqlstr, categoryId, pageNum, pageSize)
	return
}
