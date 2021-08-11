package service

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var Provides = wire.NewSet(
	wire.Struct(new(Dependency), "*"),
	wire.Struct(new(Services), "*"),
	NewAdmin,
)

type Services struct {
	*Dependency
	Admin *Admin
}

type Dependency struct {
	Db    *gorm.DB
	Redis *redis.Client
}