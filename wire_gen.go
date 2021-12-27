// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"api/app"
	"api/app/index"
	"api/app/pages"
	"api/app/roles"
	"api/bootstrap"
	"api/common"
	"github.com/gin-gonic/gin"
	"github.com/weplanx/go/api"
)

// Injectors from wire.go:

func App(value *common.Values) (*gin.Engine, error) {
	client, err := bootstrap.UseMongoDB(value)
	if err != nil {
		return nil, err
	}
	database := bootstrap.UseDatabase(client, value)
	inject := &common.Inject{
		Values:      value,
		MongoClient: client,
		Db:          database,
	}
	service := &index.Service{
		Inject: inject,
	}
	controller := &index.Controller{
		Service: service,
	}
	apiService := &api.Service{
		Db: database,
	}
	apiController := &api.Controller{
		Service: apiService,
	}
	pagesService := &pages.Service{
		Inject: inject,
	}
	pagesController := &pages.Controller{
		Service: pagesService,
	}
	rolesService := &roles.Service{
		Inject: inject,
	}
	rolesController := &roles.Controller{
		Service: rolesService,
	}
	engine := app.New(value, controller, apiController, pagesController, rolesController)
	return engine, nil
}
