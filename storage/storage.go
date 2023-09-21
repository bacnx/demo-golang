package storage

import (
	"go-gin/model"

	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *storage {
	return &storage{db: db}
}

type IStorage interface {
	CreateItem(item *model.Item) error
	GetItems() ([]model.Item, error)
	UpdateItem(id int) error
	DeleteItem(id int) (*model.Item, error)
}
