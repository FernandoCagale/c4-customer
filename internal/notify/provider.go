package notify

import (
	"github.com/FernandoCagale/c4-customer/internal/notify/grpc"
	"github.com/google/wire"
)

var Set = wire.NewSet(grpc.NewNotifyGRPC, wire.Bind(new(Notify), new(*grpc.NotifyGRPC)))
