package main

import (
	_healthcheckDelivery "base-golang-echo-setup/modules/healthcheck/delivery/http"
	_healthcheckRepo "base-golang-echo-setup/modules/healthcheck/repository"
	_healthcheckUseCase "base-golang-echo-setup/modules/healthcheck/usecase"
	"base-golang-echo-setup/services/mysqlrepo"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func init() {
	viper.SetConfigFile("config.yml") // name of config file (without extension)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	mysqlrepo.Setup()
}

func main() {
	e := echo.New()

	db, err := gorm.Open(mysql.Open(mysqlrepo.DBConfig.URL))

	if err != nil {
		log.Fatal(err.Error())
	}
	conn, _ := db.DB()
	defer conn.Close()

	healthCheckRepo := _healthcheckRepo.NewRepository(db)
	healthCheckUsecase := _healthcheckUseCase.NewUseCase(e, healthCheckRepo)
	_healthcheckDelivery.NewHandler(e, healthCheckUsecase)

	e.Logger.Fatal(e.Start(":1323"))
}
