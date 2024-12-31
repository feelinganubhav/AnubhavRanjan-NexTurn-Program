package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // SQLite driver
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "./ecommerce.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Create Products table if it does not already exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		price REAL NOT NULL,
		stock INTEGER NOT NULL DEFAULT 0,
		category_id INTEGER
	);`)
	if err != nil {
		log.Fatal("Failed to create products table:", err)
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
