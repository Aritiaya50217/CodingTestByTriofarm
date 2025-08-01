package http

import (
	"net/http"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/response"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/usecase"
	"github.com/gin-gonic/gin"
)

type VitaminHandler struct {
	usecase usecase.VitaminUsecase
}

func NewVitaminHandler(r *gin.Engine, u usecase.VitaminUsecase, api *gin.RouterGroup) {
	handler := &VitaminHandler{usecase: u}
	api.POST("/vitamin", handler.CreateVitamin)
}

func (h *VitaminHandler) CreateVitamin(c *gin.Context) {
	var vitamin domain.Vitamins

	if err := c.ShouldBindJSON(&vitamin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	maxIndex, err := h.usecase.GetMaxIndex()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get max index"})
		return
	}

	vitamin.Index = maxIndex + 1

	if err := h.usecase.CreateVitamin(&vitamin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := response.Reponse{
		ID:      vitamin.ID,
		Name:    vitamin.Name,
		TopicID: vitamin.TopicID,
		Index:   vitamin.Index,
	}

	c.JSON(http.StatusCreated, result)
}
