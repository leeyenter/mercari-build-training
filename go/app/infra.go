package app

import (
	"database/sql"
	"errors"
	"log"
	"log/slog"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var errImageNotFound = errors.New("image not found")

type Item struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Category string `db:"category" json:"category"`
	Image    string `db:"image" json:"image_name"`
}

type ItemRepository interface {
	Insert(item *Item) error
	GetAllItems() ([]*Item, error)
}

// itemRepository is an implementation of ItemRepository
type itemRepository struct {
	db *sql.DB
}

func seedDB(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS items (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, itemname TEXT, category TEXT, image TEXT)`)
	if err != nil {
		panic(err)
	}
}

// NewItemRepository creates a new itemRepository.
func NewItemRepository() ItemRepository {
	db, err := sql.Open("sqlite3", "db/db.db")
	if err != nil {
		log.Fatal(err)
	}

	seedDB(db)

	//defer db.Close()
	return &itemRepository{db: db}
}

// Insert inserts an item into the repository.
func (i *itemRepository) Insert(item *Item) error {
	stmt, err := i.db.Prepare(`INSERT INTO items (itemname, category, image) VALUES (?, ?, ?)`)
	if err != nil {
		slog.Error("Could not prepare statement", err)
		return err
	}

	_, err = stmt.Exec(item.Name, item.Category, item.Image)
	if err != nil {
		return err
	}

	return nil
}

func (i *itemRepository) GetAllItems() ([]*Item, error) {
	rows, err := i.db.Query("SELECT id, itemname, category, image FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]*Item, 0)
	for rows.Next() {
		item := &Item{}
		err = rows.Scan(&item.ID, &item.Name, &item.Category, &item.Image)
		if err != nil {
			return nil, err
		}

		item.Image = strings.Replace(item.Image, "images/", "", 1)
		items = append(items, item)
	}
	return items, nil
}

// StoreImage stores an image and returns an error if any.
// This package doesn't have a related interface for simplicity.
func StoreImage(fileName string, image []byte) error {
	// STEP 4-4: add an implementation to store an image

	return nil
}
