package http

import (
	"net/http"
	"strconv"

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
	api.GET("/microorganisms", handler.GetAllMicroorganism)
	api.POST("/microorganism/:id", handler.UpdateMicroorganism)
	api.DELETE("/microorganism/:id", handler.DeleteMicroorganism)
	api.POST("/swap/microorganisms", handler.SwapMicroorganisms)
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

func (h *MicroorganismHandler) GetAllMicroorganism(c *gin.Context) {
	microorganisms, err := h.usecase.GetAllMicroorganism()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get microorganisms"})
		return
	}

	var result []utils.Reponse
	for _, microorganism := range microorganisms {
		result = append(result, utils.Reponse{
			ID:      microorganism.ID,
			Name:    microorganism.Name,
			TopicID: microorganism.TopicID,
			Index:   microorganism.Index,
		})
	}
	c.JSON(http.StatusOK, result)
}

func (h *MicroorganismHandler) UpdateMicroorganism(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON เข้ากับ struct
	var microorganism domain.Microorganisms
	if err := c.ShouldBindJSON(&microorganism); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.usecase.GetMicroorganismByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	microorganism.ID = id

	if err := h.usecase.UpdateMicroorganism(&microorganism); err != nil {

		if err.Error() == "name already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update microorganism"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Microorganism updated successfully"})
}

func (h *MicroorganismHandler) DeleteMicroorganism(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	_, err = h.usecase.GetMicroorganismByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.usecase.DeleteMicroorganism(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Microorganism deleted successfully"})
}

func (h *MicroorganismHandler) SwapMicroorganisms(c *gin.Context) {
	var microorganisms []domain.Microorganisms

	if err := c.ShouldBindJSON(&microorganisms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.SwapMicroorganisms(microorganisms)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Microorganisms"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Microorganisms updated successfully"})
}
