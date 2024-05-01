package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/google/martian/v3/log"
	"github.com/ning-kang/grpcapi/protogen/golang/bookstore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BookStore struct {
	DB *gorm.DB
	bookstore.UnimplementedBookStoreServer
}

func NewBookStore() *BookStore {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	dsn := fmt.Sprintf("host=localhost port=5432 user=%s password=%s sslmode=disable", dbUser, dbPass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Check if the database exists
	var check string
	db.Raw("SELECT 1 FROM pg_database WHERE datname = 'bookstore'").Scan(&check)

	// If the database does not exist, create it
	if check != "1" {
		db.Exec("CREATE DATABASE bookstore")
	}

	// Close the initial connection
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database from GORM")
	}
	sqlDB.Close()

	// Connect to the new database
	dsn = fmt.Sprintf("host=localhost port=5432 dbname=bookstore user=%s password=%s sslmode=disable", dbUser, dbPass)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the new database")
	}
	// Run Migration on each model
	model := bookstore.Book{}
	err = database.AutoMigrate(&model)
	if err != nil {
		return nil
	}

	return &BookStore{DB: database}
}

func (bs *BookStore) ListBooks(ctx context.Context, in *bookstore.Empty) (*bookstore.BookList, error) {
	var bookList bookstore.BookList // have to be concrete struct
	bs.DB.Find(&bookList.Books)     // modified books value
	return &bookList, nil
}

func (bs *BookStore) GetBook(ctx context.Context, in *bookstore.BookId) (*bookstore.Book, error) {
	var book bookstore.Book // have to be concrete struct
	if err := bs.DB.Where("id = ?", in.Id).First(&book).Error; err != nil {
		log.Errorf("Failed to get %s: %s", in.Id, err.Error())
		return nil, err
	}
	return &book, nil
}
