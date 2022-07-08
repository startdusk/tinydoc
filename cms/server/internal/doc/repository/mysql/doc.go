package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/domain"
)

type docRepository struct {
	db *sql.DB
}

// NewDocRepository will create an object that represent the doc.Repository interface
func NewDocRepository(db *sql.DB) domain.DocRepository {
	return &docRepository{db: db}
}

// Delete implements domain.DocRepository
func (d *docRepository) Delete(ctx context.Context, id int64) (err error) {
	const query = `
		UPDATE doc SET deleted_at = ? WHERE id = ?
	`
	stmt, err := d.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}

// Fetch implements domain.DocRepository
func (d *docRepository) Fetch(ctx context.Context, cursor int64, num int64) (res []domain.Doc, nextCursor int64, err error) {
	const query = `
		SELECT id, title, description, status, version, internal_version, index, created_at, updated_at
			FROM doc WHERE created_at > ? AND status <> 3 AND deleted_at IS NOT NULL ORDER BY created_at LIMIT ?
	`
	res, err = d.fetch(ctx, query, cursor, num)
	if err != nil {
		return nil, -1, err
	}

	if len(res) == int(num) {
		nextCursor = res[len(res)-1].CreatedAt
	}
	return
}

// GetByID implements domain.DocRepository
func (d *docRepository) GetByID(ctx context.Context, id int64) (res domain.Doc, err error) {
	const query = `
		SELECT id, title, description, status, version, internal_version, index, created_at, updated_at
			FROM doc WHERE id = ? AND status <> 3 AND deleted_at IS NOT NULL
	`
	list, err := d.fetch(ctx, query, id)
	if err != nil {
		return domain.NilDoc, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

// GetByTitle implements domain.DocRepository
func (d *docRepository) GetByTitle(ctx context.Context, title string) (res domain.Doc, err error) {
	const query = `
		SELECT id, title, description, status, version, internal_version, index, created_at, updated_at
			FROM doc WHERE title = ? AND status <> 3 AND deleted_at IS NOT NULL
	`
	list, err := d.fetch(ctx, query, title)
	if err != nil {
		return domain.NilDoc, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

// Store implements domain.DocRepository
func (d *docRepository) Store(ctx context.Context, doc *domain.Doc) (err error) {
	const query = `
		INSERT INTO doc(title, description, status, version, index)
			VALUES(?, ?, ?, ?, ?)
	`
	stmt, err := d.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, doc.Titile, doc.Description, doc.Status, doc.Version, doc.Index)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	doc.ID = lastID
	return
}

// Update implements domain.DocRepository
func (d *docRepository) Update(ctx context.Context, doc *domain.Doc) (err error) {
	oldDoc, err := d.GetByID(ctx, doc.ID)
	if err != nil {
		return err
	}
	if doc.Titile == "" {
		doc.Titile = oldDoc.Titile
	}

	const query = `
		UPDATE doc SET title=?, description=?, status=?, version=?, index=?
			WHERE id = ?
	`
	stmt, err := d.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, doc.Titile, doc.Description, doc.Status, doc.Version, doc.Index)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}
	return
}

func (d *docRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Doc, err error) {
	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result = make([]domain.Doc, 0)
	for rows.Next() {
		d := domain.Doc{}
		err = rows.Scan(
			&d.ID,
			&d.Titile,
			&d.Description,
			&d.Status,
			&d.Version,
			&d.InternalVersion,
			&d.Index,
			&d.CreatedAt,
			&d.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
