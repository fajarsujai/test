package memberdetail

import (
	"errors"
)

type Service interface {
	CreateMemberDetail(input CreateMemberDetailInput) (MemberDetail, error)
	GetMemberDetailByID(input GetMemberDetailDetailInput) (MemberDetail, error)
	GetAllMemberDetails() ([]MemberDetail, error)
	UpdateMemberDetail(inputID GetMemberDetailDetailInput, inputData CreateMemberDetailInput) (MemberDetail, error)
	DeleteByID(input GetMemberDetailDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DeleteByID(input GetMemberDetailDetailInput) error {
	err := s.repository.DeleteByID(input.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateMemberDetail(input CreateMemberDetailInput) (MemberDetail, error) {
	memberdetail := MemberDetail{}
	memberdetail.UserID = input.UserID
	memberdetail.Nik = input.Nik
	memberdetail.PhoneNumber = input.PhoneNumber
	memberdetail.Address = input.Address
	memberdetail.ProvinsiID = input.ProvinsiID
	memberdetail.KotaID = input.KotaID
	memberdetail.KecamatanID = input.KecamatanID
	memberdetail.KelurahanID = input.KelurahanID

	newMebmerDetail, err := s.repository.Save(memberdetail)
	if err != nil {
		return newMebmerDetail, err
	}

	return newMebmerDetail, nil
}

func (s *service) GetMemberDetailByID(input GetMemberDetailDetailInput) (MemberDetail, error) {
	memberdetail, err := s.repository.FindByID(input.ID)
	if err != nil {
		return memberdetail, err
	}

	if memberdetail.ID == 0 {
		return memberdetail, errors.New("No memberdetail found on with that ID")
	}

	return memberdetail, nil
}

func (s *service) GetAllMemberDetails() ([]MemberDetail, error) {
	memberdetails, err := s.repository.FindAll()
	if err != nil {
		return memberdetails, err
	}

	return memberdetails, nil
}

func (s *service) UpdateMemberDetail(inputID GetMemberDetailDetailInput, inputData CreateMemberDetailInput) (MemberDetail, error) {
	memberdetail, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return memberdetail, err
	}

	memberdetail.UserID = inputData.UserID
	memberdetail.Nik = inputData.Nik
	memberdetail.PhoneNumber = inputData.PhoneNumber
	memberdetail.Address = inputData.Address
	memberdetail.ProvinsiID = inputData.ProvinsiID
	memberdetail.KotaID = inputData.KotaID
	memberdetail.KecamatanID = inputData.KecamatanID
	memberdetail.KelurahanID = inputData.KelurahanID

	updatedMemberDetail, err := s.repository.Update(memberdetail)
	if err != nil {
		return updatedMemberDetail, err
	}

	return updatedMemberDetail, nil
}
