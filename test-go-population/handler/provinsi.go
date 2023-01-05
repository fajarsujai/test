package handler

import (
	"net/http"
	"populationdatago/helper"

	"populationdatago/provinsi"

	"github.com/gin-gonic/gin"
)

type provinsiHandler struct {
	service provinsi.Service
}

func NewProvinsiHandler(service provinsi.Service) *provinsiHandler {
	return &provinsiHandler{service}
}

func (h *provinsiHandler) GetProvinsis(c *gin.Context) {
	var inputQuery helper.Pagination
	c.Bind(&inputQuery)

	provinsis, err := h.service.GetAllProvinsis(inputQuery)
	if err != nil {
		response := helper.APIResponse("Error to get provinsis", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of provinsis", http.StatusOK, "success", provinsis)
	c.JSON(http.StatusOK, response)
}

func (h *provinsiHandler) GetProvinsi(c *gin.Context) {
	var input provinsi.GetProvinsiDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of provinsi 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	provinsiDetail, err := h.service.GetProvinsiByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of provinsi 2", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Provinsi detail", http.StatusOK, "success", provinsi.FormatProvinsi(provinsiDetail))
	c.JSON(http.StatusOK, response)
}

func (h *provinsiHandler) CreateProvinsi(c *gin.Context) {
	var input provinsi.CreateProvinsiInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create provinsi", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProvinsi, err := h.service.CreateProvinsi(input)
	if err != nil {
		response := helper.APIResponse("Failed to create provinsi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create provinsi", http.StatusOK, "success", provinsi.FormatProvinsi(newProvinsi))
	c.JSON(http.StatusOK, response)
}

func (h *provinsiHandler) UpdateProvinsi(c *gin.Context) {
	var inputID provinsi.GetProvinsiDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update provinsi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData provinsi.CreateProvinsiInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update provinsi", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedProvinsi, err := h.service.UpdateProvinsi(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update provinsi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update provinsi", http.StatusOK, "success", provinsi.FormatProvinsi(updatedProvinsi))
	c.JSON(http.StatusOK, response)
}

func (h *provinsiHandler) Deleted(c *gin.Context) {
	var input provinsi.GetProvinsiDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to delete task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DeleteByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to delete task 2", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete data", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
