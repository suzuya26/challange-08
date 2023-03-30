package service

import "sesi_8/model"

type BookService interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookById(bookID int) (res model.Book, err error)
	UpdateBook(bookID int, in model.Book) (res model.Book, err error)
	DeleteBook(bookID int) (res string, err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	//call repo
	res, err = s.repo.CreateBook(in)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetAllBook() (res []model.Book, err error) {
	res, err = s.repo.GetAllBook()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetBookById(bookID int) (res model.Book, err error) {

	res, err = s.repo.GetBookById(bookID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) UpdateBook(bookID int, in model.Book) (res model.Book, err error) {
	res, err = s.repo.UpdateBook(bookID, in)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) DeleteBook(bookID int) (res string, err error) {
	res, err = s.repo.DeleteBook(bookID)
	if err != nil {
		return res, err
	}
	return res, nil
}
