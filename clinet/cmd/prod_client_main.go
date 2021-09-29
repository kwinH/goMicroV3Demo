package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	Wrappers2 "goMicroCli/client/Wrappers"
	"goMicroCli/client/routes"
	pd "goMicroClient/proto"
)

func main() {

	registry := etcd.NewRegistry(registry.Addrs("http://127.0.0.1:2379"))

	//MyS := micro.NewService(
	//	micro.Name("prodService.client"),
	//	micro.WrapClient(func(c client.Client) client.Client {
	//		return &Wrappers.LogWrapper{Client: c}
	//	}),
	//	micro.WrapClient(func(c client.Client) client.Client {
	//		return &Wrappers.HystrixWrapper{Client: c}
	//	}),
	//	micro.Selector(selector.NewSelector(
	//		selector.Registry(registry),
	//		selector.SetStrategy(selector.Random),
	//	)),
	//)
	//MyClient := MyS.Client()
	MyClient := client.NewClient(
		client.Selector(selector.NewSelector(
			selector.Registry(registry),
			selector.SetStrategy(selector.Random),
		)),
		client.Wrap(func(c client.Client) client.Client {
			return &Wrappers2.LogWrapper{Client: c}
		}),
		client.Wrap(func(c client.Client) client.Client {
			return &Wrappers2.HystrixWrapper{Client: c}
		}),
	)

	prodS := pd.NewProdService("prodService", MyClient)

	ginRouter := routes.NewGinRouter(prodS)
	// Create service
	//httpServer := web.NewService(
	//	web.Name("prodService"),
	//	//web.Address(":8001"),
	//	web.Handler(ginRouter),
	//	//web.Registry(registry),
	//	web.Metadata(map[string]string{"protocol": "http"}),
	//)

	//httpServer.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello world"))
	//})

	//httpServer.Init()
	//httpServer.Run()

	ginRouter.Run(":9001")
	//go run prod_main.go --server_address :8001
}
