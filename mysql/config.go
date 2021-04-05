package mysql

import "fmt"

type Config struct {
	Host string
	Port int
	User string
	Pass string
	DB   string
}

func (cfg Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&interpolateParams=true&allowNativePasswords=true",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DB)
}
