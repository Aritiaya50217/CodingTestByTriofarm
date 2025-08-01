package http

import (
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TopicHandler struct {
	topicUsecase usecase.TopicUsecase
}

func NewTopicHandler(r *gin.Engine, topicUsecase usecase.TopicUsecase, api *gin.RouterGroup) {
	handler := &TopicHandler{topicUsecase: topicUsecase}
	api.GET("/topics", handler.ListTopics)
	api.POST("/topic", handler.CreateTopic)
	api.POST("/topic/:id", handler.UpdateTopic)
	api.DELETE("/topic/:id", handler.DeleteTopic)
}

func (h *TopicHandler) ListTopics(c *gin.Context) {
	topics, err := h.topicUsecase.ListTopics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, topics)
}

func (h *TopicHandler) CreateTopic(c *gin.Context) {
	var topic domain.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.topicUsecase.CreateTopic(&topic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, topic)
}

func (h *TopicHandler) UpdateTopic(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON เข้ากับ struct
	var topic domain.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.topicUsecase.GetTopicByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	topic.ID = id

	if err := h.topicUsecase.UpdateTopic(&topic); err != nil {

		if err.Error() == "name already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Topic"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Topic updated successfully"})
}

func (h *TopicHandler) DeleteTopic(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.topicUsecase.DeleteTopic(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete topic"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Topic deleted successfully"})

}
