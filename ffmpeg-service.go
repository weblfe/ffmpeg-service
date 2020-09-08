package main

import (
		"flag"
		"github.com/tal-tech/go-zero/core/logx"
		"github.com/weblfe/ffmpeg-service/internal/config"
		"github.com/weblfe/ffmpeg-service/internal/handler"
		"github.com/weblfe/ffmpeg-service/internal/svc"

		"github.com/tal-tech/go-zero/core/conf"
		"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/ffmpeg-service-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	logx.Info(c.Name," listen at ",c.Port)
	server.Start()
}
