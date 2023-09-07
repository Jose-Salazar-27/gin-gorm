package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (tx *gorm.DB) {
	dsn := "user=postgres password=bricsport123$ host=db.hxgebjshuvcphduvezcc.supabase.co port=5432 dbname=postgres"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB CONNECTION WAS FAILED. SEE ERROR: " + err.Error())
	}

	fmt.Println("\033[32mSuccess: Db connection went well!\033[0m")
	return db
}
