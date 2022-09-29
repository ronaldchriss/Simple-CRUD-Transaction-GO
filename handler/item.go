package handler

import (
	"net/http"
	"test_aac/helper"
	"test_aac/item"
	"test_aac/user"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	service item.Service
}

func NewItemHandler(service item.Service) *ItemHandler {
	return &ItemHandler{service}
}

func (h *ItemHandler) GetItem(c *gin.Context) {
	c.Query("user_id")

	items, err := h.service.GetItems()
	if err != nil {
		response := helper.JsonResponse("Error to Get Items", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("List of Items", http.StatusOK, "success", item.FormatItems(items))
	c.JSON(http.StatusOK, response)
}

func (h *ItemHandler) GetDetail(c *gin.Context) {
	var input item.GetItemDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	itemDetail, err := h.service.GetItemByID(input)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("Detail Item", http.StatusOK, "success", item.FormatItemDetail(itemDetail))
	c.JSON(http.StatusOK, response)
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	var input item.GetItemDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	iteminfo, err := h.service.GetItemByID(input)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	h.service.DeleteItem(input)

	response := helper.JsonResponse("Success Delete Item", http.StatusOK, "success", item.FormatItemDetail(iteminfo))
	c.JSON(http.StatusOK, response)
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var input item.CreateItemInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Create Item Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	code := c.MustGet("codeUser").(user.User)

	input.User = code

	newItem, err := h.service.CreateItem(input)
	if err != nil {
		response := helper.JsonResponse("Create Item Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := item.FormatItem(newItem)
	response := helper.JsonResponse("Success to Create Item", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	var inputID item.GetItemDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.JsonResponse("Error to Update Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input item.CreateItemInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Error to Update Item", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	code := c.MustGet("codeUser").(user.User)
	input.User = code

	updateItem, err := h.service.Update(inputID, input)
	if err != nil {
		response := helper.JsonResponse("Error to Update Item", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := item.FormatItem(updateItem)
	response := helper.JsonResponse("Success to Update Item", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
