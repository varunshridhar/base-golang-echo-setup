package healthcheck

import (
	healthcheck "base-golang-echo-setup/modules/healthcheck/domain"
	"gorm.io/gorm"
)

type repository struct {
	Conn *gorm.DB
}

// NewRepository will create an object that represent the article.Repository interface
func NewRepository(conn *gorm.DB) healthcheck.Repository {
	return &repository{conn}
}

func (r repository) PingDB() (res healthcheck.ResourceStatus) {
	_conn, _ := r.Conn.DB()
	err := _conn.Ping()
	if err != nil {
		res.Message = "Database not connected"
		res.Code = 500
		return
	}
	res.Message = "Database connected"
	res.Code = 200
	return
}
