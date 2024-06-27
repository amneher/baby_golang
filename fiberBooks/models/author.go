package models

import (
	"fmt"
	"time"

	"github.com/amneher/fiberBooks/initializers"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name     string    `json:"name"`
	DOB      time.Time `json:"birthdate"`
	Deceased bool      `json:"deceased" gorm:"default:false"`
	Genre    string    `json:"genre" gorm:"index"`
	Books    []Book    `json:"books" gorm:"foreignKey:Author;nullable"`
}

func (author *Author) String() (s string) {
	s = fmt.Sprintf("<Author %v Name: %v, Genre: %v, DOB: %v>", author.ID, author.Name, author.Genre, author.DOB)
	return
}

func (author *Author) Map() (m map[string]string) {
	m = make(map[string]string)
	m["id"] = string(author.ID)
	m["createdAt"] = author.CreatedAt.String()
	m["updatedAt"] = author.UpdatedAt.String()
	m["deleted"] = author.DeletedAt.Time.String()
	m["name"] = author.Name
	m["birthdate"] = author.DOB.String()
	m["genre"] = author.Genre
	return
}

func (author *Author) Alive() (alive bool) {
	y, _, _ := time.Now().Date()
	ay, _, _ := author.DOB.Date()
	if y-ay >= 100 {
		alive = false
	} else {
		alive = true
	}
	return
}

func (author *Author) Birthday() (party bool, age int) {
	ay, am, ad := author.DOB.Date()
	y, m, d := time.Now().Date()
	if (am.String() == m.String()) && (ad == d) {
		party = true
		age = y - ay
	} else {
		party = false
		age = 0
	}
	return
}

func (author *Author) Age() (age int) {
	ay, am, ad := author.DOB.Date()
	y, m, d := time.Now().Date()
	age = y - ay
	if m < am {
		age -= 1
	} else {
		if (m == am) && (d < ad) {
			age -= 1
		}
	}
	return
}

func CreateAuthor(author *Author) *gorm.DB {
	return initializers.DB.Create(author)
}

func GetAllAuthors(dest interface{}) *gorm.DB {
	return initializers.DB.Model(&Author{}).Find(dest)
}

func FindAuthor(dest interface{}, conds ...interface{}) *gorm.DB {
	return initializers.DB.Model(&Author{}).Take(dest, conds...)
}

func GetAuthorByID(dest interface{}, id uint) *gorm.DB {
	return FindAuthor(dest, "id = ?", id)
}
