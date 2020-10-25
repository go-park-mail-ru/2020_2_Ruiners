package apiserver

import (
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/store"
	userHandler "github.com/Arkadiyche/http-rest-api/internal/pkg/user/delivery/http"
	userRep "github.com/Arkadiyche/http-rest-api/internal/pkg/user/repository"
	userUC "github.com/Arkadiyche/http-rest-api/internal/pkg/user/usecase"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

func New(config *Config) *APIServer{
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")
	//fmt.Println(s.store.Config())
	
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	user := s.InitHandler()
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/add", user.Signup)
}

func (s *APIServer) configureStore() error {
	fmt.Println(s.config.Store)
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) InitHandler() userHandler.UserHandler {
	UserRep := 	userRep.NewUserRepository(s.store.Db)
	UserUC := userUC.NewUserUseCase(UserRep)
	UserHandler := userHandler.UserHandler{
		UseCase: UserUC,
	}
	return UserHandler
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		 w.Write([]byte("Hello"))
	}
}
