package provinsi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var provinsiRepository = &RepositoryMock{Mock: mock.Mock{}}
var provinsiService = NewService(provinsiRepository)

func TestService_GetByIDNotFound(t *testing.T) {
	var input GetProvinsiDetailInput
	input.ID = 1
	//program mock
	provinsiRepository.Mock.On("FindByID", 1).Return(nil)

	category, err := provinsiService.GetProvinsiByID(input)
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestService_GetByIDSuccess(t *testing.T) {
	var input GetProvinsiDetailInput
	input.ID = 2

	provinsi := Provinsi{
		ProvinsiName: "Jakarta",
		ProvinsiCode: "JKT48",
		ID:           2,
	}

	// Programm mock
	provinsiRepository.Mock.On("FindByID", 2).Return(provinsi)

	result, err := provinsiService.GetProvinsiByID(input)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, provinsi.ProvinsiCode, result.ProvinsiCode)
	assert.Equal(t, provinsi.ProvinsiName, result.ProvinsiName)
	assert.Equal(t, provinsi.ID, result.ID)

}
