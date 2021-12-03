package db

import (
	"log"

	"github.com/go-mservice-bench/lib/account"
	"github.com/go-mservice-bench/lib/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Client  *gorm.DB
	Account account.Model
	Config  *config.Config
}

func Init(config *config.Config) (*DB, error) {
	database, err := gorm.Open(sqlite.Open(config.DbName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	database.AutoMigrate(&account.Account{})

	return &DB{
		Client:  database,
		Account: account.NewModel(database),
		Config:  config,
	}, nil
}

func (d *DB) Stop() {
	sqlDb, err := d.Client.DB()
	if err != nil {
		log.Panic(err)
	}
	sqlDb.Close()
}
