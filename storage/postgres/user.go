package postgres

import (
	"database/sql"
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

func (r *UserRepo) GetByIdUser(req *models.UserPrimaryKey) (*models.User, error) {
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
