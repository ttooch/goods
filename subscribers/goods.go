package subscribers

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro/metadata"
	"github.com/ttooch/proto/pubsub"
	"golang.org/x/net/context"
)

type Goods struct{}

func (s *Goods) ModifyStock(ctx context.Context, event *pubsub.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}
