package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	TagID      int    `gorm:"index" json:"tag_id"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (a *Article) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (a *Article) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}

func GetArticle(id int) (article Article) {
	db.Where("id=?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func GetArticles(pageNum, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func AddArticle(article Article) {
	db.Create(&article)
}

func EditArticle(id int, article Article) {
	db.Model(&Article{}).Where("id=?", id).Updates(&article)
	//db.Model(&Article{}).Where("id=?", id).Update("state", 0)
	//TODO
	//if I update record use struct ,the statement can not be 0
}

func DeleteArticle(id int) {
	db.Where("id=?", id).Delete(&Article{})
}

func ExsitArticleById(id int) bool {
	var article Article
	db.Where("id=?", id).Find(&article)
	return article.ID > 0
}
