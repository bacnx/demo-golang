package config

type Config struct {
  DBHost  string
  DBPort  string
  DBUser  string
  DBPassword  string
  DBName  string
  Port    string
}

var cfg Config

func SetEnv() {
  cfg = Config{
    DBHost: "localhost",
    DBPort: "5432",
    DBUser: "nxbac",
    DBPassword: "123456",
    DBName: "gin",
    Port: "8000",
  }
}

func GetEnv() Config {
  return cfg
}
