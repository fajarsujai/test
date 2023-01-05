package memberdetail

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(memberdetail MemberDetail) (MemberDetail, error)
	FindByID(ID int) (MemberDetail, error)
	Update(memberdetail MemberDetail) (MemberDetail, error)
	FindAll() ([]MemberDetail, error)
	DeleteByID(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) DeleteByID(ID int) error {
	var memberdetail MemberDetail
	err := r.db.Where("id = ?", ID).Delete(&memberdetail).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Save(memberdetail MemberDetail) (MemberDetail, error) {
	err := r.db.Create(&memberdetail).Error
	if err != nil {
		return memberdetail, err
	}

	return memberdetail, nil
}

func (r *repository) FindByID(ID int) (MemberDetail, error) {
	var memberdetail MemberDetail

	err := r.db.Where("id = ?", ID).Find(&memberdetail).Error
	if err != nil {
		return memberdetail, err
	}
	return memberdetail, nil
}

func (r *repository) Update(memberdetail MemberDetail) (MemberDetail, error) {
	err := r.db.Save(&memberdetail).Error

	if err != nil {
		return memberdetail, err
	}

	return memberdetail, nil
}

func (r *repository) FindAll() ([]MemberDetail, error) {
	var memberdetail []MemberDetail

	err := r.db.Find(&memberdetail).Error
	if err != nil {
		return memberdetail, err
	}

	return memberdetail, nil
}
