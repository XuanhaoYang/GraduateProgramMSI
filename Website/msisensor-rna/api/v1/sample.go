package v1

import (
	"msisensor-rna/model"
	"msisensor-rna/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 新增样本
func AddSample(ctx *gin.Context) {

	var data model.Sample
	_ = ctx.ShouldBind(&data)
	data.Uid = int(uid)
	code = errmsg.SUCCESS
	if code == errmsg.ERROR_SAMPNAME_USED {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		model.CreateSample(&data)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})

	}
}

// 查询样本列表
func GetSamples(ctx *gin.Context) {
	// pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	// pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	// switch {
	// case pageSize >= 100:
	// 	pageSize = 100
	// case pageSize <= 0:
	// 	pageSize = 10
	// }

	// if pageNum == 0 {
	// 	pageNum = 1
	// }

	data, code := model.GetSamples(int(uid))

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		// "total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除样本
func DeleteSample(ctx *gin.Context) {
	sid, _ := strconv.Atoi(ctx.Param("sid"))
	code = model.DeleteSample(sid)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func EditSample(ctx *gin.Context) {
	var data model.Sample
	sid, _ := strconv.Atoi(ctx.Param("sid"))
	_ = ctx.ShouldBindJSON(&data)
	code := model.EditSample(sid, &data)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
