package repository

import "populationdatago/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
