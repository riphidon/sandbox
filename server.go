package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sandbox-api/logs"
	"time"

	"golang.org/x/sync/errgroup"
)

const serverStarted string = "Started on port "
const serverProperExit string = "Exited Properly"

type IServer interface {
	start() error
}

type server struct {
	mainCtx    context.Context
	httpserver *http.Server
	log        *logs.AppLogger
}

func newServer(h http.Handler, addr int, c context.Context) IServer {
	return &server{
		mainCtx: c,
		httpserver: &http.Server{
			Addr:         fmt.Sprintf(":%v", addr),
			ReadTimeout:  time.Second * 15,
			WriteTimeout: time.Second * 15,
			IdleTimeout:  time.Second * 60,
			//TLSConfig:    tlsConfig,
			Handler:     h,
			BaseContext: func(l net.Listener) context.Context { return c },
		},
		log: logs.NewAppLogger(),
	}
}

func (s *server) start() error {
	g, gctx := errgroup.WithContext(s.mainCtx)

	g.Go(func() error {
		s.log.Startf(" [SERVER] %v %v", serverStarted, s.httpserver.Addr)
		return s.httpserver.ListenAndServe()
	})

	g.Go(func() error {
		<-gctx.Done()
		if err := s.httpserver.Shutdown(s.mainCtx); err != nil {
			s.log.Fatalf(" [SERVER]  :%v", err)
			return err
		}
		s.log.Infof(" [SERVER] %v", serverProperExit)
		os.Exit(0)
		return nil
	})

	if err := g.Wait(); err != nil {
		s.log.Infof(" [SERVER] Exit Reason: %s \n", err)
		return err
	}
	return nil
}
