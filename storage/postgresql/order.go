package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct{
	db	*pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) CreateOrder(ctx context.Context, req *models.CreateOrder) (string, error) {

	var (
		query	string
		id	= 	uuid.New().String()
	)

	query = `
		INSERT INTO orders(
			id,
			name,
			price,
			phone_number,
			latitude,
			longtitude,
			user_id,
			customer_id,
			courier_id,
			product_id,
			quantity,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, now())
	`

	_, err := o.db.Exec(ctx, query, 
		id,
		req.Name,
		req.Price,
		req.Phone_number,
		req.Latitude,
		req.Longtitude,
		req.User_id,
		req.Customer_id,
		req.Courier_id,
		req.Product_id,
		req.Quantity,
	)

	if err != nil{
		return "", err
	}

	return id, nil
}

func (o *orderRepo) GetByIdOrder(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error) {

	var (
		query	string
		order	models.Order
	)

	query = `
		SELECT
			id,
			name,
			COALESCE(price, 0),
			phone_number,
			latitude,
			longtitude,
			user_id,
			customer_id,
			courier_id,
			product_id,
			COALESCE(quantity, 0),
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM
			orders
		WHERE id = $1
	`

	err := o.db.QueryRow(ctx, query, req.Id).Scan(
		&order.Id,
		&order.Name,
		&order.Price,
		&order.Phone_number,
		&order.Latitude,
		&order.Longtitude,
		&order.User_id,
		&order.Customer_id,
		&order.Courier_id,
		&order.Product_id,
		&order.Quantity,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil{
		return nil, err
	}

	return &order, nil
}


func (o *orderRepo) GetListOrders(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error) {
	
	resp = &models.GetListOrderResponse{}
	
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
			COALESCE(price, 0),
			phone_number,
			latitude,
			longtitude,
			user_id,
			customer_id,
			courier_id,
			product_id,
			COALESCE(quantity, 0),
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM orders
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

	
	rows, err := o.db.Query(ctx, query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()


	for rows.Next() {

		var order models.Order
		
		err = rows.Scan(
			&resp.Count,
			&order.Id,
			&order.Name,
			&order.Price,
			&order.Phone_number,
			&order.Latitude,
			&order.Longtitude,
			&order.User_id,
			&order.Courier_id,
			&order.Courier_id,
			&order.Product_id,
			&order.Quantity,
			&order.CreatedAt,
			&order.UpdatedAt,
		)

		if err != nil{
			return nil, err
		}

		resp.Orders = append(resp.Orders, &order)
	}

	resp.Count = len(resp.Orders)

	return resp, nil
}


func (o *orderRepo) UpdateOrder(ctx context.Context, req *models.UpdateOrder) (int64, error) {

	query := `
		UPDATE
			orders
		SET
			name = $1,
			price = $2,
			phone_number = $3,
			latitude = $4,
			longtitude = $5,
			user_id = $6,
			customer_id = $7,
			courier_id = $8,
			product_id = $9,
			quantity = $10,
			updated_at = now()
		WHERE id = $11
	`

	rows, err := o.db.Exec(ctx, query, 
		req.Name,
		req.Price,
		req.Phone_number,
		req.Latitude,
		req.Longtitude,
		req.User_id,
		req.Customer_id,
		req.Courier_id,
		req.Product_id,
		req.Quantity,
		req.Id,
	)

	if err != nil{
		return 0, err
	}

	return rows.RowsAffected(), nil
}

func (o *orderRepo) PatchOrder(ctx context.Context, req *models.PatchRequest) (int64, error) {

	var (
		query	string
		set		string
	)

	if len(req.Fields) <= 0{
		return 0, errors.New("no fields to update")
	}

	for key := range req.Fields{
		set += fmt.Sprintf(" %s = :%s, ", key, key)
	}

	query = `
			UPDATE
				orders
			SET
		` + set + ` updated_at = now()
			WHERE id = :id
	`

	req.Fields["id"] = req.ID

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := o.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}


func (o *orderRepo) DeleteOrder(ctx context.Context, req *models.OrderPrimaryKey) (error) {

	_, err := o.db.Exec(ctx, 
		"DELETE FROM orders WHERE id = $1", req.Id,
	)

	if err != nil{
		return err
	}

	return nil
}