// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"sample-order/api/v1/item"
	item3 "sample-order/core/item"
	"sample-order/libs"
	item2 "sample-order/repository/item"
)

// Injectors from wire.go:

//InitializeItemControllerV1 Initialize item service
func InitializeItemControllerV1(dbCon *libs.DatabaseConnection) *item.Controller {
	dataRepository := item2.DataRepositoryFactory(dbCon)
	serviceImpl := item3.NewServiceImpl(dataRepository)
	controller := item.NewController(serviceImpl)
	return controller
}