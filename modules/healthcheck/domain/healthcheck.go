package healthcheck

type Response struct {
	Status struct {
		DB  ResourceStatus `json:"db"`
		API ResourceStatus `json:"api"`
	} `json:"status"`
}

type ResourceStatus struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type UseCase interface {
	CheckHealth() (res Response, err error)
	CheckDBHealth() (res ResourceStatus)
	CheckAPIHealth() (res ResourceStatus)
}

type Repository interface {
	PingDB() (res ResourceStatus)
}
