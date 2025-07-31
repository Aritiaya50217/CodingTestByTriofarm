package http

import (
	"net/http"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TopicHandler struct {
	topicUsecase *usecase.TopicUsecase
}

func NewTopicHandler(r *gin.Engine, topicUsecase *usecase.TopicUsecase) {
	handler := &TopicHandler{topicUsecase: topicUsecase}
	r.GET("/users", handler.ListTopics)
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
