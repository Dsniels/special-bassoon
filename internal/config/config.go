package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	CreateDatabase("fixable")
	url := "server=localhost;port=1433;user id=sa;password=Contraseñña1;database=fixable;encrypt=disable"
	db, err := gorm.Open(sqlserver.Open(url), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db: open %v", err)
	}

	return db, nil
}

func CreateDatabase(name string) {
	url := "server=localhost;port=1433;user id=sa;password=Contraseñña1;ssl=disable;encrypt=disable"
	db, err := sql.Open("sqlserver", url)
	if err != nil {
		panic(err)
	}
	retryCount := 30
	for {
		err := db.Ping()
		if err != nil {
			if retryCount == 0 {
				log.Fatalf("Not able to establish connection to database %s", err)
			}

			log.Printf(fmt.Sprintf("Could not connect to database. Wait 2 seconds. %d retries left...", retryCount))
			retryCount--
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	query := fmt.Sprintf("IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = N'%s') CREATE DATABASE [%s]", name, name)
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}


}
