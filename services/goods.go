package services

import (
	"fmt"
	"github.com/pborman/uuid"
	goodsServoce "github.com/ttooch/proto/goods"
	"github.com/ttooch/proto/pubsub"
	"time"
)

func GoodsDetail() {
	req := rpcClient.NewRequest("Goods.Detail", goodsRpc, &goodsServoce.DetailRequest{
		GoodsId: 1,
	})

	rsp := &goodsServoce.DetailResponse{}

	// Call service
	if err := rpcClient.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println(rsp)
}

func PubGoods() {

	msg := rpcClient.NewPublication(goodsTopic, &pubsub.Event{
		Id:        uuid.NewUUID().String(),
		Timestamp: time.Now().Unix(),
		Message:   fmt.Sprintf("Messaging you all day on %s", "1111"),
	})

	if err := rpcClient.Publish(ctx, msg); err != nil {
		fmt.Println("pub err: ", err)
		return
	}

	fmt.Printf("Published: %v\n", msg)
}
