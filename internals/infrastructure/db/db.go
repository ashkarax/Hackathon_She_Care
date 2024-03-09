package db

import (
	"database/sql"
	"fmt"
	"shecare/internals/config"
	"shecare/internals/domain"
	helper "shecare/pkg"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config.Config) (*gorm.DB, error) {
	connectionString := "user=postgres password=1234 host=localhost"
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = '" + config.DBName + "'")
	if err != nil {
		fmt.Println("Error checking database existence:", err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Database" + config.DBName + " already exists.")
	} else {
		_, err = sql.Exec("CREATE DATABASE " + config.DBName)
		if err != nil {
			fmt.Println("Error creating database:", err)
		}
	}

	DB, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC() // Set the timezone to UTC
		},
	})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(domain.Admin{})
	if err != nil {
		return nil, err
	}
	err = DB.AutoMigrate(domain.Users{})
	if err != nil {
		return nil, err
	}
	err = DB.AutoMigrate(domain.Posts{})
	if err != nil {
		return nil, err
	}

	CheckAndCreateAdmin(DB)

	return DB, nil
}

func CheckAndCreateAdmin(DB *gorm.DB) {
	var count int
	var (
		Name     = "she care"
		Email    = "shecare@gmail.com"
		Password = "ladiesFirst"
	)
	HashedPassword := helper.HashPassword(Password)

	query := "SELECT COUNT(*) FROM admins"
	DB.Raw(query).Row().Scan(&count)
	if count <= 0 {
		query = "INSERT INTO admins(name, email, password) VALUES(?, ?, ?)"
		DB.Exec(query, Name, Email, HashedPassword).Row().Err()
	}
}
