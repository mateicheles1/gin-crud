package config

import "fmt"

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (c DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s", c.Host, c.Port, c.User, c.Password, c.Database)
}
