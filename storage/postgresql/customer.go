package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)


type customerRepo struct {
	db 	*pgxpool.Pool
}


func NewCustomerRepo(db *pgxpool.Pool) *customerRepo {
	return &customerRepo{
		db: db,
	}
}

func (c *customerRepo) CreateCustomer(ctx context.Context, req *models.CreateCustomer) (string, error) {

	var (
		query	string
		id = 	uuid.New().String()
	)

	query = `INSERT INTO customers(
				id,
				name,
				phone,
				updated_at
				)
				VALUES($1, $2, $3, now())
			`
	_, err := c.db.Exec(ctx, query, 
		id,
		req.Name,
		req.Phone,
	)

	if err != nil{
		return "", err
	}

	return id, nil
}

func (c *customerRepo) GetByIdCustomer(ctx context.Context, req *models.CustomerPrimaryKey) (*models.Customer, error) {

	var (
		query		string
		customer	models.Customer 
	)

	query = `
		SELECT
			id,
			name,
			phone,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')	
		FROM customers
		WHERE id = $1
	`

	err := c.db.QueryRow(ctx, query, req.Id).Scan(
		&customer.Id,
		&customer.Name,
		&customer.Phone,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil{
		return nil, err
	}

	return &customer, nil

}

func (c *customerRepo) GetListCustomer(ctx context.Context, req *models.GetListCustomerRequest) (*models.GetListCustomerResponse, error) {

	var (
		query		string
		filter = 	" WHERE TRUE"
		offset =	" OFFSET 0"
		limit = 	" LIMIT 0"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			phone,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM customers
	`

	if len(req.Search) > 0{
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}  

	if req.Offset > 0{
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0{
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := c.db.Query(ctx, query)
	if err != nil{
		return nil, err
	}

	resp := &models.GetListCustomerResponse{}

	for rows.Next() {

		var customer models.Customer

		rows.Scan(
			&resp.Count,
			&customer.Id,
			&customer.Name,
			&customer.Phone,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)

		resp.Customers = append(resp.Customers, &customer)
	}

	return resp, nil
}

func (c *customerRepo) UpdateCustomer(ctx context.Context, req *models.UpdateCustomer) (int64, error) {

	query := `
		UPDATE
			customers
		SET
			name = $1,
			phone = $2,
			updated_at = now()
		WHERE id = $3
	`	

	rows, err := c.db.Exec(ctx, query, 
		req.Name,
		req.Phone,
		req.Id,
	)

	if err != nil{
		return 0, err
	}

	return rows.RowsAffected(), nil	
}

func (c *customerRepo) DeleteCustomer(ctx context.Context, req *models.CustomerPrimaryKey) (error) {

	_, err := c.db.Exec(ctx, 
		"DELETE FROM customers WHERE id = $1", req.Id,
	)

	if err != nil{
		return err
	}
	return nil
}
