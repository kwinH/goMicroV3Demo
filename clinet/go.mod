module goMicroCli

go 1.16

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/asim/go-micro/plugins/registry/etcd/v3 v3.0.0-20210904061721-270d910b7328
	github.com/asim/go-micro/v3 v3.5.2
	github.com/gin-gonic/gin v1.7.4
	goMicroSrv v1.0.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
//replace google.golang.org/grpc => google.golang.org/grpc v1.38.0
replace goMicroSrv => github.com/kwinH/goMicroV3Demo/service v0.0.0-20210929012513-35568929b6d6 // indirect
