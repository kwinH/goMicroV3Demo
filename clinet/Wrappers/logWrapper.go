package Wrappers

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/metadata"
)

type LogWrapper struct {
	client.Client
}

func (l *LogWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx:%v service:%s method:%s\n", md, req.Service(), req.Endpoint())

	return l.Client.Call(ctx, req, rsp)
}
