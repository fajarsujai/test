package kelurahan

import "gorm.io/gorm"

type Repository interface {
	Save(kelurahan Kelurahan) (Kelurahan, error)
	FindByID(ID int) (Kelurahan, error)
	Update(kelurahan Kelurahan) (Kelurahan, error)
	FindAll() ([]Kelurahan, error)
	DeleteByID(ID int) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) DeleteByID(ID int) error {
	var kelurahan Kelurahan
	err := r.db.Where("id = ?", ID).Delete(&kelurahan).Error

	if err != nil {
		return err
	}

	return nil
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(kelurahan Kelurahan) (Kelurahan, error) {
	err := r.db.Create(&kelurahan).Error
	if err != nil {
		return kelurahan, err
	}

	return kelurahan, nil
}

func (r *repository) FindByID(ID int) (Kelurahan, error) {
	var kelurahan Kelurahan

	err := r.db.Where("id = ?", ID).Find(&kelurahan).Error
	if err != nil {
		return kelurahan, err
	}

	return kelurahan, nil
}

func (r *repository) Update(kelurahan Kelurahan) (Kelurahan, error) {
	err := r.db.Save(&kelurahan).Error

	if err != nil {
		return kelurahan, err
	}

	return kelurahan, nil
}

func (r *repository) FindAll() ([]Kelurahan, error) {
	var kelurahans []Kelurahan

	err := r.db.Find(&kelurahans).Error
	if err != nil {
		return kelurahans, err
	}

	return kelurahans, nil
}
