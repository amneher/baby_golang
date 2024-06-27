package models

import (
	"time"

	"github.com/amneher/fiberBooks/initializers"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string    `json:"title"`
	Published time.Time `json:"publishDate"`
	Pages     int       `json:"pages"`
	Genre     string    `json:"genre"`
	Author    *uint     `json:"author" gorm:"index,not null"`
}

func CreateBook(book *Book) *gorm.DB {
	return initializers.DB.Create(book)
}

func GetAllBooks(dest interface{}) *gorm.DB {
	return initializers.DB.Model(&Book{}).Find(dest)
}

func FindBook(dest interface{}, conds ...interface{}) *gorm.DB {
	return initializers.DB.Model(&Book{}).Take(dest, conds...)
}

// FindTodoByAuthor finds a todo with given todo and Author identifier
func FindBookByAuthor(dest interface{}, bookIden interface{}, authorIden interface{}) *gorm.DB {
	return FindBook(dest, "id = ? AND author = ?", bookIden, authorIden)
}

// FindbooksByAuthor finds the books with Author's identifier given
func FindBooksByAuthor(dest interface{}, authorIden interface{}) *gorm.DB {
	return initializers.DB.Model(&Book{}).Find(dest, "author = ?", authorIden)
}

// Deletebook deletes a book from books' table with the given book and Author identifier
func DeleteBook(bookIden interface{}, authorIden interface{}) *gorm.DB {
	return initializers.DB.Unscoped().Delete(&Book{}, "id = ? AND author = ?", bookIden, authorIden)
}

// Updatebook allows to update the book with the given bookID and AuthorID
func UpdateBook(bookIden interface{}, authorIden interface{}, data interface{}) *gorm.DB {
	return initializers.DB.Model(&Book{}).Where("id = ? AND author = ?", bookIden, authorIden).Updates(data)
}
