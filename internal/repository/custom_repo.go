package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/models"
)

type CustomData interface {
	GetUsersCount(context.Context) (int, error)
	AddUser(context.Context, *models.User) error
	AddAccount(context.Context, *models.Account) error
}

type customData struct {
	db *sql.DB
}

func NewCustomData(db *sql.DB) (CustomData, error) {
	if db == nil {
		return nil, errors.New("db is empty")
	}

	return &customData{
		db: db,
	}, nil
}

func (r *customData) AddUser(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, insertUserQuery, &user.Email, &user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *customData) AddAccount(ctx context.Context, account *models.Account) error {
	_, err := r.db.ExecContext(ctx,
		insertAccountQuery,
		&account.LoginID,
		&account.Name,
		&account.Surname,
		&account.Birthday,
		&account.Sex,
		&account.Hobby,
		&account.City,
		&account.Avatar,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *customData) GetUsersCount(ctx context.Context) (int, error) {
	var countUsers int
	err := r.db.QueryRowContext(ctx, countUsersQuery).Scan(&countUsers)
	if err != nil {
		return 0, err
	}
	return countUsers, nil
}
