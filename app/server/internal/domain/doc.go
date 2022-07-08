package domain

import (
	"context"
	"fmt"
)

// NilDoc is an empty Doc
var NilDoc Doc

type DocID int64

func (id DocID) String() string {
	return fmt.Sprintf("%d", id)
}

func (id DocID) Into() int64 {
	return int64(id)
}

// Doc info
type Doc struct {
	ID              DocID     `json:"id,string"`
	Titile          string    `json:"title"`
	Description     string    `json:"description"`
	Status          DocStatus `json:"status"`
	Version         float64   `json:"version"`
	InternalVersion int       `json:"internal_version"`
	Index           int       `json:"index"`
	CreatedAt       int64     `json:"created_at,string"`
	UpdatedAt       int64     `json:"updated_at,string"`
}

// DocStatus doc status
type DocStatus int

const (
	// Unpublished doc unpublished
	Unpublished DocStatus = 1
	// Published doc published
	Published DocStatus = 2
	// Disabled doc disabled
	Disabled DocStatus = 3
)

// DocUsecase represent the doc's usecases
type DocUsecase interface {
	Fetch(ctx context.Context, cursor int64, num int64) ([]Doc, int64, error)
	GetByID(ctx context.Context, id int64) (Doc, error)
}

// DocRepository represent the doc's repository contract
type DocRepository interface {
	Fetch(ctx context.Context, cursor int64, num int64) ([]Doc, int64, error)
	GetByID(ctx context.Context, id DocID) (Doc, error)
}
