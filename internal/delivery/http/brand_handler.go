package http

import (
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/usecase"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/utils"
	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	usecase usecase.BrandUsecase
}

func NewBrandHandler(r *gin.Engine, u usecase.BrandUsecase, api *gin.RouterGroup) {
	handler := &BrandHandler{usecase: u}
	api.POST("/brand", handler.CreateBrand)
	api.GET("/brands", handler.GetAllBrand)
	api.POST("/brand/:id", handler.UpdateBrand)
	api.DELETE("/brand/:id", handler.DeleteBrand)
	api.POST("/swap/brands", handler.SwapBrands)
}

func (h *BrandHandler) CreateBrand(c *gin.Context) {
	var brand domain.Brands
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	maxIndex, err := h.usecase.GetMaxIndex()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get max index"})
		return
	}
	brand.Index = maxIndex + 1

	if err := h.usecase.CreateBrand(&brand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := utils.Reponse{
		ID:      brand.ID,
		Name:    brand.Name,
		TopicID: brand.TopicID,
		Index:   brand.Index,
	}
	c.JSON(http.StatusCreated, result)
}

func (h *BrandHandler) GetAllBrand(c *gin.Context) {
	brands, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get microorganisms"})
		return
	}

	var result []utils.Reponse
	for _, brand := range brands {
		result = append(result, utils.Reponse{
			ID:      brand.ID,
			Name:    brand.Name,
			TopicID: brand.TopicID,
			Index:   brand.Index,
		})
	}
	c.JSON(http.StatusOK, result)
}

func (h *BrandHandler) UpdateBrand(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON เข้ากับ struct
	var brand domain.Brands
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.usecase.GetBrandByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	brand.ID = id

	if err := h.usecase.UpdateBrand(&brand); err != nil {

		if err.Error() == "name already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Brand"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand updated successfully"})
}

func (h *BrandHandler) DeleteBrand(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	_, err = h.usecase.GetBrandByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.usecase.DeleteBrand(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand deleted successfully"})
}

func (h *BrandHandler) SwapBrands(c *gin.Context) {
	var brands []domain.Brands

	if err := c.ShouldBindJSON(&brands); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.SwapBrand(brands)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Brand"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Brands updated successfully"})
}
