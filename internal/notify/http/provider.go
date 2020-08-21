package http

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(NewNotifyHTTP)
