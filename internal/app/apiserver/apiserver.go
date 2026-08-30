package apiserver

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type APIserver struct {
	config *Config
	logger *slog.Logger
	router *mux.Router
}

func New(config *Config) *APIserver {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &APIserver{
		config: config,
		logger: logger,
		router: mux.NewRouter(),
	}
}

func (s *APIserver) Start() error {
	s.configureRouter()

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handelHello())
}

func (s *APIserver) handelHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Word!"))
	}
}
