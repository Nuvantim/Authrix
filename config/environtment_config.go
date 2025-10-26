package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_NAME"`
}

type RedisConfig struct {
	Host     string `envconfig:"REDIS_HOST"`
	Password string `envconfig:"REDIS_PASSWORD"`
	Port     int    `envconfig:"REDIS_PORT"`
}

type ServerConfig struct {
	Port    string `envconfig:"PORT"`
	AppName string `envconfig:"APP_NAME"`
	Url     string `envconfig:"URL"`
}

type MailConfig struct {
	Mailer   string `envconfig:"MAIL_MAILER"`
	Port     int    `envconfig:"MAIL_PORT"`
	Username string `envconfig:"MAIL_USERNAME"`
	Password string `envconfig:"MAIL_PASSWORD"`
	Address  string `envconfig:"MAIL_FROM_ADDRESS"`
}

func CheckEnv() error {
	if err := godotenv.Load(); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func GetDatabaseConfig() (*DatabaseConfig, error) {
	var db DatabaseConfig
	if err := envconfig.Process("", &db); err != nil {
		return nil, err
	}
	return &db, nil
}

func GetRedisConfig() (*RedisConfig, error) {
	var rds RedisConfig
	if err := envconfig.Process("", &rds); err != nil {
		return nil, err
	}
	return &rds, nil
}

func GetServerConfig() (*ServerConfig, error) {
	var serv ServerConfig
	if err := envconfig.Process("", &serv); err != nil {
		return nil, err
	}
	return &serv, nil
}

func GetMailConfig() (*MailConfig, error) {
	var mail MailConfig
	if err := envconfig.Process("", &mail); err != nil {
		return nil, err
	}
	return &mail, nil
}
