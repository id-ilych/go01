package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"go01/pkg/models"
)

type Config struct {
	Filename string
	Table    string
}

type HotdogDatabase struct {
	db         *sql.DB
	stmtInsert string
}

func OpenHotdogDatabase(cfg Config) (*HotdogDatabase, error) {
	log.Printf("opening database: %s", cfg.Filename)
	db, err := sql.Open("sqlite3", cfg.Filename)
	if err != nil {
		return nil, err
	}

	log.Printf("ensuring '%s' table exists", cfg.Table)
	stmtEnsureTable := stmtEnsureHotdogsTable(cfg.Table)
	if _, err = db.Exec(stmtEnsureTable); err != nil {
		return nil, err
	}

	log.Printf("hotdogs database is ready")
	stmtInsert := stmtInsertHotdog(cfg.Table)
	return &HotdogDatabase{db, stmtInsert}, nil
}

func (db *HotdogDatabase) SaveHotdog(hotdog *models.Hotdog) error {
	tx, err := db.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(db.stmtInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(hotdog.Title, hotdog.Calories, hotdog.Price)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (db *HotdogDatabase) Close() error {
	return db.db.Close()
}

func stmtEnsureHotdogsTable(table string) string {
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (title TEXT, calories INTEGER, price INTEGER);", table)
}

func stmtInsertHotdog(table string) string {
	return fmt.Sprintf("INSERT INTO %s (title, calories, price) VALUES(?, ?, ?)", table)
}
