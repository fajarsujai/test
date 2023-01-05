package kota

import (
	"errors"
)

type Service interface {
	CreateKota(input CreateKotaInput) (Kota, error)
	GetKotaByID(input GetKotaDetailInput) (Kota, error)
	GetAllKotas() ([]Kota, error)
	UpdateKota(inputID GetKotaDetailInput, inputData CreateKotaInput) (Kota, error)
	DeleteByID(input GetKotaDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DeleteByID(input GetKotaDetailInput) error {
	err := s.repository.DeleteByID(input.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateKota(input CreateKotaInput) (Kota, error) {
	kota := Kota{}
	kota.KotaName = input.KotaName
	kota.ProvinsiID = input.ProvinsiID

	newKota, err := s.repository.Save(kota)
	if err != nil {
		return newKota, err
	}

	return newKota, nil
}

func (s *service) GetKotaByID(input GetKotaDetailInput) (Kota, error) {
	kota, err := s.repository.FindByID(input.ID)
	if err != nil {
		return kota, err
	}

	if kota.ID == 0 {
		return kota, errors.New("No kota found on with that ID")
	}

	return kota, nil
}

func (s *service) GetAllKotas() ([]Kota, error) {
	kotas, err := s.repository.FindAll()
	if err != nil {
		return kotas, err
	}

	return kotas, nil
}

func (s *service) UpdateKota(inputID GetKotaDetailInput, inputData CreateKotaInput) (Kota, error) {
	kota, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return kota, err
	}

	kota.KotaName = inputData.KotaName
	kota.ProvinsiID = inputData.ProvinsiID

	updatedKota, err := s.repository.Update(kota)
	if err != nil {
		return updatedKota, err
	}

	return updatedKota, nil
}
