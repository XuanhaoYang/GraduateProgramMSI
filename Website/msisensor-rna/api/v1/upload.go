package v1

import (
	"msisensor-rna/model"
	"msisensor-rna/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{"err": err})
		return
	}

	dst := "./data/" + file.Filename
	// 上传文件至指定的完整文件路径
	code = model.UploadFile(c, file, dst)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"path":    dst,
	})

}
