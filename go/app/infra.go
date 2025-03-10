package app

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
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
	db *pgx.Conn
}

func seedDB(db *pgx.Conn) {
	_, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS items (id SERIAL NOT NULL PRIMARY KEY, itemname TEXT, category TEXT, image TEXT)`)
	if err != nil {
		panic(err)
	}
}

// NewItemRepository creates a new itemRepository.
func NewItemRepository() ItemRepository {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:mypassword@"+dbHost+":5432/postgres")
	if err != nil {
		panic(err)
	}

	seedDB(conn)

	//defer db.Close()
	return &itemRepository{db: conn}
}

// Insert inserts an item into the repository.
func (i *itemRepository) Insert(item *Item) error {
	_, err := i.db.Exec(context.Background(), `INSERT INTO items (itemname, category, image) VALUES ($1, $2, $3)`, item.Name, item.Category, item.Image)
	if err != nil {
		slog.Error("Could not prepare statement", err)
		return err
	}

	return nil
}

func (i *itemRepository) GetAllItems() ([]*Item, error) {
	rows, err := i.db.Query(context.Background(), "SELECT id, itemname, category, image FROM items")
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
