package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-gonic/gin"
	"github.com/wa1kman999/goblog-web-svc/global"
	"github.com/wa1kman999/goblog-web-svc/initialize"
	"github.com/wa1kman999/goblog-web-svc/internal/http"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
	"go-micro.dev/v4/web"
)

var (
	serviceName    = "goblog-web-svc"
	serviceVersion = "latest"
)

func main() {
	r := gin.Default()

	// 初始化路由
	http.InitRouter(r)

	srv := web.NewService(
		web.Name(serviceName),
		// 需要注册中心
		web.Registry(consul.NewRegistry(
			registry.Addrs(fmt.Sprintf("%s:%s", global.GoBlogWebConfig.Consul.Host, global.GoBlogWebConfig.Consul.Port)))),
		web.Version(serviceVersion),
		web.Address(global.GoBlogWebConfig.System.Port),
		web.Handler(r),
	)

	// 初始化
	if err := srv.Init(); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// 配置文件初始化
	if err := initialize.ConfigInit(); err != nil {
		panic(err)
	}
}
