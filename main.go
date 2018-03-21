package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"time"
	"github.com/ttooch/goods/handler"
	myService "github.com/ttooch/goods/service"
	"github.com/ttooch/goods/subscriber"
	goods "github.com/ttooch/proto/goods"
	_ "github.com/micro/go-plugins/broker/nsq"
	_ "github.com/micro/go-plugins/registry/etcd"
)

var (
	topic = "topic.go.micro.srv.goods"
)

func main() {

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.goods"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	// di service
	myService.InitService(service)

	// Register Handler
	goods.RegisterGoodsHandler(service.Server(), new(handler.Goods))

	micro.RegisterSubscriber(topic, service.Server(), new(subscriber.Goods), server.SubscriberQueue("consumer"))

	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
