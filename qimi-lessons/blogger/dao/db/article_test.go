package db

import (
	"blogger/model"
	"testing"
	"time"
)
// 初始化方法里初始化数据库，建立连接。
func init() {
	//参数parseTime=true：将 mysql中的时间类型，自动解析为go结构体中的时间类型。
	dns := "root:123456@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// TestInsertArticle 测试： 插入文章
func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryID = 5
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.CommentCount = 0
	article.Content = "来自：\n TestInsertArticle 测试： 插入文章"
	article.ArticleInfo.Summary = " TestInsertArticle 测试： 插入文章"
	article.ArticleInfo.Title = "TestONNNNN"
	article.ArticleInfo.UserName = "panda8z"
	article.ArticleInfo.ViewCount = 0
	id, err := InsertArticle(article)
	if err != nil {
		t.Logf(err.Error())
		return
	}
	t.Logf("article id = %#v\n", id)
}


func TestGetArticleInfo(t *testing.T) {
	articleInfo, err := GetArticleDetail(1)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Logf("get article succ, article:%#v\n", articleInfo)
}

func TestGetArticleDetail(t *testing.T) {

	detail, err := GetArticleDetail(1)
	if err != nil {
		t.Logf(err.Error())
		return
	}
	t.Logf("article detail : %#v\n", detail.Content)
}

// TestGetArticleList 测试：按页码获取文章列表
func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 20)
	if err != nil {
		return
	}
	t.Logf("articleList : %#v\n", len(articleList))
}
