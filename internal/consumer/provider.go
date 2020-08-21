package consumer

import "github.com/google/wire"

var Set = wire.NewSet(New, wire.Bind(new(Consumer), new(*ConsumerKafka)))
