package repo

import (
	"go-gin/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

type IRepo interface {
	GetItems(*gin.Context) ([]model.Item, error)
}
