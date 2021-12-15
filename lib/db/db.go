package db

import (
	"database/sql"

	"github.com/go-mservice-bench/lib/account"
	"github.com/go-mservice-bench/lib/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Client  *sql.DB
	Account account.Model
	Config  *config.Config
}

func Init(config *config.Config) (*DB, error) {
	database, err := gorm.Open(sqlite.Open(config.DbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	client, err := database.DB()
	if err != nil {
		return nil, err
	}

	database.AutoMigrate(&account.Account{})

	return &DB{
		Client:  client,
		Account: account.NewModel(database),
		Config:  config,
	}, nil
}
