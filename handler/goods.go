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
	"github.com/ttooch/goods/model"
)

type Goods struct{}

func (e *Goods) Detail(ctx context.Context, req *goods.Request, rsp *goods.Response) error {

	rpcGoods := make([]*goods.Good, 0)

	ads := model.GetGoodsForPage(1, 3)

	DeepCopy(ads, &rpcGoods)

	rsp.Data = rpcGoods

	fmt.Println(ads[0])

	return nil
}
