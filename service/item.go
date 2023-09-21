package service

import (
	"context"
	"go-gin/model"
	"go-gin/storage"
)

type Item struct {
	storage storage.IStorage
}

func NewItem(storage storage.IStorage) *Item {
	return &Item{
		storage: storage,
	}
}

type IItem interface {
	CreateItem(ctx context.Context, item *model.Item) error
	GetItems(ctx context.Context) ([]model.Item, error)
	UpdateItem(ctx context.Context, id int) error
	DeleteItem(ctx context.Context, id int) (*model.Item, error)
}

func (i *Item) CreateItem(cxt context.Context, item *model.Item) error {
	err := i.storage.CreateItem(item)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) GetItems(ctx context.Context) ([]model.Item, error) {
	items, err := i.storage.GetItems()
	if err != nil {
		return items, err
	}
	return items, nil
}

func (i *Item) UpdateItem(ctx context.Context, id int) error {
	if err := i.storage.UpdateItem(id); err != nil {
		return err
	}
	return nil
}

func (i *Item) DeleteItem(ctx context.Context, id int) (*model.Item, error) {
	item, err := i.storage.DeleteItem(id)
	if err != nil {
		return item, err
	}
	return item, nil
}
