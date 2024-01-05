package db

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

type Database struct {
	Client *gorm.DB
}

// TODO: retrieve search criteria for optimized search

func NewDatabase() (*Database, error) {
	// Postgres setup
	connectionConfigs := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_TABLE"), os.Getenv("DB_PASSWORD"), os.Getenv("SSL_MODE"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionConfigs,
		PreferSimpleProtocol: true,
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		return nil, err
	}

	// Redis setup
	return &Database{
		Client: db,
	}, nil
}

// PingPostgres pings the Postgres database
func (d *Database) PingPostgres(ctx context.Context) error {
	// Retrieve the underlying SQL DB from the Gorm DB
	sqlDB, err := d.Client.DB()

	// Ping the database
	err = sqlDB.PingContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
