package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres dialect
)

var _db *gorm.DB
var _once = sync.Once{}

func connectionStrFromEnvs() string {
	connStr := "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	return fmt.Sprintf(connStr, os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_SSLMODE"))
}

// Initialize connection to database and returns this connection
func Initialize() *gorm.DB {
	_once.Do(func() {
		db, err := gorm.Open("postgres", connectionStrFromEnvs())
		if err != nil {
			panic(err)
		}

		db.LogMode(true) // TODO(khanek) Get value from env
		_db = db
	})
	return _db
}

// DB returns connection to database
func DB() *gorm.DB {
	if _db == nil {
		panic("Database not initialized!")
	}
	return _db
}
