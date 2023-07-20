package postgres

import (
	"database/sql"
	"fmt"
	"user/models"

	"github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(req *models.CreateUser) (*models.UserPrimaryKey, error) {
	id := uuid.New().String()

	query := `INSERT INTO "users"(id, full_name, login, password) VALUES($1, $2, $3, $4)`

	_, err := r.db.Exec(query, id, req.FullName, req.Login, req.Password)

	if err != nil {
		return nil, err
	}

	return &models.UserPrimaryKey{
		Id: id,
	}, nil
}

func (r *UserRepo) GetById(req *models.UserPrimaryKey) (*models.User, error) {
	user := models.User{}
	query := `SELECT (id, full_name, login, password) FROM "users" WHERE id = $1`
	err := r.db.QueryRow(query, req.Id).Scan(
		&user.Id,
		&user.FullName,
		&user.Login,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetList(req *models.GetListRequest) (*models.GetListResponse, error) {
	var (
		resp   = &models.GetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			full_name,
			login,
			password
		FROM users
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND title ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&resp.Count,
			&user.Id,
			&user.FullName,
			&user.Login,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, user)
	}

	return resp, nil
}

func (r *UserRepo) Update(req *models.User) (*models.User, error) {
	query := `UPDATE users SET full_name = $1, login = $2, password = $3 WHERE id = $4`
	_, err := r.db.Exec(query, req.FullName, req.Login, req.Password, req.Id)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (r *UserRepo) Delete(req *models.UserPrimaryKey) error {
	query := `DELETE users WHERE id = $1`
	_, err := r.db.Exec(query, req.Id)

	return err
}
