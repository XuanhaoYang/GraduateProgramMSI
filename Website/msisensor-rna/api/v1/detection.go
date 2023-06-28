package v1

import (
	"msisensor-rna/model"
	"msisensor-rna/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BulkDetection(ctx *gin.Context) {
	fileName := ctx.Query("filename")

	res := model.BulkDetection(fileName)
	if res.Code == 1 {
		code = errmsg.SUCCESS
	} else {
		code = errmsg.ERROR
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":     code,
		"weight":     res.Weight,
		"value":      res.Value,
		"score":      res.Score,
		"prediction": res.Pre,
		"message":    res.Message,
	})
}
