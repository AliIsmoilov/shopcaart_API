package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"app/api/models"
	"app/pkg/helper"
)

type bookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *bookRepo {
	return &bookRepo{
		db: db,
	}
}

func (r *bookRepo) Create(req *models.CreateBook) (string, error) {

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

	_, err := r.db.Exec(query,
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

func (r *bookRepo) GetByID(req *models.BookPrimaryKey) (*models.Book, error) {

	var (
		query     string
		id        		sql.NullString
		name      		sql.NullString
		price     		sql.NullFloat64
		count			sql.NullInt64
		came_price		sql.NullFloat64
		profit_status	sql.NullString
		profit			sql.NullFloat64
		sell_price		sql.NullFloat64
		createdAt 		sql.NullString
		updatedAt 		sql.NullString
	)

	query = `
		SELECT
			id, 
			name, 
			price,
			count,
			came_price,
			profit_status,
			profit,
			sell_price,
			created_at, 
			updated_at
		FROM book
		WHERE id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&id,
		&name,
		&price,
		&count,
		&came_price,
		&profit_status,
		&profit,
		&sell_price,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Book{
		Id:        	id.String,
		Name:      	name.String,
		Price:     	price.Float64,
		Count: 		int(count.Int64),
		Came_price: came_price.Float64,
		Profit_status: profit_status.String,
		Profit: 	profit.Float64,
		Sell_price: sell_price.Float64,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *bookRepo) GetList(req *models.GetListBookRequest) (resp *models.GetListBookResponse, err error) {

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
			count,
			came_price,
			profit_status,
			profit,
			sell_price,
			created_at, 
			updated_at
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

	rows, err := r.db.Query(query)
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

func (r *bookRepo) Update(req *models.UpdateBook) (int64, error) {

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

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *bookRepo) Delete(req *models.BookPrimaryKey) error {

	_, err := r.db.Exec(
		"DELETE FROM book WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
