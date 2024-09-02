package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnector struct {
}

func (p *PostgresConnector) GetConnection() (db *gorm.DB, err error) {
	//e := godotenv.Load()
	// if e != nil {
	// 	fmt.Print(e)
	// }
	// username := os.Getenv("db_user")
	// password := os.Getenv("db_pass")
	// dbName := os.Getenv("db_name")
	// dbHost := os.Getenv("db_host")
	username := "postgres"
	password := ",HTjgKRo"
	dbName := "exercise_ease_db"
	dbHost := "tf-20240901160038747700000001.cpsc4kmqkaao.af-south-1.rds.amazonaws.com"
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s port=5432 sslmode=disable password=%s", dbHost, username, dbName, password)
	return gorm.Open(postgres.Open(dbURI), &gorm.Config{})
}
