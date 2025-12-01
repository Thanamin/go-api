package bootstrap

import (
	"log"

	"gorm.io/gorm"

	dbcfg "github.com/thanamin/go-api/internal/infrastructures/config/db"
	envs "github.com/thanamin/go-api/internal/infrastructures/config/environments"
)

// InitDatabase initializes and returns database connection
func InitDatabase() *gorm.DB {
	postgresCfg := dbcfg.PostgresConfig{
		Host:     envs.GetDBHost(),
		Port:     envs.GetDBPort(),
		User:     envs.GetDBUser(),
		Password: envs.GetDBPassword(),
		DBName:   envs.GetDBName(),
		SSLMode:  envs.GetDBSSLMode(),
	}
	db, err := dbcfg.NewPostgresConnection(postgresCfg)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	return db
}
