package apiserver

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kotoninja/http-rest-api/internal/store"
)

type APIserver struct {
	config *Config
	logger *slog.Logger
	router *mux.Router
	store  *store.Store
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

	if err := s.configureStore(); err != nil {
		return err
	}
	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handelHello())
}

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIserver) handelHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Word!"))
	}
}
