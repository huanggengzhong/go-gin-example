package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/go-gin-example/models"
	"github.com/huanggengzhong/go-gin-example/pkg/e"
	"github.com/huanggengzhong/go-gin-example/pkg/logging"
	"github.com/huanggengzhong/go-gin-example/pkg/setting"
	"github.com/huanggengzhong/go-gin-example/pkg/util"
	"github.com/unknwon/com"
)

// ShowAccount godoc
//
//	@Summary		列表
//	@Description	获取标签列表
//	@Tags			标签
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	false	"标签名"
//	@Param			state	path		int	false	"状态"
//	@Param			page	path		int	false	"起始页"
//	@Success		200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/tags [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	data := make(map[string]interface{})

	data["lists"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// ShowAccount godoc
//
//	@Summary		新增
//	@Description	新增标签
//	@Tags			标签
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"标签名"
//	@Param			state	path		int	true	"状态"
//	@Param			created_by	path		string	true	"创建人"
//	@Success		200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/tags [post]
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	//校验
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0和1")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			// log.Printf("err.key是:%s,err.message是:%s", err.Key, err.Message)
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// ShowAccount godoc
//
//	@Summary		编辑
//	@Description	编辑标签
//	@Tags			标签
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"标签名"
//	@Param			state	path		int	true	"状态"
//	@Param			modified_by	path		string	true	"修改人"
//	@Success		200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/tags/:id [put]
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")

	modified_by := c.Query("modified_by")

	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("id不能为空")
	valid.Required(modified_by, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modified_by, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modified_by

			if state != -1 {
				data["state"] = state
			}
			if name != "" {
				data["name"] = name
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			// log.Printf("err.key是:%s,err.message是:%s", err.Key, err.Message)
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// ShowAccount godoc
//
//	@Summary		删除
//	@Description	删除标签
//	@Tags			标签
//	@Accept			json
//	@Produce		json
//	@Success		200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/tags/:id [delete]
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于0")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			// log.Printf("err.key是:%s,err.message是:%s", err.Key, err.Message)
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
