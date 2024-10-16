package api

type ApiService interface {
	RegisterRoutes(server *ApiServer)
}
