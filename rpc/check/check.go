// Code generated by goctl. DO NOT EDIT!
// Source: check.proto

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/weblfe/ffmpeg-service/rpc/check/internal/config"
	"github.com/weblfe/ffmpeg-service/rpc/check/internal/server"
	"github.com/weblfe/ffmpeg-service/rpc/check/internal/svc"
	check "github.com/weblfe/ffmpeg-service/rpc/check/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rpcx"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/check.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	checkerSrv := server.NewCheckerServer(ctx)

	s, err := rpcx.NewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		check.RegisterCheckerServer(grpcServer, checkerSrv)
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}