package router

import (
	"go-gin/config"
	"go-gin/handler"
	"go-gin/service"
	"go-gin/storage"
)

type Service struct {
	*config.App
}

func NewService() *Service {
	s := Service{
		config.NewApp(),
	}

	db := s.GetDB()
	storage := storage.NewStorage(db)

	itemService := service.NewItem(storage)
	item := handler.NewItem(itemService)

	router := s.Router
	v1 := router.Group("/v1")

	v1.POST("/items", item.CreateItem)
	v1.GET("/items", item.GetItems)
	v1.PUT("/items/:id", item.UpdateItem)
	v1.DELETE("/items/:id", item.DeleteItem)

	return &s
}
