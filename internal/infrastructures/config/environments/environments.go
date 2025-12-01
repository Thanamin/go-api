package environments

import (
	"os"
	"strconv"
	"sync"
)

type envConfig struct {
	Port             string
	DBURI            string
	DBName           string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBSSLMode        string
	HashMethod       string
	HashSecret       string
	BcryptCost       int
	JWTSecret        string
	JWTExpirationHrs int
}

var cfg envConfig
var once sync.Once

func load() {
	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" {
		cfg.Port = "3000"
	}
	cfg.DBURI = os.Getenv("DB_URI")
	cfg.DBName = os.Getenv("DB_NAME")

	// Postgres specific
	cfg.DBHost = os.Getenv("DB_HOST")
	if cfg.DBHost == "" {
		cfg.DBHost = "localhost"
	}
	cfg.DBPort = os.Getenv("DB_PORT")
	if cfg.DBPort == "" {
		cfg.DBPort = "5432"
	}
	cfg.DBUser = os.Getenv("DB_USER")
	if cfg.DBUser == "" {
		cfg.DBUser = "postgres"
	}
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	if cfg.DBPassword == "" {
		cfg.DBPassword = "postgres"
	}
	cfg.DBSSLMode = os.Getenv("DB_SSLMODE")
	if cfg.DBSSLMode == "" {
		cfg.DBSSLMode = "disable"
	}

	// Password hashing config
	cfg.HashMethod = os.Getenv("HASH_METHOD")
	if cfg.HashMethod == "" {
		cfg.HashMethod = "bcrypt"
	}
	cfg.HashSecret = os.Getenv("HASH_SECRET")
	if cfg.HashSecret == "" {
		cfg.HashSecret = "default-secret-key"
	}
	cfg.BcryptCost = 10
	if costStr := os.Getenv("BCRYPT_COST"); costStr != "" {
		if cost, err := strconv.Atoi(costStr); err == nil && cost >= 4 && cost <= 31 {
			cfg.BcryptCost = cost
		}
	}

	// JWT config
	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	if cfg.JWTSecret == "" {
		cfg.JWTSecret = "default-jwt-secret"
	}
	cfg.JWTExpirationHrs = 24
	if expStr := os.Getenv("JWT_EXPIRATION_HOURS"); expStr != "" {
		if exp, err := strconv.Atoi(expStr); err == nil && exp > 0 {
			cfg.JWTExpirationHrs = exp
		}
	}
}

// GetPort returns the configured port (default "3000")
func GetPort() string {
	once.Do(load)
	return cfg.Port
}

// GetDBURI returns the DB URI (may be empty if not set)
func GetDBURI() string {
	once.Do(load)
	return cfg.DBURI
}

// GetDBName returns the DB name (may be empty if not set)
func GetDBName() string {
	once.Do(load)
	return cfg.DBName
}

// GetDBHost returns the Postgres host
func GetDBHost() string {
	once.Do(load)
	return cfg.DBHost
}

// GetDBPort returns the Postgres port
func GetDBPort() string {
	once.Do(load)
	return cfg.DBPort
}

// GetDBUser returns the Postgres user
func GetDBUser() string {
	once.Do(load)
	return cfg.DBUser
}

// GetDBPassword returns the Postgres password
func GetDBPassword() string {
	once.Do(load)
	return cfg.DBPassword
}

// GetDBSSLMode returns the Postgres SSL mode
func GetDBSSLMode() string {
	once.Do(load)
	return cfg.DBSSLMode
}

// GetHashMethod returns the password hash method
func GetHashMethod() string {
	once.Do(load)
	return cfg.HashMethod
}

// GetHashSecret returns the password hash secret
func GetHashSecret() string {
	once.Do(load)
	return cfg.HashSecret
}

// GetBcryptCost returns the bcrypt cost
func GetBcryptCost() int {
	once.Do(load)
	return cfg.BcryptCost
}

// GetJWTSecret returns the JWT secret
func GetJWTSecret() string {
	once.Do(load)
	return cfg.JWTSecret
}

// GetJWTExpirationHours returns the JWT expiration in hours
func GetJWTExpirationHours() int {
	once.Do(load)
	return cfg.JWTExpirationHrs
}
