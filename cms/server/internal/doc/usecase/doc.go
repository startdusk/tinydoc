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

// Delete implements domain.DocUsecase
func (d *docUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, d.ctxTimeout)
	defer cancel()
	existed, err := d.docRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existed == domain.NilDoc {
		return domain.ErrNotFound
	}
	return d.docRepo.Delete(ctx, id)
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
	return d.docRepo.GetByID(ctx, id)
}

// GetByTitle implements domain.DocUsecase
func (d *docUsecase) GetByTitle(c context.Context, title string) (res domain.Doc, err error) {
	ctx, cancel := context.WithTimeout(c, d.ctxTimeout)
	defer cancel()
	return d.docRepo.GetByTitle(ctx, title)
}

// Store implements domain.DocUsecase
func (d *docUsecase) Store(c context.Context, doc *domain.Doc) (err error) {
	ctx, cancel := context.WithTimeout(c, d.ctxTimeout)
	defer cancel()
	existed, _ := d.docRepo.GetByTitle(ctx, doc.Titile)
	if existed != domain.NilDoc {
		return domain.ErrConflict
	}
	return d.docRepo.Store(ctx, doc)
}

// Update implements domain.DocUsecase
func (d *docUsecase) Update(c context.Context, doc *domain.Doc) (err error) {
	ctx, cancel := context.WithTimeout(c, d.ctxTimeout)
	defer cancel()

	return d.docRepo.Update(ctx, doc)
}
