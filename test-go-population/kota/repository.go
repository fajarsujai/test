package kota

import "gorm.io/gorm"

type Repository interface {
	Save(kota Kota) (Kota, error)
	FindByID(ID int) (Kota, error)
	Update(kota Kota) (Kota, error)
	FindAll() ([]Kota, error)
	DeleteByID(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) DeleteByID(ID int) error {
	var kota Kota
	err := r.db.Where("id = ?", ID).Delete(&kota).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Save(kota Kota) (Kota, error) {
	err := r.db.Create(&kota).Error
	if err != nil {
		return kota, err
	}

	return kota, nil
}

func (r *repository) FindByID(ID int) (Kota, error) {
	var kota Kota

	err := r.db.Where("id = ?", ID).Find(&kota).Error
	if err != nil {
		return kota, err
	}

	return kota, nil
}

func (r *repository) Update(kota Kota) (Kota, error) {
	err := r.db.Save(&kota).Error

	if err != nil {
		return kota, err
	}

	return kota, nil
}

func (r *repository) FindAll() ([]Kota, error) {
	var kota []Kota

	err := r.db.Find(&kota).Error
	if err != nil {
		return kota, err
	}

	return kota, nil
}
