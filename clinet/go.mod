module goMicroCli

go 1.16

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/asim/go-micro/plugins/registry/etcd/v3 v3.0.0-20210904061721-270d910b7328
	github.com/asim/go-micro/v3 v3.5.2
	github.com/gin-gonic/gin v1.7.4
	github.com/kr/pretty v0.3.0 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/sys v0.0.0-20210903071746-97244b99971b // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210903162649-d08c68adba83 // indirect
	goMicroSrv v1.0.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
//replace google.golang.org/grpc => google.golang.org/grpc v1.38.0
replace goMicroSrv => github.com/kwinH/go-micro-v3-service-demo v1.0.0 // indirect
