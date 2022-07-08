package http

import (
	"fmt"
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
	g.GET("/docs/:doc_id", handler.DocDetails)
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

	c.Header("X-Cursor", fmt.Sprintf("%d", nextCursor))
	c.JSON(http.StatusOK, docs)
}

func (h *docHandler) DocDetails(c *gin.Context) {
	
}
