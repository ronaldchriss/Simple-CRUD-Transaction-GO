package handler

import (
	"net/http"
	"test_aac/helper"
	"test_aac/transaction"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service transaction.Service
}

func NewTransHandler(service transaction.Service) *TransactionHandler {
	return &TransactionHandler{service}
}

func (h *TransactionHandler) GetAllTransaction(c *gin.Context) {
	trans, err := h.service.GetAllTransaction()
	if err != nil {
		response := helper.JsonResponse("Error to Get Transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("List of Transaction", http.StatusOK, "success", transaction.FormatListTrans(trans))
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetTransactionbyNumber(c *gin.Context) {
	var input transaction.InputGetTransactionNumber

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	itemDetail, err := h.service.GetTransactionbyNumber(input.Number)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("Detail Item", http.StatusOK, "success", transaction.FormatDetailTransaction(itemDetail))
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetTransactionbyDate(c *gin.Context) {
	var input transaction.InputGetTransactionDate

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	itemDetail, err := h.service.GetTransactionbyDate(input.Date)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("Detail Item", http.StatusOK, "success", transaction.FormatListTrans(itemDetail))
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) TransactionCreate(c *gin.Context) {
	var input transaction.InputCreateTrans
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Create Transaction Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newItem, err := h.service.CreateTrans(input)
	if err != nil {
		response := helper.JsonResponse("Create Transaction Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	success, err := h.service.GetTransactionbyID(newItem.ID)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("Success Create Transaction", http.StatusOK, "success", transaction.FormatDetailTransaction(success))
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) DeleteTrans(c *gin.Context) {
	var input transaction.InputGetTransactionID

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transinfo, err := h.service.GetTransactionbyID(input.ID)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.DeleteTrans(input)
	if err != nil {
		response := helper.JsonResponse("Error to Delete Transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("Success Delete Transaction", http.StatusOK, "success", transaction.FormatDetailTransaction(transinfo))
	c.JSON(http.StatusOK, response)
}
