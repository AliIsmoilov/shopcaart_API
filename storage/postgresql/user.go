package postgresql

import (
	"fmt"

	"github.com/google/uuid"

	"app/api/models"
	"database/sql"
)

type userRepo struct{
	db *sql.DB
}


func NewUserRepo(db *sql.DB) *userRepo{
	return &userRepo{
		db: db,
	}
} 


func (u *userRepo) CreateUser(req *models.CreateUser) (string, error){

	var (
		query string
		id 	= uuid.New().String()
	)

	query = `INSERT INTO users(
			id,
			name,
			balance,
			updated_at
			)
			VALUES($1, $2, $3, now())
		`
	_, err := u.db.Exec(query, 
		id,
		req.Name,
		req.Balance,
	)

	if err != nil{
		return "", err
	}

	return id, err

}

func (u *userRepo) UserGetByID(req *models.UserPrimaryKey) (*models.User, error){

	var (
		query string
		user  models.User
	)

	query = `
			SELECT
				id,
				name,
				balance,
				created_at,
				updated_at
			FROM users
			WHERE id = $1
		`
	
	err := u.db.QueryRow(query, req.Id).Scan(
		&user.Id,
		&user.Name,
		&user.Balance,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil{
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) UserGetList(req *models.GetListUserRequest) (*models.GetListUserResponse, error){

	resp := &models.GetListUserResponse{}

	var (
		query string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit = " LIMIT 0"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			balance,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			updated_at
		FROM users
	`

	if len(req.Search) > 0{
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0{
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := u.db.Query(query)
	if err != nil{
		return nil, err
	}

	for rows.Next(){

		var user models.User

		rows.Scan(
			&resp.Count,
			&user.Id,
			&user.Name,
			&user.Balance,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		resp.Users = append(resp.Users, &user)

	}


	return resp, nil
}


func (u *userRepo) UpdateUser(req *models.UpdateUser) (int64, error){

	query := `
		UPDATE
			users
		SET
			name = $1,
			balance = $2,
			updated_at = now()
		WHERE id = $3
	`

	rows, err := u.db.Exec(query, 
		req.Name,
		req.Balance,
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


func (u *userRepo) DeleteUser(req *models.UserPrimaryKey) error {

	_, err := u.db.Exec(
		"DELETE FROM users WHERE id = $1", req.Id,
	)

	if err != nil{
		return err
	}

	return nil
}