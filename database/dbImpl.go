package database

import (
	"adivii/forum-discussion/config"
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	DbInstance *PostgresDatabase
)

func NewPostgresDatabase(conf *config.Config) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d",
			conf.Database.Host,
			conf.Database.User,
			conf.Database.Password,
			conf.Database.DBName,
			conf.Database.Port)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect to database")
		}

		DbInstance = &PostgresDatabase{Db: db}
	})

	return DbInstance
}

// GetDb implements Database.
func (m *PostgresDatabase) GetDb() *gorm.DB {
	return DbInstance.Db
}
