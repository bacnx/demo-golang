package service
import (
  "go-gin/model"
  "go-gin/repo"
)

type User struct {
  repo repo.IRepo
}

func NewUser(repo repo.IRepo) *User {
  return &User{
    repo: repo,
  }
}

type IUser interface {
  GetUsers() (users []model.User, err error)
}

func (s *User) GetUsers() (users []model.User, err error) {
  users, err = s.repo.GetUsers()
  if err != nil {
    return users, err
  }
  return users, nil
}
