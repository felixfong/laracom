package main

import (
	"context"
	"github.com/micro/go-micro"
	pb "github.com/felixfong/laracom/demo-service/proto/demo"
	"log"
)

type DemoService struct {}

func (s *DemoService) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp.Text = "你好， " + req.Name
	return nil
}

func main() {
	service := micro.NewService(micro.Name("laracom.demo.service"))
	service.Init()
	pb.RegisterDemoServiceHandler(service.Server(), &DemoService{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败：%v", err)
	}
}

