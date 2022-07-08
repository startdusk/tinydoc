package http

import (
	"net/http"
	"server/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type docHandler struct {
	usecase domain.DocUsecase
}

// NewDocHandler will initialize the docs / resources endpoint
func NewDocHandler(g *gin.Engine, du domain.DocUsecase) {
	handler := &docHandler{
		usecase: du,
	}

	g.GET("/docs", handler.Docs)
	g.POST("/docs", handler.Store)
}

func (h *docHandler) Docs(c *gin.Context) {
	num, _ := strconv.Atoi(c.Param("num"))
	cursor, _ := strconv.Atoi(c.Param("cursor"))
	ctx := c.Request.Context()

	docs, nextCursor, err := h.usecase.Fetch(ctx, int64(cursor), int64(num))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "",
		})
		return
	}

	c.Header("X-Cursor", strconv.FormatInt(nextCursor, 10))
	c.JSON(http.StatusOK, docs)
}

// Store will store the doc by given request body
func (h *docHandler) Store(c *gin.Context) {

}