package handler

import (
	"net/http"
	"populationdatago/helper"
	"populationdatago/kelurahan"

	"github.com/gin-gonic/gin"
)

type kelurahanHandler struct {
	service kelurahan.Service
}

func NewKelurahanHandler(service kelurahan.Service) *kelurahanHandler {
	return &kelurahanHandler{service}
}

func (h *kelurahanHandler) GetKelurahans(c *gin.Context) {

	kelurahans, err := h.service.GetAllKelurahans()
	if err != nil {
		response := helper.APIResponse("Error to get kelurahans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of kelurahans", http.StatusOK, "success", kelurahan.FormatKelurahans(kelurahans))
	c.JSON(http.StatusOK, response)
}

func (h *kelurahanHandler) GetKelurahan(c *gin.Context) {
	var input kelurahan.GetKelurahanDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of kelurahan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	kelurahanDetail, err := h.service.GetKelurahanByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of kelurahan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Kelurahan detail", http.StatusOK, "success", kelurahan.FormatKelurahan(kelurahanDetail))
	c.JSON(http.StatusOK, response)
}

func (h *kelurahanHandler) CreateKelurahan(c *gin.Context) {
	var input kelurahan.CreateKelurahanInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create kelurahan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKelurahan, err := h.service.CreateKelurahan(input)
	if err != nil {
		response := helper.APIResponse("Failed to create kelurahan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create kelurahan", http.StatusOK, "success", kelurahan.FormatKelurahan(newKelurahan))
	c.JSON(http.StatusOK, response)
}

func (h *kelurahanHandler) UpdateKelurahan(c *gin.Context) {
	var inputID kelurahan.GetKelurahanDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update kelurahan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData kelurahan.CreateKelurahanInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update kelurahan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedKelurahan, err := h.service.UpdateKelurahan(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update kelurahan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update kelurahan", http.StatusOK, "success", kelurahan.FormatKelurahan(updatedKelurahan))
	c.JSON(http.StatusOK, response)
}

func (h *kelurahanHandler) Deleted(c *gin.Context) {
	var input kelurahan.GetKelurahanDetailInput

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
