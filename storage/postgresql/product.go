package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct{
	db	*pgxpool.Pool
}

func NewProductRepoI(db *pgxpool.Pool) *productRepo{
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) CreateProduct(ctx context.Context, req *models.CreateProduct) (string, error) {

	var (
		query	string
		id = 	uuid.New().String()
	)

	query = `
		INSERT INTO products (
			id,
			name,
			price,
			category_id,
			updated_at
		) VALUES ($1, $2, $3, $4, now())
	`

	_, err := p.db.Exec(ctx, query, 
		id,
		req.Name,
		req.Price,
		req.Category_id,
	)
	if err != nil{
		return "", err
	}

	return id, nil
}

func (p *productRepo) GetByIdProduct(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {

	var(
		query	string
		product	models.Product
	)

	query = `
		SELECT
			id,
			name,
			COALESCE(price, 0),
			category_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM
			products
		WHERE id = $1
	`

	err := p.db.QueryRow(ctx, query, req.Id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Category_id,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil{
		return nil, err
	}

	return &product, nil
}

func (p *productRepo) GetListProduct(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error) {

	var (
		query		string
		filter =	" WHERE TRUE"
		offset = 	" OFFSET 0"
		limit = 	" LIMIT 0" 
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			COALESCE(price, 0),
			category_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM products
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

	fmt.Println(query)

	rows, err := p.db.Query(ctx, query)
	if err != nil{
		return nil, err
	}

	resp := models.GetListProductResponse{}

	for rows.Next(){

		var product models.Product

		rows.Scan(
			&resp.Count,
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Category_id,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		resp.Products = append(resp.Products, &product)
	}

	resp.Count = len(resp.Products)

	return &resp, nil

}

func (p *productRepo) UpdateProduct(ctx context.Context, req *models.UpdateProduct) (int64, error) {

	query := `
		UPDATE
			products
		SET
			name = $1,
			price = $2,
			category_id = $3,
			updated_at = now()
		WHERE id = $4
	`

	res, err := p.db.Exec(ctx, query,
		req.Name,
		req.Price,
		req.Category_id,
		req.Id,
	)
	if err != nil{
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (p *productRepo) DeleteProduct(ctx context.Context, req *models.ProductPrimaryKey) (error) {

	_, err := p.db.Exec(ctx,
		"DELETE FROM products WHERE id = $1", req.Id,
	)

	if err != nil{
		return err
	}

	return nil
}