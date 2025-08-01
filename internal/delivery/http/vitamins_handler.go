package http

import (
	"net/http"
	"strconv"

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
	api.GET("/vitamins", handler.GetAllVitamin)
	api.POST("/vitamin/:id", handler.UpdateVitamin)
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

func (h *VitaminHandler) GetAllVitamin(c *gin.Context) {
	vitamins, err := h.usecase.GetAllVitamin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get medicines"})
		return
	}
	var result []response.Reponse
	for _, vitamin := range vitamins {
		result = append(result, response.Reponse{
			ID:      vitamin.ID,
			Name:    vitamin.Name,
			TopicID: vitamin.TopicID,
			Index:   vitamin.Index,
		})
	}
	c.JSON(http.StatusOK, result)

}

func (h *VitaminHandler) UpdateVitamin(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON เข้ากับ struct
	var vitamin domain.Vitamins
	if err := c.ShouldBindJSON(&vitamin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.usecase.GetVitaminByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	vitamin.ID = id

	if err := h.usecase.UpdateVitamin(&vitamin); err != nil {
		if err.Error() == "name already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vitamin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vitamin updated successfully"})
}
