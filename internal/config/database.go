package config

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	url := os.Getenv("sql_connection")
	slog.Info("Database", slog.String("url", url))
	db, err := PingDb(url)
	if err != nil {
		return nil, fmt.Errorf("db: open %v", err)
	}
	return db, nil

}

func PingDb(url string) (*gorm.DB, error) {
	db, err := sql.Open("sqlserver", url)
	if err != nil {
		panic(err)
	}
	retryCount := 30
	for {
		err := db.Ping()
		if err != nil {
			if strings.Contains(err.Error(), "requested by the login") {
				log.Fatal(slog.String("SQL Exception", err.Error()))
				break

			}
			if retryCount == 0 {
				log.Fatal("No se pudo conectar con la base de datos", slog.String("Exception", err.Error()))
			}
			slog.Warn("Reintenteando conexion con la base de datos", slog.Int("Retries left", retryCount))
			retryCount--
			time.Sleep(4 * time.Second)
		} else {
			break
		}
	}

	slog.Info("Db conectada")
	dbG, err := gorm.Open(sqlserver.Open(url), &gorm.Config{})

	return dbG, nil
}
