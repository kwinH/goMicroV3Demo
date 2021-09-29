package middlewares

import (
	"github.com/gin-gonic/gin"
	pd "goMicroClient/proto"
)

func InitMiddleware(prodService pd.ProdService) gin.HandlerFunc {
	return func(cxt *gin.Context) {
		cxt.Keys=make(map[string]interface{})
		cxt.Keys["prodService"] = prodService
		cxt.Next()
	}
}