package config

import "os"

func SetupEnv() {
	os.Setenv("DB_CONNECTION", "mysql")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_DATABASE", "goweb1")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_CHARSET", "utf8")
	os.Setenv("DB_PARSETIME", "True")

	os.Setenv("SERVER_PORT", ":8080")
}