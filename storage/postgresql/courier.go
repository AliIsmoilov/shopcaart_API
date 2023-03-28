package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type courierRepo struct{
	db	*pgxpool.Pool
}

func NewCourierRepo(db *pgxpool.Pool) *courierRepo{
	return &courierRepo{
		db: db,
	}
}

func (c *courierRepo) CreateCourier(ctx context.Context, req *models.CreateCourier) (string, error) {

	var (
		query	string
		id = 	uuid.New().String()
	)

	query = `
		INSERT INTO courier (
			id,
			name,
			phone_number,
			updated_at
		) VALUES($1, $2, $3, now())
	`

	_, err := c.db.Exec(ctx, query,
		id,
		req.Name,
		req.Phone_number,
	)
	if err != nil{
		return "", err
	}

	return id, nil
}

func (c *courierRepo) GetByIDCourier(ctx context.Context, req *models.CourierPrimaryKey) (*models.Courier, error) {

	var (
		query	string
		courier	models.Courier
	)

	query = `
		SELECT
			id,
			name,
			phone_number,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM courier
		WHERE id = $1
	`

	err := c.db.QueryRow(ctx, query, req.Id).Scan(
		&courier.Id,
		&courier.Name,
		&courier.Phone_number,
		&courier.CreatedAt,
		&courier.UpdatedAt,
	)

	if err != nil{
		return nil, err
	} 

	return &courier, nil
}

func (c *courierRepo) GetListCourier(ctx context.Context, req *models.GetListCourierRequest) (*models.GetListCourierResponse, error) {

	var(
		query	string
		filter=	" WHERE TRUE"
		offset=	" OFFSET 0"
		limit=	" LIMIT 0"
	)

	query = `
		SELECT
		COUNT(*) OVER(),
		id,
		name,
		phone_number,
		TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
		TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
	FROM courier
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

	resp := models.GetListCourierResponse{}

	for rows.Next() {
		
		var courier models.Courier
		
		rows.Scan(
			&resp.Count,
			&courier.Id,
			&courier.Name,
			&courier.Phone_number,
			&courier.CreatedAt,
			&courier.UpdatedAt,
		)

		resp.Couriers = append(resp.Couriers, &courier)
	}

	return &resp, nil
}

func (c *courierRepo) UpdateCourier(ctx context.Context, req *models.UpdateCourier) (int64, error) {

	query := `
		UPDATE 
			courier
		SET
			name = $1,
			phone_number = $2,
			updated_at = now()
		WHERE id = $3
	`

	rows, err := c.db.Exec(ctx, query, 
		req.Name,
		req.Phone_number,
		req.Id,
	)
	if err != nil{
		return 0, err
	}

	return rows.RowsAffected(), nil
}

func (c *courierRepo) DeleteCourier(ctx context.Context, req *models.CourierPrimaryKey) (error) {

	_, err := c.db.Exec(ctx, 
		"DELETE FROM courier WHERE id = $1", req.Id,
	)
	if err != nil{
		return err
	}

	return nil
}