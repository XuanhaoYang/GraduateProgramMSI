package model

import (
	"mime/multipart"
	"msisensor-rna/utils/errmsg"

	"github.com/gin-gonic/gin"
)

// 上传样本到指定路径
func UploadFile(c *gin.Context, file *multipart.FileHeader, dst string) (code int) {

	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		return errmsg.ERROR_UPLOAD_WRONG
	}

	return errmsg.SUCCESS
}
