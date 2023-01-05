package provinsi

import (
	"fmt"
	"math"
	"populationdatago/helper"

	"gorm.io/gorm"
)

type Repository interface {
	Save(provinsi Provinsi) (Provinsi, error)
	FindByID(ID int) (*Provinsi, error)
	Update(provinsi Provinsi) (Provinsi, error)
	FindAll(pagination helper.Pagination) (*helper.Pagination, error)
	DeleteByID(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func paginate(value interface{}, pagination *helper.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	fmt.Println(pagination)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func (r *repository) FindAll(pagination helper.Pagination) (*helper.Pagination, error) {
	var provinsis []Provinsi

	err := r.db.Scopes(paginate(provinsis, &pagination, r.db)).Find(&provinsis).Error
	pagination.Rows = provinsis
	if err != nil {
		return &pagination, err
	}

	return &pagination, nil
}

func (r *repository) DeleteByID(ID int) error {
	var provinsis Provinsi
	err := r.db.Where("id = ?", ID).Delete(&provinsis).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Save(provinsi Provinsi) (Provinsi, error) {
	err := r.db.Create(&provinsi).Error
	if err != nil {
		return provinsi, err
	}

	return provinsi, nil
}

func (r *repository) FindByID(ID int) (*Provinsi, error) {
	var provinsis Provinsi

	err := r.db.Where("id = ?", ID).Find(&provinsis).Error
	if err != nil {
		return &provinsis, err
	}

	return &provinsis, nil
}

func (r *repository) Update(provinsis Provinsi) (Provinsi, error) {
	err := r.db.Save(&provinsis).Error

	if err != nil {
		return provinsis, err
	}

	return provinsis, nil
}

// func (r *repository) FindAll() ([]Provinsi, error) {
// 	var provinsis []Provinsi

// 	err := r.db.Find(&provinsis).Error
// 	if err != nil {
// 		return &pagination, err
// 	}

// 	return &pagination, nil
// }

// func (cg *CategoryGorm) List(pagination pkg.Pagination) (*pkg.Pagination, error) {
// 	var categories []*Category

// 	cg.db.Scopes(paginate(categories, &pagination, cg.db)).Find(&categories)
// 	pagination.Rows = categories

// 	return &pagination, nil
// }
// return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
// err := r.db.Find(&provinsis).Error
