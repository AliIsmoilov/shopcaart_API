package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)


type authorRepo struct{
	db *pgxpool.Pool
}

func NewAuthorRepo(db *pgxpool.Pool) *authorRepo {
	return &authorRepo{
		db: db,
	}
}

func (a *authorRepo) CreateAuthor(ctx context.Context, req *models.CreateAuthor) (string, error) {

	var (
		query 	string
		id 	=	uuid.New().String()	
	)

	query = `INSERT INTO author(
				id,
				name
				)
			VALUES($1, $2)
		`

	_, err := a.db.Exec(ctx, query, 
		id,
		req.Name,
	)

	if err != nil{
		return "", err
	}

	return id, nil
}

func (a *authorRepo) AuthorGetById(ctx context.Context, req *models.AuthorPrimaryKey) (*models.Author, error) {

	var resp models.Author

	query := `SELECT
				id,
				name,
				TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS')
			FROM author
			WHERE id = $1  	
	`

	err := a.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.CreatedAt,
	)

	if err != nil{
		return nil, err
	}

	return &resp, nil
}

func (a *authorRepo) GetListAuthor(ctx context.Context, req *models.GetListAuthorRequest) (*models.GetListAuthorResponse, error) {

	var (
		query 		string
		filter	= 	" WHERE  TRUE"
		offset 	= 	" OFFSET 0"
		limit	=	" LIMIT 0"	
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM author
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

	var resp models.GetListAuthorResponse

	rows, err := a.db.Query(ctx, query)
	if err != nil{
		return nil, err
	}


	for rows.Next() {

		var author models.Author

		rows.Scan(
			&resp.Count,
			&author.Id,
			&author.Name,
			&author.CreatedAt,
		)

		resp.Authors = append(resp.Authors, &author)

	}

	return &resp, nil

}

func (a *authorRepo) UpdateAuthor(ctx context.Context, req *models.UpdateAuthor) (int64, error) {

	query := `
		UPDATE
			author
		SET
			name = $1
		WHERE id = $2
	`

	rows, err := a.db.Exec(ctx, query, 
		req.Name,
		req.Id,
	)

	if err != nil{
		return 0, err
	}

	return rows.RowsAffected(), nil
}

func (a *authorRepo) DeleteAuthor(ctx context.Context, req *models.AuthorPrimaryKey) error {

	_, err := a.db.Exec(
		ctx, "DELETE FROM author WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}
	
	return nil
}