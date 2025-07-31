package http

import (
	"net/http"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TopicHandler struct {
	topicUsecase usecase.TopicUsecase
}

func NewTopicHandler(r *gin.Engine, topicUsecase usecase.TopicUsecase) {
	handler := &TopicHandler{topicUsecase: topicUsecase}
	r.GET("/topics", handler.ListTopics)
	r.POST("/topic", handler.CreateTopic)
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
