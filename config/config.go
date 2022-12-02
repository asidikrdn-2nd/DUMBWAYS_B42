package config

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

// membuat fungsi untuk koneksi ke database, fungsi ini dapat dipanggil oleh package lainnya
func CreateConnection() (*pgx.Conn, error) {
	// memulai koneksi ke database
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	// cek error saat menghubungkan ke database
	if err != nil {
		return nil, err
	}

	return conn, nil
}
