package config

import "os"

func SetupEnv() {
	os.Setenv("DB_CONNECTION", "postgres")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_DATABASE", "goweb1")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "123456")
	os.Setenv("DB_CHARSET", "utf8")
	os.Setenv("DB_PARSETIME", "True")
	os.Setenv("SSLMODE", "disable")
	os.Setenv("SERVER_PORT", ":8080")
}