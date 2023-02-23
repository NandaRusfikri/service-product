package dto

import "github.com/gin-contrib/cors"

type AppConfig struct {
	Debug        bool
	Host         string
	Port         string
	UseTLS       bool
	CertFilePath string
	KeyFilePath  string
	Cors         *cors.Config
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Dialect  string
	TimeZone string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Password string
	DB       int
}

type SMTPConfig struct {
	Host     string
	Port     int
	Email    string
	Password string
	Name     string
}
