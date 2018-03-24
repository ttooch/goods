package handler

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/micro/go-log"
	//"github.com/ttooch/goods/service"
	goods "github.com/ttooch/proto/goods"
	//"github.com/ttooch/goods/model"
	//"os/user"
	//"github.com/ttooch/goods/model"
)

type Goods struct{}

func (e *Goods) Detail(ctx context.Context, req *goods.Request, rsp *goods.Response) error {
	//log.Log("Received Goods.Detail request")
	//ads := model.GetAdsForPage(1, 20)

	//rpcAds := make([]*goods.Ads, 0, 20)

	//DeepCopy(ads, &rpcAds)

	//rsp.Data = rpcAds

	rsp.Msg = "goodsDetail"

	//personDB := make(map[string]goods.Response)

	//var person = goods.Response{123, "msg", "bbb"}
	//ads := goods.Ads{1,"2","2","2","2",3,3}
	ad := goods.Ads{1,"2","2","2","2",3,3}

	rsp.Status = 1

	ads  := make([]*goods.Ads, 0)

	ads = append(ads, &ad)

	fmt.Println(ad)

	fmt.Println(&ad)

	fmt.Println(*&ad)

	rsp.Data = ads

	//fmt.Println("111")
	//
	//go service.PubGoods()
	//
	//fmt.Println("PubGoods1")
	//
	//go service.PubGoods()
	//
	//fmt.Println("PubGoods2")

	//var m1 map[string]int32
	//
	//m1 = make(map[string]int32)
	//
	//m1["a"] = 12312
	//m1["b"] = 565464

	fmt.Println(ads)

	return nil
}
