package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"gomicro_warmhome/PostReg/handler"
	"gomicro_warmhome/PostReg/subscriber"

	example "gomicro_warmhome/PostReg/proto/example"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostReg"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostReg", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostReg", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
