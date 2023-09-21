package storage

import (
	"go-gin/model"
)

func (s *storage) CreateItem(item *model.Item) error {
	if err := s.db.Create(item).Error; err != nil {
		return err
	}

	return nil
}

func (s *storage) GetItems() ([]model.Item, error) {
	var items []model.Item
	if err := s.db.Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (s *storage) UpdateItem(id int) error {
	var item *model.Item
	if err := s.db.Where("id = ?", id).First(&item).Error; err != nil {
		return err
	}
	s.db.Save(item)
	return nil
}

func (s *storage) DeleteItem(id int) (*model.Item, error) {
	var item *model.Item
	if err := s.db.Where("id = ?", id).Delete(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}
