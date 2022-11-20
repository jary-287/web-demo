package v1

import (
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jary-287/web-demo/models"
	"github.com/jary-287/web-demo/pkg/e"
	"github.com/jary-287/web-demo/pkg/setting"
	"github.com/jary-287/web-demo/pkg/util"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	if arg := c.Query("state"); arg != "" {
		state, _ := strconv.Atoi(arg)
		maps["state"] = state
	}
	code := e.SUCCESS
	data["data"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	var tag models.Tag
	c.BindJSON(&tag)
	code := e.INVALID_PARAMS
	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("name 不能为空")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")
	if !valid.HasErrors() {
		if models.ExsitByName(tag.Name) {
			code = e.ERROR_EXIST_TAG
		} else {
			models.AddTag(tag)
			code = e.SUCCESS
		}
	}

	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})

}

//修改文章标签
func EditTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	var tag models.Tag
	c.BindJSON(&tag)
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if !models.ExsitById(id) {
			code = e.ERROR_NOT_EXIST_TAG
		} else {
			models.EditTag(id, tag)
		}
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if !models.ExsitById(id) {
			code = e.ERROR_NOT_EXIST_TAG
		} else {
			models.DeleteTag(id)
		}
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})

}
