package service

import (
	"context"
	"go-gin/model"
)

type IItem interface {
	CreateItem(item *model.Item) error
	GetItems() ([]model.Item, error)
	UpdateItem(item *model.Item) error
	DeleteItem(id int) (*model.Item, error)
}

type Item struct {
	storage IItem
}

func NewItem(storage IItem) *Item {
	return &Item{
		storage: storage,
	}
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

func (i *Item) UpdateItem(ctx context.Context, item *model.Item) error {
	if err := i.storage.UpdateItem(item); err != nil {
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
