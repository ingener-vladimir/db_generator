package service

import (
	"context"
	"errors"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/logger"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/models"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/repository"
	"math/rand"
)

type CustomFacker interface {
	AddUser(context.Context, string, string) error
	AddAccount(context.Context, string, string, string, string, string, string, string) error
}

type customFacker struct {
	repository repository.CustomData
	logger     *logger.Logger
	loginIDs   map[int]string
}

func NewCustomFacker(repository repository.CustomData, logger *logger.Logger) (CustomFacker, error) {
	if repository == nil {
		return nil, errors.New("repository is empty")
	}
	if logger == nil {
		return nil, errors.New("logger is empty")
	}

	return &customFacker{
		repository: repository,
		logger:     logger,
		loginIDs:   make(map[int]string),
	}, nil
}

func (s *customFacker) AddUser(ctx context.Context, email, password string) error {
	user := &models.User{
		Email:    email,
		Password: password,
	}

	err := user.EncryptPassword()
	if err != nil {
		s.logger.Print(err)
		return err
	}

	err = s.repository.AddUser(ctx, user)
	if err != nil {
		s.logger.Print(err)
		return err
	}
	return nil
}

func (s *customFacker) AddAccount(ctx context.Context, name, surname, avatar, birthday, sex, hobby, city string) error {
	countUsers, err := s.repository.GetUsersCount(ctx)
	if err != nil {
		s.logger.Print(err)
		return err
	}

	var loginID int
	for {
		loginID = rand.Intn(countUsers) + 1
		if _, ok := s.loginIDs[loginID]; !ok {
			s.loginIDs[loginID] = ""
			break
		}

	}

	account := &models.Account{
		LoginID:  loginID,
		Avatar:   avatar,
		Name:     name,
		Surname:  surname,
		Birthday: birthday,
		Sex:      sex,
		Hobby:    hobby,
		City:     city,
	}

	err = s.repository.AddAccount(ctx, account)
	if err != nil {
		s.logger.Print(err)
		return err
	}
	return nil
}
