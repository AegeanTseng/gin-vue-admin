package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
)

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
func CreateApi(c *gin.Context) {
	var api model.SysApi
	_ = c.ShouldBindJSON(&api)
	err := api.CreateApi()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "创建成功", c)
	}
}

// @Tags SysApi
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysModel.SysApi true "删除api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/deleteApi [post]
func DeleteApi(c *gin.Context) {
	var a model.SysApi
	_ = c.ShouldBindJSON(&a)
	err := a.DeleteApi()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "删除成功", c)
	}
}

type AuthAndPathIn struct {
	AuthorityId string `json:"authorityId"`
	ApiIds      []uint `json:"apiIds"`
}

//条件搜索后端看此api

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PageInfo true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func GetApiList(c *gin.Context) {
	// 此结构体仅本方法使用
	type searchParams struct {
		model.SysApi
		model.PageInfo
	}
	var sp searchParams
	_ = c.ShouldBindJSON(&sp)
	err, list, total := sp.SysApi.GetInfoList(sp.PageInfo)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{
			"list":     list,
			"total":    total,
			"page":     sp.PageInfo.Page,
			"pageSize": sp.PageInfo.PageSize,
		}, "删除成功", c)
	}
}

// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func GetApiById(c *gin.Context) {
	var idInfo GetById
	_ = c.ShouldBindJSON(&idInfo)
	err, api := new(model.SysApi).GetApiById(idInfo.Id)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{
			"api": api,
		}, "获取数据成功", c)

	}
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/updateApi [post]
func UpdateApi(c *gin.Context) {
	var api model.SysApi
	_ = c.ShouldBindJSON(&api)
	err := api.UpdateApi()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("修改数据失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "修改数据成功", c)
	}
}

// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func GetAllApis(c *gin.Context) {
	err, apis := new(model.SysApi).GetAllApis()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{
			"apis": apis,
		}, "获取数据成功", c)
	}
}
