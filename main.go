package main

import (
	//"github.com/micro/go-micro"
	//"github.com/micro/go-micro/server"
	_ "github.com/micro/go-plugins/broker/nsq"
	_ "github.com/micro/go-plugins/registry/etcd"
	//"github.com/ttooch/goods/handlers"
	//"github.com/ttooch/goods/models"
	//myService "github.com/ttooch/goods/services"
	//"github.com/ttooch/goods/subscribers"
	//goodsService "github.com/ttooch/proto/goods"
	//"log"
	//"time"
	//"github.com/ttooch/goods/elastics"
	"github.com/ttooch/goods/elastics"
	//"fmt"
	//"encoding/json"
	//"github.com/ttooch/goods/models"
	//"gopkg.in/olivere/elastic.v5"
	//"fmt"
	"fmt"
)

var (
	topic = "topic.go.micro.api.v1.goods"
)

func main() {

	//goodsList := elastics.GetGoodsList()
	//
	//b,_ := json.Marshal(goodsList)
	//
	//fmt.Println(string(b))

	//elastics.UpdateGoods("AWJr9TvbpO3ECuIBreLn",&elastics.Goods{
	//	GoodsBrief:"1111",
	//})

	goods, _ := elastics.GetGoodsByID("AWJr9TvbpO3ECuIBreLn")

	fmt.Printf("%+v", goods.Snapshot)

	goods = new(elastics.Goods)

	goods.BarCode = "3333"

	elastics.UpdateGoodsById(goods)

	//query := elastic.NewBoolQuery()
	//
	//query.Must(elastic.NewTermQuery("goods_id", 1))
	//
	//goods, _ := elastics.GetGoodsList(query,map[string]bool{"goods_id":true})
	//
	//fmt.Println(goods)
	//
	//b,_ := json.Marshal(goods)
	//
	//fmt.Println(string(b))
}
