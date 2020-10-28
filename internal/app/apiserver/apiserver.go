package apiserver

import (
	"fmt"
	filmHandler"github.com/Arkadiyche/http-rest-api/internal/pkg/film/delivery/http"
	filmRep "github.com/Arkadiyche/http-rest-api/internal/pkg/film/repository"
	filmUC "github.com/Arkadiyche/http-rest-api/internal/pkg/film/usecase"
	sessionRep"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/sesession/repository"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/middleware"
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
	user, film := s.InitHandler()
	//User routes ...
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/signup", user.Signup)
	s.router.HandleFunc("/login", user.Login)
	s.router.HandleFunc("/me", user.Me)
	s.router.HandleFunc("/logout", user.Logout)
	s.router.HandleFunc("/chengelogin", user.ChangeLogin())
	s.router.HandleFunc("/chengepass", user.ChangePassword())
	//Film routes ...
	s.router.HandleFunc("/film/{id:[0-9]+}", film.FilmById)

	s.router.Use(middleware.CORSMiddleware(s.config.CORS))
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

func (s *APIServer) InitHandler() (userHandler.UserHandler, filmHandler.FilmHandler) {
	//user
	UserRep := 	userRep.NewUserRepository(s.store.Db)
	SessionRep := sessionRep.NewSessionRepository(s.store.Db)
	UserUC := userUC.NewUserUseCase(UserRep, SessionRep)
	UserHandler := userHandler.UserHandler{
		UseCase: UserUC,
	}
	//film
	FilmRep := filmRep.NewFilmRepository(s.store.Db)
	FilmUC := filmUC.NewFilmUseCase(FilmRep)
	FilmHandler := filmHandler.FilmHandler{
		UseCase: FilmUC,
	}
	return UserHandler, FilmHandler
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		 w.Write([]byte("Hello"))
	}
}
