package app

import (
	"github.com/ingener-vladimir/go_practices/db_generator/internal/config"
	db2 "github.com/ingener-vladimir/go_practices/db_generator/internal/db"
	customFacker "github.com/ingener-vladimir/go_practices/db_generator/internal/facker"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/logger"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/repository"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/service"
	facker "github.com/jaswdr/faker"
)

type App struct {
	config     *config.Config
	customData repository.CustomData
	service    service.CustomFacker
	fkr        *customFacker.CustomFaker
}

func NewApp(config *config.Config) *App {
	if config == nil {
		panic("empty config!")
	}
	return &App{config: config}
}

func (a *App) Initialize() {
	a.initRepository()
	a.initService()
	a.initFacker()
}

func (a *App) initRepository() {
	db, err := db2.NewConnection(a.config)
	if err != nil {
		panic(err)
	}

	a.customData, err = repository.NewCustomData(db)
	if err != nil {
		panic(err)
	}
}

func (a *App) initService() {
	service, err := service.NewCustomFacker(a.customData, a.config.Log)
	if err != nil {
		panic(err)
	}
	a.service = service
}

func (a *App) initFacker() {
	fkr := facker.New()
	customFacker, err := customFacker.NewFacker(a.service, &fkr, a.config.Log)
	if err != nil {
		panic(err)
	}
	a.fkr = customFacker
}

func (a *App) GenerateData(count int) {
	if err := a.fkr.GenerateUsers(count); err != nil {
		logger.LogError(err, a.config.Log)
		return
	}

	if err := a.fkr.GenerateAccounts(10); err != nil {
		logger.LogError(err, a.config.Log)
		return
	}
}
