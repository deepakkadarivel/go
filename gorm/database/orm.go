package database

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

var Ormdb *gorm.DB

// TODO step 1 : open db
func Opendb(dialect string, dbName string)  {
	db, error := gorm.Open(dialect, dbName)
	if error != nil {
		fmt.Println(error)
		panic("failed to connect to database.")
	}

	Ormdb = db

	defer Ormdb.Close()
}

// TODO step 2 : Migrate the scheme
func Migratedb(values ...interface{}) {
	Ormdb.AutoMigrate(values)
}
// TODO step 3 : Create
// TODO step 4 : Read
// TODO step 5 : Update
// TODO step 6 : Delete


