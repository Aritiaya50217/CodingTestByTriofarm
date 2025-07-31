package http

import (
	"net/http"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/usecase"
	"github.com/gin-gonic/gin"
)

type MedicineHandler struct {
	usecase usecase.MedicineUsecase
}

func NewMedicineHandler(r *gin.Engine, u usecase.MedicineUsecase, api *gin.RouterGroup) {
	handler := &MedicineHandler{usecase: u}
	api.POST("/medicines", handler.CreateMedicine)
}

func (h *MedicineHandler) CreateMedicine(c *gin.Context) {
	var medicine domain.Medicines
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	maxIndex, err := h.usecase.GetMaxIndex()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get max index"})
		return
	}
	medicine.Index = maxIndex + 1

	if err := h.usecase.CreateMedicine(&medicine); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.PreloadTopic(&medicine); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load topic"})
		return
	}

	c.JSON(http.StatusCreated, medicine)
}
