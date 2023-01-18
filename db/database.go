package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("gagal membaca file .env ", err)
	}
	fmt.Println("berhasil membaca file .env")
	return os.Getenv(key)
}

func InitDB() (*pgx.Conn, error) {

	conn, err := pgx.Connect(context.Background(), GetEnv("DB_URL"))
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
		return nil, err
	}
	fmt.Println("berhasil terhubung dengan database")

	return conn, nil
}

func CloseDB(conn *pgx.Conn) {
	defer conn.Close(context.Background())
}
