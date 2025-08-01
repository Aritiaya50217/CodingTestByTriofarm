package http

import (
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/usecase"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/utils"
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
	api.DELETE("/vitamin/:id", handler.DeleteVitamin)
	api.POST("/swap/vitamins", handler.SwapVitamins)
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

	result := utils.Reponse{
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
	var result []utils.Reponse
	for _, vitamin := range vitamins {
		result = append(result, utils.Reponse{
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

func (h *VitaminHandler) DeleteVitamin(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	_, err = h.usecase.GetVitaminByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.usecase.DeleteVitamin(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vitamin deleted successfully"})
}

func (h *VitaminHandler) SwapVitamins(c *gin.Context) {
	var vitamins []domain.Vitamins

	if err := c.ShouldBindJSON(&vitamins); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.SwapVitamins(vitamins)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vitamins"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Vitamins updated successfully"})
}
