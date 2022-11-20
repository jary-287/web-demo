package models

import (
	"time"

	"github.com/jary-287/web-demo/pkg/logging"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (t *Tag) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
func GetTags(PageNum int, PageSize int, maps interface{}) (tags []Tag) {
	if result := db.Where(maps).Offset(PageNum).Limit(PageSize).Find(&tags); result.Error != nil {
		logging.Error("err:", result.Error)
	}
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExsitByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	return tag.ID > 0
}

func AddTag(tag Tag) bool {
	if result := db.Create(&tag); result.Error != nil {
		logging.Error(result.Error)
		return false
	}
	return true
}
func DeleteTag(id int) {
	db.Where("id=?", id).Delete(&Tag{})
}

func ExsitById(id int) bool {
	var tag Tag
	db.Where("id=?", id).First(&tag)
	return tag.ID > 0
}
func EditTag(id int, tag Tag) {
	db.Model(&Tag{}).Where("id=?", id).Updates(&tag)
}
