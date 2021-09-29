package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"goMicroCli/routes"
	Wrappers2 "goMicroCli/wrappers"
	pd "goMicroSrv/proto"
)

// go run main.go
func main() {
	registry := etcd.NewRegistry(registry.Addrs("http://127.0.0.1:2379"))

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

	ginRouter.Run(":8809")
}
