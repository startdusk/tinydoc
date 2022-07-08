package domain

import (
	"context"
	"fmt"
)

type DetailID int64

func (id DetailID) String() string {
	return fmt.Sprintf("%d", id)
}

func (id DetailID) Into() int64 {
	return int64(id)
}

// Doc detail
type DocDetail struct {
	ID        DetailID      `json:"id,string"`
	DocID     DocID         `json:"doc_id,string"`
	Titile    string        `json:"title"`
	Content   string        `json:"content"`
	Type      DocDetailType `json:"type"`
	Index     int           `json:"index"`
	CreatedAt int64         `json:"created_at,string"`
	UpdatedAt int64         `json:"updated_at,string"`
}

type DocDetailType int

const (
	// DocFile file
	DocFile DocDetailType = 1
	// DocFolder folder
	DocFolder DocDetailType = 2
)

// DocDetailUsecase represent the doc detail's usecases
type DocDetailUsecase interface {
	Fetch(ctx context.Context, docID int64) ([]DocDetail, error)
	GetByID(ctx context.Context, detailID int64) (DocDetail, error)
}

// DocDetailRepository represent the doc detail's repository contract
type DocDetailRepository interface {
	Fetch(ctx context.Context, docID DocID) ([]DocDetail, error)
	GetByID(ctx context.Context, detailID DetailID) (DocDetail, error)
}
