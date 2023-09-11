package repo
import (
  "go-gin/model"
)

func (r *Repo) GetUsers() ([]model.User, error) {
  var users []model.User
  if err := r.db.Find(&users).Error; err != nil {
    return []model.User{}, err
  }
  return users, nil
}
