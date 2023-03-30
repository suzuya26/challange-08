package repository

import (
	"fmt"
	"sesi_8/model"
	"time"
)

// interface employee, index repo book
type BookRepo interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookById(bookID int) (res model.Book, err error)
	UpdateBook(bookID int, in model.Book) (res model.Book, err error)
	DeleteBook(bookID int) (res string, err error)
}

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllBook() (res []model.Book, err error) {
	err = r.gorm.Where("deleted_at is null").Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetBookById(bookID int) (res model.Book, err error) {
	err = r.gorm.Where("deleted_at is null").First(&res, bookID).Error
	if err != nil {
		return res, err
	}
	return res, err
}

func (r Repo) UpdateBook(bookID int, in model.Book) (res model.Book, err error) {
	err = r.gorm.Model(&res).Where("id = ?", bookID).Updates(
		model.Book{
			Title:       in.Title,
			Author:      in.Author,
			Description: in.Description,
		}).Scan(&res).Error

	if err != nil {
		return res, err
	}
	// res = fmt.Sprintf("Book dengan Id %v sudah ter-update", bookID)
	return res, err
}

func (r Repo) DeleteBook(bookID int) (res string, err error) {
	book := model.Book{}
	err = r.gorm.Model(&book).Where("id = ?", bookID).Update("deleted_at", time.Now()).Error

	if err != nil {
		return res, err
	}
	res = fmt.Sprintf("Book dengan Id %v sudah ter-hapus :<", bookID)
	return res, err
}
