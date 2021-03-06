package routes

import (
	"github.com/gin-gonic/gin"
	"goMicroCli/handler"
	middlewares2 "goMicroCli/middlewares"
	pd "goMicroSrv/proto"
)

func NewGinRouter(prodService pd.ProdService) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middlewares2.InitMiddleware(prodService))
	v1Group := ginRouter.Group("v1")

	{

		v1Group.Handle("POST", "/prods", handler.GetProdsList)

		v1Group.Handle("POST", "/prods1", func(context *gin.Context) {
			//context.Param("size")
			//context.JSON(200,
			//	gin.H{
			//		"data": ProdService.NewProdList(5),
			//	},
			//
			//)
		})

	}

	return ginRouter
}
