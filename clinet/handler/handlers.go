package handler

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"goMicroClient/handler"
	pd "goMicroClient/proto"
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
		context.JSON(200, gin.H{"code":prodRes.Code,"data": prodRes.Data})

		return

		//熔断 第一步：配置config
		configA := hystrix.CommandConfig{
			Timeout: 1000,
		}

		//熔断 第一步：配置command
		hystrix.ConfigureCommand("getprods", configA)

		//熔断 第一步：配置command
		err := hystrix.Do("getprods", func() error {
			prodRes, err = prodS.GetProdsList(context, &prodReq)
			return err
		}, func(err error) error {
			prodRes, err = defaultGetProdsList()
			return err
		})

		if err != nil {
			context.JSON(500, gin.H{"status": err.Error()})
		} else {
			context.JSON(200, gin.H{"data": prodRes.Data})
		}
	}

}
