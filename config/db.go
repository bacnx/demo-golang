package config
import (
  "fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "log"
)

var dbDefault *gorm.DB

func GetDB() *gorm.DB {
  if dbDefault == nil {
    return initDB()
  }

  return dbDefault
}

func initDB() *gorm.DB {
  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s",
    cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName,
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
  })
  if err != nil {
    log.Fatalf("Failed to connect to PostgreSQL: %v", err)
  }

  return db
}
