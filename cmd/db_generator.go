package main

import (
	"fmt"
	app "github.com/ingener-vladimir/go_practices/db_generator/internal/app"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/config"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/logger"
	"log"
	"os"
	"strconv"
)

func main() {
	if !(len(os.Args) == 2) {
		fmt.Println("usage : ./generator count_rows \nPROJECT DATABASE GENERATION ABORT")
		os.Exit(0)
	}

	countRows, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid value of countRows")
		os.Exit(0)
	}

	// читаем конфигурационные настройки
	cfg, err := config.LoadConfig("configs")
	if err != nil {
		log.Println(err.Error())

		return
	}

	os.Getenv("rows")

	cfg.Log = logger.NewConsole(cfg.LogLevel == config.DebugLevel)

	application := app.NewApp(cfg)
	application.Initialize()
	application.GenerateData(countRows)
}
