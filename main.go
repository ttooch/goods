package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	_ "github.com/micro/go-plugins/broker/nsq"
	_ "github.com/micro/go-plugins/registry/etcd"
	"github.com/ttooch/goods/handlers"
	myService "github.com/ttooch/goods/services"
	"github.com/ttooch/goods/subscribers"
	goodsService "github.com/ttooch/proto/goods"
	"log"
	"time"
	"github.com/ttooch/goods/models"
)

var (
	topic = "topic.go.micro.api.v1.goods"
)

func main() {

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.v1.goods"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.AfterStop(func() error {
			return models.Engine.Close()
		}),
	)

	// di service
	myService.InitService(service)

	// Register Handler
	goodsService.RegisterGoodsHandler(service.Server(), new(handlers.Goods))

	micro.RegisterSubscriber(topic, service.Server(), new(subscribers.Goods), server.SubscriberQueue("consumer"))

	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
