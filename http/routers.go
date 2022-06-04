package http

import "forum/service"

type API struct {
	service *service.Service
}

func NewAPI(service service.Service) *API {
	return &API{service: &service}
}

func (a *API) InitRoutes() {
}
