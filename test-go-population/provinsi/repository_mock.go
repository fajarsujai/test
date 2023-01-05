package provinsi

import (
	"errors"
	"populationdatago/helper"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (repository *RepositoryMock) FindByID(ID int) (*Provinsi, error) {
	arguments := repository.Mock.Called(ID)

	if arguments.Get(0) == nil {
		return nil, errors.New("Error")
	} else {
		provinsi := arguments.Get(0).(Provinsi)
		return &provinsi, nil
	}
}

func (repository *RepositoryMock) DeleteByID(ID int) error {
	panic("implement me")
}

func (repository *RepositoryMock) FindAll(pagination helper.Pagination) (*helper.Pagination, error) {
	panic("implement me")
}

func (repository *RepositoryMock) Save(provinsi Provinsi) (Provinsi, error) {
	panic("implement me")
}

func (repository *RepositoryMock) Update(provinsis Provinsi) (Provinsi, error) {
	panic("implement me")
}
