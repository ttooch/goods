package handlers

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ttooch/goods/models"
	goodsService "github.com/ttooch/proto/goods"
)

type Goods struct{}

func (g *Goods) Detail(ctx context.Context, req *goodsService.DetailRequest, rsp *goodsService.DetailResponse) error {

	rpcGoods := new(goodsService.RpcGoods)

	goods, err := models.GetGoodsByID(req.GoodsId)

	DeepCopy(goods, rpcGoods)

	rsp.Data = rpcGoods

	if err != nil {

		rsp.Status = ERROR_STAUTS

		rsp.Msg = err.Error()

	} else {
		rsp.Status = SUCCESS_STAUTS

		rsp.Msg = SUCCESS_MSG
	}

	return nil
}

func (g *Goods) List(ctx context.Context, req *goodsService.ListRequest, rsp *goodsService.ListResponse) error {

	session := models.Engine.Where("goods_id = ?", 1).Desc("goods_id")

	//goods, err := models.GetGoodsListForPage(session, 0, 20)

	goods, err := models.GetGoodsList(session)

	rpcGoods := make([]*goodsService.RpcGoods, len(goods), 10)

	DeepCopy(goods, &rpcGoods)

	rsp.Data = rpcGoods

	if err != nil {

		rsp.Status = ERROR_STAUTS

		rsp.Msg = err.Error()

	} else {
		rsp.Status = SUCCESS_STAUTS

		rsp.Msg = SUCCESS_MSG
	}

	return nil
}

func (g *Goods) Update(ctx context.Context, req *goodsService.UpdateRequest, rsp *goodsService.UpdateResponse) error {

	goods := new(models.Goods)

	DeepCopy(req, goods)

	session := models.Engine.Where("goods_id = ?", req.GoodsId)

	_, err := models.UpdateGoods(session, goods)

	if err != nil {
		rsp.Status = ERROR_STAUTS

		rsp.Msg = err.Error()
	} else {
		rsp.Status = SUCCESS_STAUTS

		rsp.Msg = "修改成功"
	}

	rsp.Data = new(goodsService.EmptyObject)

	return nil
}

func (g *Goods) Add(ctx context.Context, req *goodsService.AddRequest, rsp *goodsService.AddResponse) error {

	goods := new(models.Goods)

	DeepCopy(req, goods)

	err := models.AddGoods(goods)

	if err != nil {
		rsp.Status = ERROR_STAUTS

		rsp.Msg = err.Error()
	} else {
		rsp.Status = SUCCESS_STAUTS

		rsp.Msg = "添加成功"
	}

	rsp.Data = new(goodsService.EmptyObject)

	return nil
}

func (g *Goods) Delete(ctx context.Context, req *goodsService.DeleteRequest, rsp *goodsService.DeleteResponse) error {

	_,err :=models.DelGoodsById(req.GoodsId)

	if err != nil {
		rsp.Status = ERROR_STAUTS

		rsp.Msg = err.Error()
	} else {
		rsp.Status = SUCCESS_STAUTS

		rsp.Msg = "删除成功"
	}

	rsp.Data = new(goodsService.EmptyObject)

	return nil
}
