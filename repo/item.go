package repo

import (
	"go-gin/model"

	"github.com/gin-gonic/gin"
)

func (r *Repo) GetItems(c *gin.Context) ([]model.Item, error) {
	var items []model.Item
	if err := r.db.Find(&items).Error; err != nil {
		return []model.Item{}, err
	}
	return items, nil
}
