package handler

import (
	"go-gin/model"
	"go-gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Item struct {
	service service.IItem
}

func NewItem(service service.IItem) *Item {
	return &Item{
		service: service,
	}
}

func (h *Item) CreateItem(c *gin.Context) {
	var item model.Item
	c.ShouldBind(&item)
	if err := h.service.CreateItem(c, &item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Item) GetItems(c *gin.Context) {
	item, err := h.service.GetItems(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Item) UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	if err := h.service.UpdateItem(c, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Item) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	item, err := h.service.DeleteItem(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
		"data":    item,
	})
}
