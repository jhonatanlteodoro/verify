package mysql_connector

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Get db connection, you must provide a retry limit and waitSeconds to wait before try
// to connect again in case of connection error
func GetConnection(uri string, waitSecondsCaseError int, retry int) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: uri}), &gorm.Config{})
	if err != nil {
		log.Printf("DB CONNECTION ERROR: %v", err)
		if retry > 0 {
			time.Sleep(time.Duration(waitSecondsCaseError) * time.Second)
			log.Println("Trying to connect the database again...")
			retry -= 1
			return GetConnection(uri, waitSecondsCaseError, retry)
		}
		log.Fatal(err)
	}

	return db
}
