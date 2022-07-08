package domain

import "context"

// NilDoc is an empty Doc
var NilDoc Doc

// Doc info
type Doc struct {
	ID              int64     `json:"id,string"`
	Titile          string    `json:"title"`
	Description     string    `json:"description"`
	Status          DocStatus `json:"status"`
	Version         float64   `json:"version"`
	InternalVersion int       `json:"internal_version"`
	Index           int       `json:"index"`
	CreatedAt       int64     `json:"created_at,string"`
	UpdatedAt       int64     `json:"updated_at,string"`
	DeletedAt       int64     `json:"deleted_at,string"`
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
	Update(ctx context.Context, doc *Doc) error
	GetByTitle(ctx context.Context, title string) (Doc, error)
	Store(ctx context.Context, doc *Doc) error
	Delete(ctx context.Context, id int64) error
}

// DocRepository represent the doc's repository contract
type DocRepository interface {
	Fetch(ctx context.Context, cursor int64, num int64) ([]Doc, int64, error)
	GetByID(ctx context.Context, id int64) (Doc, error)
	Update(ctx context.Context, doc *Doc) error
	GetByTitle(ctx context.Context, title string) (Doc, error)
	Store(ctx context.Context, doc *Doc) error
	Delete(ctx context.Context, id int64) error
}
