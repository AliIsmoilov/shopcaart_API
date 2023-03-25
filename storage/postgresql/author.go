package postgresql

import (
	"app/api/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)


type authorRepo struct{
	db *sql.DB	
}

func NewAuthorRepo(db *sql.DB) *authorRepo {
	return &authorRepo{
		db: db,
	}
}

func (a *authorRepo) CreateAuthor(req *models.CreateAuthor) (string, error) {

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

	_, err := a.db.Exec(query, 
		id,
		req.Name,
	)

	if err != nil{
		return "", err
	}

	return id, nil
}

func (a *authorRepo) AuthorGetById(req *models.AuthorPrimaryKey) (*models.Author, error) {

	var resp models.Author

	query := `SELECT
				id,
				name,
				created_at
			FROM author
			WHERE id = $1  	
	`

	err := a.db.QueryRow(query, req.Id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.CreatedAt,
	)

	if err != nil{
		return nil, err
	}

	return &resp, nil
}

func (a *authorRepo) GetListAuthor(req *models.GetListAuthorRequest) (*models.GetListAuthorResponse, error) {

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

	rows, err := a.db.Query(query)
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

func (a *authorRepo) UpdateAuthor(req *models.UpdateAuthor) (int64, error) {

	query := `
		UPDATE
			author
		SET
			name = $1
		WHERE id = $2
	`

	rows, err := a.db.Exec(query, 
		req.Name,
		req.Id,
	)

	if err != nil{
		return 0, err
	}

	RowsAffected, err := rows.RowsAffected()
	if err != nil{
		return 0, err
	}

	return RowsAffected, nil
}