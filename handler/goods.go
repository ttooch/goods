package handler

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-log"
	goods "github.com/ttooch/proto/goods"
	"github.com/ttooch/goods/model"
)

type Goods struct{}

func (e *Goods) Detail(ctx context.Context, req *goods.DetailRequest, rsp *goods.DetailResponse) error {
	log.Log("Received Goods.Detail request")

	//goodsModel := new(model.Goods)

	//goodsModel.GetGoodsByPk()

	goodsModel := model.GetGoodsByPk()

	fmt.Println(*goodsModel)

	rpcGoods := new(goods.Good)

	DeepCopy(goodsModel,rpcGoods)

 	rsp.Data = rpcGoods

	//ads := model.GetAdsForPage(1, 20)
	//
	//rpcAds := make([]*goods.Ads, 0, 20)
	//
	//DeepCopy(ads, &rpcAds)
	//
	//rsp.Data = rpcAds
	//
	//rsp.Msg = "goodsDetail"
	//
	//fmt.Println("111")
	//
	//go service.PubGoods()
	//
	//fmt.Println("PubGoods1")
	//
	//go service.PubGoods()
	//
	//fmt.Println("PubGoods2")
	//
	//fmt.Println("2222")
	return nil
}
func (e *Goods) List(ctx context.Context, req *goods.ListRequest, rsp *goods.ListResponse) error {

	goodsList := model.GetGoodsForPage(1,20)

	rpcGoodsList := make([]*goods.Good, 0, 20)

	DeepCopy(goodsList,&rpcGoodsList)

	rsp.Data = rpcGoodsList

	return nil

}
