package handlers

import (
	"net/http"

	"sandbox-api/logs"
	"sandbox-api/services"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var logger *logs.AppLogger

type IHandler interface {
	RegisterRoutes()
	ApplyCors() http.Handler
	AppHandler(next http.Handler) http.HandlerFunc
}

type handler struct {
	users  IUsers
	items  IItems
	makers IMakers
	router *mux.Router
}

type accessCheck struct {
	granted bool
}

func NewHandler(as *services.Services, r *mux.Router) IHandler {
	logger = logs.NewAppLogger()
	return &handler{
		users:  NewUsers(as.IUserService),
		items:  newItems(as.IItemService),
		makers: NewMakers(as.IMakerService),
		router: r,
	}
}

func (h *handler) ApplyCors() http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	return c.Handler((h.router))
}

func (h *handler) AppHandler(next http.Handler) http.HandlerFunc {
	return h.HandlerFn(next.ServeHTTP)
}

func (h *handler) HandlerFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("[HANDLER] %v requested %v to %v", r.RemoteAddr, r.Method, r.URL.Path)

		if r.URL.Path == "/signin/" {
			next(w, r)
			return
		}

		h.users.AuthHandler(next)

	})
}

func (h *handler) RegisterRoutes() {
	h.items.Register(h.router)
	h.makers.Register(h.router)
	h.users.Register(h.router)
}
