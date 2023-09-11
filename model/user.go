package model

type User struct {
  Id        int64   `json:"id"`
  UserName  string  `json:"user_name"`
  Password  string  `json:"password"`
}

func (User) TableName() string {
  return "users"
}
