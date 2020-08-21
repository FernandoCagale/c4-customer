package notify

import (
	"github.com/FernandoCagale/c4-customer/internal/notify/grpc"
	"github.com/FernandoCagale/c4-customer/internal/notify/http"
	"github.com/google/wire"
)

var SetGRPC = wire.NewSet(grpc.NewNotifyGRPC,
	wire.Bind(new(Notify), new(*grpc.NotifyGRPC)), //grpc
)

var SetHTTP = wire.NewSet(http.NewNotifyHTTP,
	wire.Bind(new(Notify), new(*http.NotifyHTTP)), // http
)
