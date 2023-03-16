package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB Global DB connection
var DB *gorm.DB

func init() {
	var err error
	if DB, err = gorm.Open("sqlite3", "test.db"); err != nil {
		panic(err)
	}
}
