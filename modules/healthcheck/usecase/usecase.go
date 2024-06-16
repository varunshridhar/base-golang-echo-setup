package healthcheck

import (
	"base-golang-echo-setup/modules/healthcheck/domain"
	"errors"
	"github.com/labstack/echo/v4"
)

type useCase struct {
	DB healthcheck.Repository
}

// NewUseCase will initialize the health check use case
func NewUseCase(e *echo.Echo, d healthcheck.Repository) healthcheck.UseCase {
	return &useCase{d}
}

func (u useCase) CheckHealth() (res healthcheck.Response, err error) {
	res.Status.DB = u.CheckDBHealth()
	res.Status.API = u.CheckAPIHealth()

	if res.Status.DB.Code != 200 || res.Status.API.Code != 200 {
		err = errors.New("health check failed")
	}
	return
}

func (u useCase) CheckDBHealth() (res healthcheck.ResourceStatus) {
	res = u.DB.PingDB()
	return
}

// CheckAPIHealth checks if a basic API is working as expected
func (u useCase) CheckAPIHealth() (res healthcheck.ResourceStatus) {
	res.Message = "Success"
	res.Code = 200
	return
}
