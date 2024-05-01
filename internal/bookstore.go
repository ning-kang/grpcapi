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

func (bs *BookStore) CreateBook(ctx context.Context, in *bookstore.CreateBookInput) (*bookstore.Book, error) {
	book := &bookstore.Book{Title: in.Title, Author: in.Author}
	result := bs.DB.Create(book) // Primary key ID is already inserted
	if result.Error != nil {
		log.Errorf("Failed to create book %s: %s", book, result.Error.Error())
		return nil, result.Error
	}
	return book, nil
}

func (bs *BookStore) UpdateBook(ctx context.Context, in *bookstore.UpdateBookInput) (*bookstore.Book, error) {
	book, err := bs.GetBook(ctx, &bookstore.BookId{Id: in.ID})
	if err != nil {
		return nil, err
	}
	result := bs.DB.Model(book).Updates(in)
	if result.Error != nil {
		log.Errorf("Failed to create book %s: %s", book, result.Error.Error())
		return nil, result.Error
	}
	return book, nil

}

func (bs *BookStore) DeleteBook(ctx context.Context, in *bookstore.BookId) (*bookstore.Empty, error) {
	book, err := bs.GetBook(ctx, in)
	if err != nil {
		return nil, err
	}
	result := bs.DB.Delete(book)
	if result.Error != nil {
		log.Errorf("Failed to delete book %s: %s", book, result.Error.Error())
		return nil, result.Error
	}
	return nil, nil
}
