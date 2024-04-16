package orm

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var connMap sync.Map

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

// Get Data Source Name
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		c.Host,
		c.Username,
		c.Password,
		c.DBName,
		c.Port,
		"Asia/Taipei",
	)
}

// Connect Database
func (c *Config) Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(c.DSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
}

// GetDB GetDB
func GetDB(c *Config) *gorm.DB {
	if db, ok := connMap.Load(c); ok {
		return db.(*gorm.DB)
	}

	db, err := c.Connect()
	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}

	if db, ok := connMap.LoadOrStore(c, db); ok {
		return db.(*gorm.DB)
	}

	return db
}
