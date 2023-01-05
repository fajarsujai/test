package provinsi

import (
	"errors"
	"populationdatago/helper"
)

type Service interface {
	CreateProvinsi(input CreateProvinsiInput) (Provinsi, error)
	GetProvinsiByID(input GetProvinsiDetailInput) (Provinsi, error)
	GetAllProvinsis(pagination helper.Pagination) (*helper.Pagination, error)
	UpdateProvinsi(inputID GetProvinsiDetailInput, inputData CreateProvinsiInput) (Provinsi, error)
	DeleteByID(input GetProvinsiDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DeleteByID(input GetProvinsiDetailInput) error {
	err := s.repository.DeleteByID(input.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateProvinsi(input CreateProvinsiInput) (Provinsi, error) {
	provinsi := Provinsi{}
	provinsi.ProvinsiName = input.ProvinsiName
	provinsi.ProvinsiCode = input.ProvinsiCode

	newProvinsi, err := s.repository.Save(provinsi)
	if err != nil {
		return newProvinsi, err
	}

	return newProvinsi, nil
}

func (s *service) GetProvinsiByID(input GetProvinsiDetailInput) (*Provinsi, error) {
	provinsi, err := s.repository.FindByID(input.ID)

	if err != nil {
		return provinsi, err
	}

	if provinsi.ID == 0 {
		return provinsi, errors.New("No provinsi found on with that ID")
	}

	return provinsi, nil
}

func (s *service) GetAllProvinsis(pagination helper.Pagination) (*helper.Pagination, error) {
	provinsis, err := s.repository.FindAll(pagination)
	if err != nil {
		return provinsis, err
	}

	return provinsis, nil
}

// func (s *service) UpdateProvinsi(inputID GetProvinsiDetailInput, inputData CreateProvinsiInput) (Provinsi, error) {
// 	provinsi, err := s.repository.FindByID(inputID.ID)
// 	if err != nil {
// 		return provinsi, err
// 	}

// 	provinsi.ProvinsiName = inputData.ProvinsiName
// 	provinsi.ProvinsiCode = inputData.ProvinsiCode

// 	updatedProvinsi, err := s.repository.Update(provinsi)
// 	if err != nil {
// 		return updatedProvinsi, err
// 	}

// 	return updatedProvinsi, nil
// }
