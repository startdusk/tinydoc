package usecase

import (
	"context"
	"server/internal/domain"
	"time"
)

type docUsecase struct {
	docRepo domain.DocRepository

	ctxTimeout time.Duration
}

// NewDocUsecase will create new an docUsecase object representation of domain.DocUsecase interface
func NewDocUsecase(d domain.DocRepository, timeout time.Duration) domain.DocUsecase {
	return &docUsecase{docRepo: d, ctxTimeout: timeout}
}

// Fetch implements domain.DocUsecase
func (d *docUsecase) Fetch(c context.Context, cursor int64, num int64) (res []domain.Doc, nextCursor int64, err error) {
	if num == 0 {
		num = 10
	}
	ctx, cancel := context.WithTimeout(c, d.ctxTimeout)
	defer cancel()

	return d.docRepo.Fetch(ctx, cursor, num)
}

// GetByID implements domain.DocUsecase
func (d *docUsecase) GetByID(c context.Context, id int64) (res domain.Doc, err error) {
	ctx, cancel := context.WithTimeout(c, d.ctxTimeout)
	defer cancel()
	return d.docRepo.GetByID(ctx, domain.DocID(id))
}
