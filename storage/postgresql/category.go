package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct{
	db	*pgxpool.Pool
}

func NewCategoryRepoI(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}


func (c *categoryRepo) CreateCategory(ctx context.Context, req *models.CreateCategory) (string, error) {

	var(
		query	string
		id = 	uuid.New().String()
	)

	query = `
		INSERT INTO categories(
			id,
			name
		) VALUES ($1, $2)
	`

	_, err := c.db.Exec(ctx, query, 
		id,
		req.Name,
	)

	if err != nil{
		return "", err
	}

	return id, nil
}

func (c *categoryRepo) GetByIdCategory(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error) {

	var (
		query		string
		category	models.Category  
	)	

	query = `
		SELECT
			id,
			name
		FROM
			categories
		WHERE id = $1
	`

	err := c.db.QueryRow(ctx, query, req.Id).Scan(
		&category.Id,
		&category.Name,
	)
	
	if err != nil{
		return nil, err
	}

	return &category, nil
}


func (c *categoryRepo) GetListCategory(ctx context.Context, req *models.GetListCatogoryRequest) (*models.GetListCategoryResponse, error) {

	var (
		query		string
		filter = 	" WHERE TRUE"
		offset = 	" OFFSET 0"
		limit = 	" LIMIT 0"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name
		FROM 
			categories
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

	resp := models.GetListCategoryResponse{}

	for rows.Next(){

		var category models.Category

		rows.Scan(
			&resp.Count,
			&category.Id,
			&category.Name,
		)

		resp.Categories = append(resp.Categories, &category)

	}

	// resp.Count = len(resp.Categories)

	return &resp, nil
}

func (c *categoryRepo) UpdateCategory(ctx context.Context, req *models.UpdateCategory) (int64, error) {

	query := `
		UPDATE
			categories
		SET
			name = $1
		WHERE id = $2
	`

	res, err := c.db.Exec(ctx, query, 
		req.Name,
		req.Id,
	)
	if err != nil{
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (c *categoryRepo) DeleteCategory(ctx context.Context, req *models.CategoryPrimaryKey) (error) {

	_, err := c.db.Exec(ctx, 
		"DELETE FROM categories WHERE id = $1", req.Id,
	)

	if err != nil{
		return err
	}

	return nil
}