package handler

import (
	"net/http"
	"populationdatago/helper"
	"populationdatago/kecamatan"

	"github.com/gin-gonic/gin"
)

type kecamatanHandler struct {
	service kecamatan.Service
}

func NewKecamatanHandler(service kecamatan.Service) *kecamatanHandler {
	return &kecamatanHandler{service}
}

func (h *kecamatanHandler) GetKecamatans(c *gin.Context) {

	kecamatans, err := h.service.GetAllKecamatans()
	if err != nil {
		response := helper.APIResponse("Error to get kecamatans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of kecamatans", http.StatusOK, "success", kecamatan.FormatKecamatans(kecamatans))
	c.JSON(http.StatusOK, response)
}

func (h *kecamatanHandler) GetKecamatan(c *gin.Context) {
	var input kecamatan.GetKecamatanDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of kecamatan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	kecamatanDetail, err := h.service.GetKecamatanByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of kecamatan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Kecamatan detail", http.StatusOK, "success", kecamatan.FormatKecamatan(kecamatanDetail))
	c.JSON(http.StatusOK, response)
}

func (h *kecamatanHandler) CreateKecamatan(c *gin.Context) {
	var input kecamatan.CreateKecamatanInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create kota", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKecamatan, err := h.service.CreateKecamatan(input)
	if err != nil {
		response := helper.APIResponse("Failed to create kota", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create kota", http.StatusOK, "success", kecamatan.FormatKecamatan(newKecamatan))
	c.JSON(http.StatusOK, response)
}

func (h *kecamatanHandler) UpdateKecamatan(c *gin.Context) {
	var inputID kecamatan.GetKecamatanDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update kecamatan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData kecamatan.CreateKecamatanInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update kota", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedKecamatan, err := h.service.UpdateKecamatan(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update kecamatan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update kecamatan", http.StatusOK, "success", kecamatan.FormatKecamatan(updatedKecamatan))
	c.JSON(http.StatusOK, response)
}

func (h *kecamatanHandler) Deleted(c *gin.Context) {
	var input kecamatan.GetKecamatanDetailInput

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
