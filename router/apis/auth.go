package apis

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jary-287/web-demo/models"
	"github.com/jary-287/web-demo/pkg/e"
	"github.com/jary-287/web-demo/pkg/logging"
	"github.com/jary-287/web-demo/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	var auth auth
	if err := c.BindJSON(&auth); err != nil {
		logging.Error("bind json failed", err.Error())
	}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&auth)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckToken(auth.Username, auth.Password)
		if isExist {
			token, err := util.GetToken(auth.Username, auth.Password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
