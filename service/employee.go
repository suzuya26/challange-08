package service

// usecase

import "sesi_8/model"

type EmployeeService interface {
	CreateEmployee(in model.Employee) (res model.Employee, err error)
	GetEmployeeByID(id int64) (res model.Employee, err error)
	UpdateEmployee(in model.Employee) (res model.Employee, err error)
	DeleteEmployee(id int64) (err error)
}

func (s *Service) CreateEmployee(in model.Employee) (res model.Employee, err error) {
	// call repo
	res, err = s.repo.CreateEmployee(in)
	if err != nil {
		return res, err
	}

	return res, nil

	// return s.repo.CreateEmployee(in)
}

func (s *Service) GetEmployeeByID(id int64) (res model.Employee, err error) {
	return s.repo.GetEmployeeByID(id)
}

func (s *Service) UpdateEmployee(in model.Employee) (res model.Employee, err error) {
	return s.repo.UpdateEmployee(in)
}

func (s *Service) DeleteEmployee(id int64) (err error) {
	return s.repo.DeleteEmployee(id)
}

// func (s *Service) UpdateEmployee(in model.Employee) (res model.Employee, err error) {
// 	// TODO: implement
// }
