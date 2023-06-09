package postgresql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type bookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) *bookRepo {
	return &bookRepo{
		db: db,
	}
}

func (r *bookRepo) Create(ctx context.Context, req *models.CreateBook) (string, error) {

	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO book(
			id, 
			name, 
			price,
			count,
			came_price,
			profit_status,
			profit,
			sell_price, 
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now())
	`

	_, err := r.db.Exec(ctx, query,
		id.String(),
		req.Name,
		req.Price,
		req.Count,
		req.Came_price,
		req.Profit_status,
		req.Profit,
		req.Sell_price,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *bookRepo) GetByID(ctx context.Context, req *models.BookPrimaryKey) (*models.Book, error) {

	var (
		query     string
		resp 	  models.Book
	)

	query = `
		SELECT
			id, 
			name, 
			price,
			COALESCE(count,0),
			came_price,
			profit_status,
			COALESCE(profit,0),
			sell_price,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'), 
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM book
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.Price,
		&resp.Count,
		&resp.Came_price,
		&resp.Profit_status,
		&resp.Profit,
		&resp.Sell_price,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *bookRepo) GetList(ctx context.Context, req *models.GetListBookRequest) (resp *models.GetListBookResponse, err error) {

	resp = &models.GetListBookResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id, 
			name, 
			price,
			COALESCE(count,0),
			came_price,
			profit_status,
			COALESCE(profit,0),
			sell_price,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'), 
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM book
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var book models.Book
		err = rows.Scan(
			&resp.Count,
			&book.Id,
			&book.Name,
			&book.Price,
			&book.Count,
			&book.Came_price,
			&book.Profit_status,
			&book.Profit,
			&book.Sell_price,
			&book.CreatedAt,
			&book.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Books = append(resp.Books, &book)
	}

	return resp, nil
}

func (r *bookRepo) Update(ctx context.Context, req *models.UpdateBook) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			book
		SET
			name = :name,
			price = :price,
			count = :count,
			came_price = :came_price,
			profit_status = :profit_status,
			profit	= :profit,
			sell_price = :sell_price,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":    req.Id,
		"name":  req.Name,
		"price": req.Price,
		"count": req.Count,
		"came_price": req.Came_price,
		"profit_status": req.Profit_status,
		"profit" : req.Profit,
		"sell_price" : req.Sell_price,

	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *bookRepo) Delete(ctx context.Context, req *models.BookPrimaryKey) error {

	_, err := r.db.Exec(
		ctx, "DELETE FROM book WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
