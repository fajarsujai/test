package kecamatan

import "gorm.io/gorm"

type Repository interface {
	Save(kecamatan Kecamatan) (Kecamatan, error)
	FindByID(ID int) (Kecamatan, error)
	Update(kecamatan Kecamatan) (Kecamatan, error)
	FindAll() ([]Kecamatan, error)
	DeleteByID(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) DeleteByID(ID int) error {
	var kecamatan Kecamatan
	err := r.db.Where("id = ?", ID).Delete(&kecamatan).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Save(kecamatan Kecamatan) (Kecamatan, error) {
	err := r.db.Create(&kecamatan).Error
	if err != nil {
		return kecamatan, err
	}

	return kecamatan, nil
}

func (r *repository) FindByID(ID int) (Kecamatan, error) {
	var kecamatan Kecamatan

	err := r.db.Where("id = ?", ID).Find(&kecamatan).Error
	if err != nil {
		return kecamatan, err
	}

	return kecamatan, nil
}

func (r *repository) Update(kecamatan Kecamatan) (Kecamatan, error) {
	err := r.db.Save(&kecamatan).Error

	if err != nil {
		return kecamatan, err
	}

	return kecamatan, nil
}

func (r *repository) FindAll() ([]Kecamatan, error) {
	var kecamatan []Kecamatan

	err := r.db.Find(&kecamatan).Error
	if err != nil {
		return kecamatan, err
	}

	return kecamatan, nil
}
