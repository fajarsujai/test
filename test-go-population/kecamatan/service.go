package kecamatan

import "errors"

type Service interface {
	CreateKecamatan(input CreateKecamatanInput) (Kecamatan, error)
	GetKecamatanByID(input GetKecamatanDetailInput) (Kecamatan, error)
	GetAllKecamatans() ([]Kecamatan, error)
	UpdateKecamatan(inputID GetKecamatanDetailInput, inputData CreateKecamatanInput) (Kecamatan, error)
	DeleteByID(input GetKecamatanDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DeleteByID(input GetKecamatanDetailInput) error {
	err := s.repository.DeleteByID(input.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateKecamatan(input CreateKecamatanInput) (Kecamatan, error) {
	kecamatan := Kecamatan{}
	kecamatan.KecamatanName = input.KecamatanName
	kecamatan.KotaID = input.KotaID

	newKecamatan, err := s.repository.Save(kecamatan)
	if err != nil {
		return newKecamatan, err
	}

	return newKecamatan, nil
}

func (s *service) GetKecamatanByID(input GetKecamatanDetailInput) (Kecamatan, error) {
	kecamatan, err := s.repository.FindByID(input.ID)
	if err != nil {
		return kecamatan, err
	}

	if kecamatan.ID == 0 {
		return kecamatan, errors.New("No kecamatan found on with that ID")
	}

	return kecamatan, nil
}

func (s *service) GetAllKecamatans() ([]Kecamatan, error) {
	kecamatans, err := s.repository.FindAll()
	if err != nil {
		return kecamatans, err
	}

	return kecamatans, nil
}

func (s *service) UpdateKecamatan(inputID GetKecamatanDetailInput, inputData CreateKecamatanInput) (Kecamatan, error) {
	kecamatan, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return kecamatan, err
	}

	kecamatan.KecamatanName = inputData.KecamatanName
	kecamatan.KotaID = inputData.KotaID

	updatedKecamatan, err := s.repository.Update(kecamatan)
	if err != nil {
		return updatedKecamatan, err
	}

	return updatedKecamatan, nil
}
