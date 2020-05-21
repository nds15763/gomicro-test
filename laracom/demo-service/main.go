package main

import (
	"context"
	"log"

	pb "laracom/demo-service/proto/demo"

	"github.com/micro/go-micro/v2"
)

type DemoServiceHandler struct {
}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp.Text = "你好, " + req.Name
	return nil
}

func (s *DemoServiceHandler) SayHelloByUserId(ctx context.Context, req *pb.HelloRequest, rsp *pb.DemoResponse) error {
	return nil
}

func main() {
	// 注册服务名必须和 demo.proto 中的 package 声明一致
	service := micro.NewService(micro.Name("laracom.demo.service"))
	service.Init()

	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
