package facker

import (
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ingener-vladimir/go_practices/db_generator/internal"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/logger"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/service"
	"github.com/jaswdr/faker"
)

type CustomFaker struct {
	service service.CustomFacker
	facker  *faker.Faker
	logger  *logger.Logger
	ctx     context.Context
}

func NewFacker(service service.CustomFacker, facker *faker.Faker, logger *logger.Logger) (*CustomFaker, error) {
	if service == nil {
		return nil, errors.New("service is empty")
	}
	if facker == nil {
		return nil, errors.New("facker is empty")
	}
	if logger == nil {
		return nil, errors.New("logger is empty")
	}

	return &CustomFaker{
		service: service,
		facker:  facker,
		logger:  logger,
		ctx:     context.Background(),
	}, nil
}

func (f *CustomFaker) GenerateUsers(count int) error {
	for index := 0; index < count; index++ {
		internet := f.facker.Internet()

		err := f.service.AddUser(f.ctx, internet.Email(), internet.Password())
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *CustomFaker) GenerateAccounts(count int) error {
	for index := 0; index < count; index++ {
		person := f.facker.Person()
		name := person.FirstNameFemale()
		surname := person.LastName()
		sex := "f"
		if count%2 == 0 {
			name = person.FirstNameMale()
			sex = "m"
		}

		hobby := f.facker.Company().JobTitle()
		city := f.facker.Address().Country()

		err := f.service.AddAccount(f.ctx, name, surname, internal.DefaultAvatar, "2006-01-02", sex, hobby, city)
		if err != nil {
			return err
		}
	}

	return nil
}
