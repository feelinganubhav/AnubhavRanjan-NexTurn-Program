package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // SQLite driver
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "./blog_management.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Create Blogs table if it does not already exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Fatal("Failed to create blogs table:", err)
	}

	// Create Users table if it does not already exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	log.Println("Database initialized successfully")
}

func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call InitializeDatabase first.")
	}
	return DB
}
