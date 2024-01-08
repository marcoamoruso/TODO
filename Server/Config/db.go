package Config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// Build database configuration from environment variables
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     portConv(os.Getenv("DB_PORT")),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

// Convert string to int
func portConv(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// Build connection string
func DbURL(dbConfig *DBConfig) string {
	conString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	return conString
}

// Initialize database
func DatabaseInit() {
	dbConfig := BuildDBConfig()
	dbURL := DbURL(dbConfig)

	var err error
	DB, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
}

// Get database
func GetDB() *gorm.DB {
	return DB
}
