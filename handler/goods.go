package handler

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-log"
	"github.com/ttooch/goods/service"
	goods "github.com/ttooch/proto/goods"
	"github.com/ttooch/goods/model"
)

type Goods struct{}

func (e *Goods) Detail(ctx context.Context, req *goods.Request, rsp *goods.Response) error {
	log.Log("Received Goods.Detail request")
	ads := model.GetAdsForPage(1, 20)

	rpcAds := make([]*goods.Ads, 0, 20)

	DeepCopy(ads, &rpcAds)

	rsp.Data = rpcAds

	rsp.Msg = "goodsDetail"

	fmt.Println("111")

	go service.PubGoods()

	fmt.Println("PubGoods1")

	go service.PubGoods()

	fmt.Println("PubGoods2")

	fmt.Println("2222")
	return nil
}
