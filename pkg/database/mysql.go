package database

import (
	"fmt"

	"github.com/humamalamin/test-case-dating/pkg/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Gorm interface {
	Connect() (*gorm.DB, error)
}

type Options struct {
	master  string
	maxOpen int
	maxIdle int
}

func NewGorm(cfg *config.Config) Gorm {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabasHost, cfg.DatabasePort, cfg.DatabaseName)
	opt := new(Options)
	opt.master = conn
	opt.maxIdle = cfg.DatabaseMaxIdleConnections
	opt.maxOpen = cfg.DatabaseMaxOpenConnections

	return opt
}

func (o *Options) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(o.master), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Error().Err(err).Msg("[Connect-1] failed to connect to database " + o.master)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[Connect-2] failed to connect to database " + o.master)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(o.maxOpen)
	sqlDB.SetMaxIdleConns(o.maxIdle)

	return db, nil
}
