package score

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"advance-go/internal/config"
)

type request struct {
	Score int `json:"score" binding:"required"`
}

type grade struct {
	Score int    `json:"score" binding:"required"`
	Grade string `json:"grade" binding:"required"`
}

type score int

func (s score) Calculate() grade {
	si := int(s)
	switch {
	case si > 90:
		return grade{si, "A"}
	case si > 80:
		return grade{si, "B"}
	case si > 70:
		return grade{si, "C"}
	case si > 60:
		return grade{si, "D"}
	default:
		return grade{si, "F"}
	}
}

type Handler struct {
	conf *config.Config
}

func NewHandler(conf *config.Config) *Handler {
	return &Handler{conf: conf}
}

func (h *Handler) GetScore(c *gin.Context) {
	var json request

	c.Bind(&json)
	s := score(json.Score)
	c.JSON(http.StatusOK, s.Calculate())
}
