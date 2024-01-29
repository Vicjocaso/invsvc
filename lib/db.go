package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// Initializing the database, params: url -> the connection string form the mysql database
func InitDB(url string) (*sql.DB, error) {
	// Open a connection to the database
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	// Db = db
	// db.SetMaxIdleConns(20)
	// db.SetMaxOpenConns(20)
	// db.SetConnMaxLifetime(time.Duration(2) * time.Minute)

	log.Println("Successfully connected to PlanetScale!")

	// db.Exec(`CREATE TABLE IF NOT EXISTS invoice (
	// 	id int NOT NULL AUTO_INCREMENT,
	// 	client varchar(255),
	// 	price int,
	// 	status varchar(255)
	// 	PRIMARY KEY (id)
	// 	);`)

	return db, nil

}
