package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"sandbox-api/config"
	"sandbox-api/database"
	"sandbox-api/handlers"
	"sandbox-api/services"
)

func main() {

	cfg := config.LoadConfig(false)
	dba := database.NewDBAccess(cfg.DataBase)
	as := services.NewAppService(dba)

	r := NewRouter()

	h := handlers.NewHandler(as, r)
	h.RegisterRoutes()
	hdlr := h.ApplyCors()

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := newServer(h.AppHandler(hdlr), cfg.Port, mainCtx)

	if err := srv.start(); err != nil {
		panic(err)
	}

}
