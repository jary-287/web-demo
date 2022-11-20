package v1

import (
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jary-287/web-demo/models"
	"github.com/jary-287/web-demo/pkg/e"
	"github.com/jary-287/web-demo/pkg/logging"
	"github.com/jary-287/web-demo/pkg/setting"
	"github.com/jary-287/web-demo/pkg/util"
)

// @Summary 获取所有文章
// @Description 获取所有的文章
// @Tags 文章
// @Param id query string false "文章id"
// @Success 200 {object} []models.Article "desc"
// @Failure 400 {object} string "{"msg": "who are you"}"
// @Router /articles [get]
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	code := e.INVALID_PARAMS
	valid := validation.Validation{}
	var state = -1
	if c.Query("state") != "" {
		maps["state"], _ = strconv.Atoi(c.Query("state"))
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	var tagId = -1
	if c.Query("tag_id") != "" {
		maps["tag_id"], _ = strconv.Atoi(c.Query("tag_id"))
		valid.Min(tagId, 1, "tag_id").Message("标签大于0")
	}
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["data"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
	} else {
		for _, err := range valid.Errors {
			logging.Info("key:", err.Key, "err:", err.Message)
		}
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func GetArticle(c *gin.Context) {

}

func AddArticle(c *gin.Context) {
	var article models.Article
	if err := c.BindJSON(&article); err != nil {
		logging.Warn("bind json failed:", err.Error())
	}
	code := e.SUCCESS
	if models.ExsitById(article.TagID) {
		models.AddArticle(article)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})

}

func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var article models.Article
	if err := c.BindJSON(&article); err != nil {
		logging.Warn("bind json failed:", err.Error())
	}
	code := e.SUCCESS
	if !models.ExsitArticleById(id) {
		code = e.ERROR_NOT_EXIST_ARTICLE
	} else {
		models.EditArticle(id, article)
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := e.SUCCESS
	if models.ExsitArticleById(id) {
		models.DeleteArticle(id)
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})
}
