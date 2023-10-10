package models

import (
	"time"
)

type Author struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	// Books []Book
}

type Book struct {
	ID       uint   `gorm:"primaryKey"`
	ISBN     string `gorm:"uniqueKey"`
	Title    string
	Genre    string
	AuthorId uint
	// Loans    []Loan
	AuthorID uint
}

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"uniqueKey"`
	Password string
	// Loans    []Loan
}

type Loan struct {
	ID         uint `gorm:"primaryKey"`
	BorrowDate *time.Time
	ReturnDate *time.Time
	UserID     uint
	BookID     uint
}
