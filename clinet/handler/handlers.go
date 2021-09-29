package handler

import (
	"github.com/gin-gonic/gin"
	"goMicroSrv/handler"
	pd "goMicroSrv/proto"
	"strconv"
)

func defaultGetProdsList() (rsp *pd.ProdListResponse, err error) {
	models := make([]*pd.ProdModel, 0)
	var i int32
	for i = 0; i < 2; i++ {
		models = append(models, handler.NewProd(30+i, "prodname"+strconv.Itoa(int(30+i))))
	}
	rsp = &pd.ProdListResponse{}
	rsp.Data = models
	return
}

func GetProdsList(context *gin.Context) {
	var prodReq pd.ProdsRequest
	err := context.Bind(&prodReq)

	if err != nil {
		context.JSON(500, gin.H{"status": err.Error()})
	} else {
		prodS := context.Keys["prodService"].(pd.ProdService)
		var prodRes *pd.ProdListResponse

		prodRes, _ = prodS.GetProdsList(context, &prodReq)
		context.JSON(200, gin.H{"code": prodRes.Code, "data": prodRes.Data})

		return
	}
}
