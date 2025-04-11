package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"lowcode/internal/config"
	"lowcode/internal/errcode"
	"lowcode/internal/handler"
	"lowcode/internal/svc"
	_ "lowcode/internal/validate"
)

var configFile = flag.String("f", "etc/lowcode-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.DisableStat()
	logx.MustSetup(c.Log)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()
	// File Handler
	// helper.FileRegisterHandlers(server, "/api/files", "/opt/framework/post", false)

	// API Handler
	handler.RegisterHandlers(server, ctx)
	// Swagger Doc Handler
	handler.RegisterSwaggerHandlers(server, ctx)
	// Error Handler
	// logx error's msg and error's stack to log file
	httpx.SetErrorHandler(errcode.HandleError)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
