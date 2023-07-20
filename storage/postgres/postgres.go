package postgres

import (
	"database/sql"
	"fmt"
	"user/config"
	"user/storage"

	_ "github.com/lib/pq"
)

type store struct {
	db   *sql.DB
	user *UserRepo
}

func NewConnectionPostgres(cfg config.Config) (storage.StorageI, error) {
	connectionStr := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresDatabase, cfg.PostgresPassword, cfg.PostgresPort,
	)
	sqlDB, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return &store{
		db:   sqlDB,
		user: NewUserRepo(sqlDB),
	}, nil
}

func (s *store) Close() {
	err := s.db.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *store) User() storage.UserRepoI {
	s.user = NewUserRepo(s.db)

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}
	return s.user
}
