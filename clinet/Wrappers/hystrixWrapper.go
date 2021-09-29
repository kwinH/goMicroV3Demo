package wrappers

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/v3/client"
	"goMicroSrv/handler"
	pd "goMicroSrv/proto"
	"strconv"
)

type HystrixWrapper struct {
	client.Client
}

func defaultData(rsp interface{}) {
	models := make([]*pd.ProdModel, 0)
	var i int32
	for i = 0; i < 2; i++ {
		models = append(models, handler.NewProd(30+i, "prodname"+strconv.Itoa(int(30+i))))
	}

	switch rsp.(type) {
	case *pd.ProdListResponse:
		res := rsp.(*pd.ProdListResponse)
		res.Data = models
	}

	return
}

func (l *HystrixWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdNmae := req.Service() + "." + req.Endpoint()

	//熔断 第一步：配置config
	configA := hystrix.CommandConfig{
		Timeout: 1000,
		//2次内50%出错则直接走熔断，5秒后再请求真实接口
		RequestVolumeThreshold: 2,
		ErrorPercentThreshold:  50,
		SleepWindow:            5000,
	}

	//熔断 第二步：配置command
	hystrix.ConfigureCommand("getprods", configA)

	//熔断 第三步：配置command
	return hystrix.Do(cmdNmae, func() error {
		return l.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		fmt.Printf("hystrix err:%s\n", err.Error())
		defaultData(rsp)
		return nil
	})

}
