// Package docgrp maintains the group of handlers for doc access.
package docgrp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handlers manages the set of doc endpoints.
type Handlers struct {
}

// Query returns a list of docs with paging.
func (h *Handlers) Query(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// QueryByID returns a doc by its ID.
func (h *Handlers) QueryByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
