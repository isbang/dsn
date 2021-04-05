package postgres

import (
	"fmt"
	"strings"
)

type Config struct {
	Host string
	Port int
	User string
	Pass string
	DB   string
}

func (cfg Config) DSN() string {
	return strings.Join([]string{
		fmt.Sprintf("user=%s", cfg.User),
		fmt.Sprintf("password=%s", cfg.Pass),
		fmt.Sprintf("host=%s", cfg.Host),
		fmt.Sprintf("port=%d", cfg.Port),
		fmt.Sprintf("dbname=%s", cfg.DB),
	}, " ")
}
