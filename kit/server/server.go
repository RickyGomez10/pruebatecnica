package server

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	mux http.ServeMux
}

func NewServer() Server {
	svc := Server{
		mux: *http.NewServeMux(),
	}
	return svc

}

func (s *Server) Run(port string) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()
	errs := make(chan error, 2)
	go func() {
		sugar.Info("Listening to port:", port)
		errs <- http.ListenAndServe(":"+port, &s.mux)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		sugar.Sync()
	}()

	sugar.Info("terminated ", <-errs)

}
func (s *Server) RegisterRoutes(path string, handler http.Handler) {
	s.mux.Handle(path, handler)
}
