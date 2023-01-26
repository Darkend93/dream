package datastore

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDB(dbSource *string) *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(*dbSource),
		&gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}},
	)
	if err != nil {
		panic(err)
	}

	return db
}
