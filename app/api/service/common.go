package service

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type Dependency struct {
	Db    *gorm.DB
	Redis *redis.Client
}

var Provides = wire.NewSet(
	wire.Struct(new(Dependency), "*"),
	NewIndex,
)
