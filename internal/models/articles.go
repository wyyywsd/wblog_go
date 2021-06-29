package models

import (
	"fmt"
	"wblogApi/internal/db"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UserId         uint
	ArticleTitle   string `gorm:"not null" form:"article_title" json:"article_title" binding:"required"`
	ArticleContent string `gorm:"not null" form:"article_content" json:"article_content" binding:"required"`
}

func CreateArticle(articleTitle string, articleContent string, userID uint) *gorm.DB {
	article := Article{ArticleTitle: articleTitle, ArticleContent: articleContent, UserId: userID}
	result := db.W_Db.Create(&article) // 通过数据的指针来创建
	return result
}
func FindArticleById(articleId string) (*Article, *gorm.DB) {
	var article Article
	result := db.W_Db.Where("id = ?", articleId).First(&article)
	//err := db.W_Db.Table("users").Where("user_name = ?", userName).Find(&user).Error
	fmt.Println(result)
	return &article, result
}
func FindAllArticles(pageSize int, currentPage int) ([]*Article, *gorm.DB, int64) {
	var articles []*Article
	var allArticles []*Article
	result := db.W_Db.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&articles)
	resultAll := db.W_Db.Find(&allArticles)
	return articles, result, resultAll.RowsAffected
}
