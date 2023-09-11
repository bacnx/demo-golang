package repo
import (
  "go-gin/model"
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
  GetUsers(name string) ([]model.User, error)
}
