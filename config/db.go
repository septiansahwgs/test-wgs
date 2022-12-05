package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

type DB struct {
	SQL *sqlx.DB
}

var dbConn = &DB{}

func ConnectSQL() (*DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")
	hostName := os.Getenv("DB_HOST")
	userName := os.Getenv("DB_USER")
	hostPort := os.Getenv("DB_PORT")

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		userName,
		password,
		hostName,
		hostPort,
		dbName)

	fmt.Println("Datasource", url)

	d, err := sqlx.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	d.SetMaxIdleConns(10)
	d.SetMaxOpenConns(10)
	d.SetConnMaxLifetime(5 * time.Minute)

	migrations := &migrate.FileMigrationSource{
		Dir: "./config/migrations",
	}

	migrate.SetTable("migrations")

	sql := d.DB
	n, err := migrate.Exec(sql, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println("Error occcured:", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)

	dbConn.SQL = d
	return dbConn, err
}

func ConnectSQLTest() (*DB, error) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		"admin123",
		"localhost",
		"5432",
		"db_products")

	fmt.Println("Datasource", url)

	d, err := sqlx.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	d.SetMaxIdleConns(10)
	d.SetMaxOpenConns(10)
	d.SetConnMaxLifetime(5 * time.Minute)

	dbConn.SQL = d
	return dbConn, err
}
