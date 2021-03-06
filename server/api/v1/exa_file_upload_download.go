package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/upload [post]
func UploadFile(c *gin.Context) {
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("上传文件失败，%v", err), c)
	} else {
		//文件上传后拿到文件路径
		err, filePath, key := utils.Upload(header, USER_HEADER_BUCKET, USER_HEADER_IMG_PATH)
		if err != nil {
			response.Result(response.ERROR, gin.H{}, fmt.Sprintf("接收返回值失败，%v", err), c)
		} else {
			//修改数据库后得到修改后的user并且返回供前端使用
			var file model.ExaFileUploadAndDownload
			file.Url = filePath
			file.Name = header.Filename
			s := strings.Split(file.Name, ".")
			file.Tag = s[len(s)-1]
			file.Key = key
			if noSave == "0" {
				err = file.Upload()
			}
			if err != nil {
				response.Result(response.ERROR, gin.H{}, fmt.Sprintf("修改数据库链接失败，%v", err), c)
			} else {
				response.Result(response.SUCCESS, gin.H{"file": file}, "上传成功", c)

			}
		}
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body dbModel.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
func DeleteFile(c *gin.Context) {
	var file model.ExaFileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	err, f := file.FindFile()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("删除失败，%v", err), c)
	} else {
		err = utils.DeleteFile(USER_HEADER_BUCKET, f.Key)
		if err != nil {
			response.Result(response.ERROR, gin.H{}, fmt.Sprintf("删除失败，%v", err), c)

		} else {
			err = f.DeleteFile()
			if err != nil {
				response.Result(response.ERROR, gin.H{}, fmt.Sprintf("删除失败，%v", err), c)
			} else {
				response.Result(response.SUCCESS, gin.H{}, "删除成功", c)
			}
		}
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PageInfo true "分页获取文件户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
func GetFileList(c *gin.Context) {
	var pageInfo model.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(model.ExaFileUploadAndDownload).GetInfoList(pageInfo)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取数据成功", c)
	}
}
