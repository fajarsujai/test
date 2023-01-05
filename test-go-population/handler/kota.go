package handler

import (
	"net/http"
	"populationdatago/helper"
	"populationdatago/kota"

	"github.com/gin-gonic/gin"
)

type kotaHandler struct {
	service kota.Service
}

func NewKotaHandler(service kota.Service) *kotaHandler {
	return &kotaHandler{service}
}

func (h *kotaHandler) GetKotas(c *gin.Context) {

	kotas, err := h.service.GetAllKotas()
	if err != nil {
		response := helper.APIResponse("Error to get kotas", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of kotas", http.StatusOK, "success", kota.FormatKotas(kotas))
	c.JSON(http.StatusOK, response)
}

func (h *kotaHandler) GetKota(c *gin.Context) {
	var input kota.GetKotaDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of kota", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	kotaDetail, err := h.service.GetKotaByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of kota", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Kota detail", http.StatusOK, "success", kota.FormatKota(kotaDetail))
	c.JSON(http.StatusOK, response)
}

func (h *kotaHandler) CreateKota(c *gin.Context) {
	var input kota.CreateKotaInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create kota", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKota, err := h.service.CreateKota(input)
	if err != nil {
		response := helper.APIResponse("Failed to create kota", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create kota", http.StatusOK, "success", kota.FormatKota(newKota))
	c.JSON(http.StatusOK, response)
}

func (h *kotaHandler) UpdateKota(c *gin.Context) {
	var inputID kota.GetKotaDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update kota", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData kota.CreateKotaInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update kota", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedKota, err := h.service.UpdateKota(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update kota", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update kota", http.StatusOK, "success", kota.FormatKota(updatedKota))
	c.JSON(http.StatusOK, response)
}

func (h *kotaHandler) Deleted(c *gin.Context) {
	var input kota.GetKotaDetailInput

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
