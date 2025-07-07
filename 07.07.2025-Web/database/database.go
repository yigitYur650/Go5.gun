// Package database provides database connection and initialization logic.
package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB // db değişkenini evrensel bir değişkene atadık

func Connect() {
	db, err := gorm.Open(sqlite.Open("Databesm.db")) // gorma burda diyoruz ki sen git sqlite dosyasını databasm de aç burda open fonksiyonu bize 2 şey return ediyor biri db biride err
	if err != nil {
		fmt.Println("CONNETC ERROR")
	}
	DB = db
	fmt.Println("CONNETC SUCCESFULY")
}
