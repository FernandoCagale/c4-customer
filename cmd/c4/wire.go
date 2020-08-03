//+build wireinject

package main

import (
	"github.com/FernandoCagale/c4-customer/api/routers"
	"github.com/FernandoCagale/c4-customer/event"
	"github.com/FernandoCagale/c4-customer/internal/datastore"
	notifyGRPC "github.com/FernandoCagale/c4-customer/internal/notify/grpc"
	"github.com/FernandoCagale/c4-customer/pkg"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func SetupApplication(*gorm.DB, *grpc.ClientConn) (*routers.SystemRoutes, error) {
	wire.Build(pkg.Container)
	return nil, nil
}

func SetupPostgres() (*gorm.DB, error) {
	wire.Build(datastore.Set)
	return nil, nil
}

func SetupEvents(*gorm.DB, *grpc.ClientConn) (*event.CustomerEvent, error) {
	wire.Build(event.Set)
	return nil, nil
}

func SetupClientGRPC() (*grpc.ClientConn, error) {
	wire.Build(notifyGRPC.Set)
	return nil, nil
}
