package http

import (
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/response"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/usecase"
	"github.com/gin-gonic/gin"
)

type MedicineHandler struct {
	usecase usecase.MedicineUsecase
}

func NewMedicineHandler(r *gin.Engine, u usecase.MedicineUsecase, api *gin.RouterGroup) {
	handler := &MedicineHandler{usecase: u}
	api.POST("/medicines", handler.CreateMedicine)
	api.GET("/medicines", handler.GetAllMedicine)
	api.POST("/medicines/:id", handler.UpdateMedicine)
	api.DELETE("/medicines/:id", handler.DeleteMedicine)
	api.POST("/swap/medicines", handler.SwapMedicines)
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
	result := response.Reponse{
		ID:      medicine.ID,
		Name:    medicine.Name,
		TopicID: medicine.TopicID,
		Index:   medicine.Index,
	}
	c.JSON(http.StatusCreated, result)
}

func (h *MedicineHandler) GetAllMedicine(c *gin.Context) {
	medicines, err := h.usecase.GetAllMedicines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get medicines"})
		return
	}
	var result []response.Reponse
	for _, medicine := range medicines {
		result = append(result, response.Reponse{
			ID:      medicine.ID,
			Name:    medicine.Name,
			TopicID: medicine.TopicID,
			Index:   medicine.Index,
		})
	}
	c.JSON(http.StatusOK, result)
}

func (h *MedicineHandler) UpdateMedicine(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON เข้ากับ struct
	var medicine domain.Medicines
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.usecase.GetMedicineByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	medicine.ID = id

	if err := h.usecase.UpdateMedicine(&medicine); err != nil {

		if err.Error() == "name already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update medicine"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Medicine updated successfully"})
}

func (h *MedicineHandler) GetMedicineByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	medicine, err := h.usecase.GetMedicineByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medicine not found"})
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func (h *MedicineHandler) DeleteMedicine(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	_, err = h.usecase.GetMedicineByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.usecase.DeleteMedicine(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Medicine deleted successfully"})
}

func (h *MedicineHandler) SwapMedicines(c *gin.Context) {
	var medicines []domain.Medicines

	if err := c.ShouldBindJSON(&medicines); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.SwapMedicines(medicines)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update medicines"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Medicines updated successfully"})
}
