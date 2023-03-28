package database

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type sqlxInstance struct {
	read, write *sqlx.DB
}

func (s *sqlxInstance) Read() *sqlx.DB {
	return s.read
}

func (s *sqlxInstance) Write() *sqlx.DB {
	return s.write
}

func (g *sqlxInstance) Close() {
}

type SqlxDatabase interface {
	Read() *sqlx.DB
	Write() *sqlx.DB
	Close()
}

func GeneratePostgreCredentials() string {

	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbTimeout := os.Getenv("DB_TIMEOUT")

	constring := "user=%s password=%s dbname=%s host=%s port=%s sslmode=disable connect_timeout=%s"

	return fmt.Sprintf(constring, dbUser, dbPassword, dbName, dbHost, dbPort, dbTimeout)
}

func InitSQLX() SqlxDatabase {

	dsn := GeneratePostgreCredentials()

	sqlRead, errRead := sqlx.Open("postgres", dsn)
	if errRead != nil {
		panic(errRead.Error())
	}

	sqlRead.SetMaxIdleConns(5)
	sqlRead.SetMaxOpenConns(20)
	sqlRead.SetConnMaxLifetime(time.Hour)

	sqlWrite, errWrite := sqlx.Open("postgres", dsn)
	if errWrite != nil {
		panic(errWrite.Error())
	}

	sqlWrite.SetMaxIdleConns(5)
	sqlWrite.SetMaxOpenConns(20)
	sqlWrite.SetConnMaxLifetime(time.Hour)

	return &sqlxInstance{
		read:  sqlRead,
		write: sqlWrite,
	}
}
