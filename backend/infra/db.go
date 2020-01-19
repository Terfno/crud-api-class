package infra

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(127.0.0.1:63306)"
	DBNAME := "crud_db"

	CONNECTION := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECTION)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
