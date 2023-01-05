package handler

import (
	"net/http"
	"populationdatago/helper"
	"populationdatago/memberdetail"

	"github.com/gin-gonic/gin"
)

type memberdetailHandler struct {
	service memberdetail.Service
}

func NewMemberDetailHandler(service memberdetail.Service) *memberdetailHandler {
	return &memberdetailHandler{service}
}

func (h *memberdetailHandler) GetMemberDetails(c *gin.Context) {

	memberdetails, err := h.service.GetAllMemberDetails()
	if err != nil {
		response := helper.APIResponse("Error to get memberdetails", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of memberdetails", http.StatusOK, "success", memberdetail.FormatMemberDetails(memberdetails))
	c.JSON(http.StatusOK, response)
}

func (h *memberdetailHandler) GetMemberDetail(c *gin.Context) {
	var input memberdetail.GetMemberDetailDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of memberdetail 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	memberdetailDetail, err := h.service.GetMemberDetailByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of memberdetail 2", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("memberdetail detail", http.StatusOK, "success", memberdetail.FormatMemberDetail(memberdetailDetail))
	c.JSON(http.StatusOK, response)
}

func (h *memberdetailHandler) CreateMemberDetail(c *gin.Context) {
	var input memberdetail.CreateMemberDetailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create memberdetail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMemberDetail, err := h.service.CreateMemberDetail(input)
	if err != nil {
		response := helper.APIResponse("Failed to create memberdetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create memberdetail", http.StatusOK, "success", memberdetail.FormatMemberDetail(newMemberDetail))
	c.JSON(http.StatusOK, response)
}

func (h *memberdetailHandler) UpdateMemberDetail(c *gin.Context) {
	var inputID memberdetail.GetMemberDetailDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update memberdetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData memberdetail.CreateMemberDetailInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update memberdetail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedMemberDetail, err := h.service.UpdateMemberDetail(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update memberdetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update memberdetail", http.StatusOK, "success", memberdetail.FormatMemberDetail(updatedMemberDetail))
	c.JSON(http.StatusOK, response)
}

func (h *memberdetailHandler) Deleted(c *gin.Context) {
	var input memberdetail.GetMemberDetailDetailInput

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
