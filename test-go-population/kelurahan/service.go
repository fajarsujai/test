package kelurahan

import "errors"

type Service interface {
	CreateKelurahan(input CreateKelurahanInput) (Kelurahan, error)
	GetKelurahanByID(input GetKelurahanDetailInput) (Kelurahan, error)
	GetAllKelurahans() ([]Kelurahan, error)
	UpdateKelurahan(inputID GetKelurahanDetailInput, inputData CreateKelurahanInput) (Kelurahan, error)
	DeleteByID(input GetKelurahanDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DeleteByID(input GetKelurahanDetailInput) error {
	err := s.repository.DeleteByID(input.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateKelurahan(input CreateKelurahanInput) (Kelurahan, error) {
	kelurahan := Kelurahan{}
	kelurahan.KelurahanName = input.KelurahanName
	kelurahan.KecamatanID = input.KecamatanID

	newKelurahan, err := s.repository.Save(kelurahan)
	if err != nil {
		return newKelurahan, err
	}

	return newKelurahan, nil
}

func (s *service) GetKelurahanByID(input GetKelurahanDetailInput) (Kelurahan, error) {
	kelurahan, err := s.repository.FindByID(input.ID)
	if err != nil {
		return kelurahan, err
	}

	if kelurahan.ID == 0 {
		return kelurahan, errors.New("No kelurahan found on with that ID")
	}

	return kelurahan, nil
}

func (s *service) GetAllKelurahans() ([]Kelurahan, error) {
	kelurahans, err := s.repository.FindAll()
	if err != nil {
		return kelurahans, err
	}

	return kelurahans, nil
}

func (s *service) UpdateKelurahan(inputID GetKelurahanDetailInput, inputData CreateKelurahanInput) (Kelurahan, error) {
	kelurahan, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return kelurahan, err
	}

	kelurahan.KelurahanName = inputData.KelurahanName
	kelurahan.KecamatanID = inputData.KecamatanID

	updatedKelurahan, err := s.repository.Update(kelurahan)
	if err != nil {
		return updatedKelurahan, err
	}

	return updatedKelurahan, nil
}
