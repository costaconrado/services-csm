// Package postgres implements postgres connection.
package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/costaconrado/services-csm/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB gorm.DB
}

// New -.
func New(cfg config.PG) (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo",
		cfg.Host, cfg.User, cfg.Pass, cfg.Database, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	connTimeout := time.Duration(cfg.ConnSecondsTimeout) * time.Second
	connAttempts := cfg.MaxConnAttempts

	for connAttempts > 0 && err == nil {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", connAttempts)

		time.Sleep(connTimeout)

		connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - not able to get sql connection: %w", err)
	}
	sqlDB.SetMaxIdleConns(cfg.PoolMax)

	return &Postgres{DB: *db}, nil
}

// Close -.
func (postgres *Postgres) Close() {
	sqlDB, err := postgres.DB.DB()
	if err == nil {
		sqlDB.Close()
	}
}
