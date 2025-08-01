package http

import (
	"net/http"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/usecase"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/utils"
	"github.com/gin-gonic/gin"
)

type MicroorganismHandler struct {
	usecase usecase.MicroorganismUsecase
}

func NewMicroorganismHandler(r *gin.Engine, u usecase.MicroorganismUsecase, api *gin.RouterGroup) {
	handler := &MicroorganismHandler{usecase: u}
	api.POST("/microorganism", handler.CreateMicroorganism)
}

func (h *MicroorganismHandler) CreateMicroorganism(c *gin.Context) {
	var microorganism domain.Microorganisms
	if err := c.ShouldBindJSON(&microorganism); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	maxIndex, err := h.usecase.GetMaxIndex()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get max index"})
		return
	}
	microorganism.Index = maxIndex + 1

	if err := h.usecase.CreateMicroorganism(&microorganism); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := utils.Reponse{
		ID:      microorganism.ID,
		Name:    microorganism.Name,
		TopicID: microorganism.TopicID,
		Index:   microorganism.Index,
	}
	c.JSON(http.StatusCreated, result)
}
