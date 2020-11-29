package apiserver

import (
	"fmt"
	filmHandler "github.com/Arkadiyche/http-rest-api/internal/pkg/film/delivery/http"
	filmRep "github.com/Arkadiyche/http-rest-api/internal/pkg/film/repository"
	filmUC "github.com/Arkadiyche/http-rest-api/internal/pkg/film/usecase"
	sessionRep "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/session/repository"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/middleware"
	personHandler "github.com/Arkadiyche/http-rest-api/internal/pkg/person/deliver/http"
	personRep "github.com/Arkadiyche/http-rest-api/internal/pkg/person/repository"
	personUC "github.com/Arkadiyche/http-rest-api/internal/pkg/person/usecase"
	playlistHandler "github.com/Arkadiyche/http-rest-api/internal/pkg/playlist/delivery/http"
	playlistRep "github.com/Arkadiyche/http-rest-api/internal/pkg/playlist/repository"
	playlistUC "github.com/Arkadiyche/http-rest-api/internal/pkg/playlist/usecase"
	ratingHandler "github.com/Arkadiyche/http-rest-api/internal/pkg/rating/delivery/http"
	ratingRep "github.com/Arkadiyche/http-rest-api/internal/pkg/rating/repository"
	ratingUC "github.com/Arkadiyche/http-rest-api/internal/pkg/rating/usecase"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/store"
	subscibeHandler "github.com/Arkadiyche/http-rest-api/internal/pkg/subscribe/deliver/http"
	subscibeRep "github.com/Arkadiyche/http-rest-api/internal/pkg/subscribe/repository"
	subscribeUC "github.com/Arkadiyche/http-rest-api/internal/pkg/subscribe/usecase"
	userHandler "github.com/Arkadiyche/http-rest-api/internal/pkg/user/delivery/http"
	userRep "github.com/Arkadiyche/http-rest-api/internal/pkg/user/repository"
	userUC "github.com/Arkadiyche/http-rest-api/internal/pkg/user/usecase"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	config    *Config
	logger    *logrus.Logger
	router    *mux.Router
	store     *store.Store
	sanitazer *bluemonday.Policy
}

func New(config *Config) *APIServer {
	return &APIServer{
		config:    config,
		logger:    logrus.New(),
		router:    mux.NewRouter(),
		sanitazer: bluemonday.UGCPolicy(),
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
	user, film, rating, person, playlist, subscribe := s.InitHandler()
	//User routes ...
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/signup", user.Signup)
	s.router.HandleFunc("/login", user.Login)
	s.router.HandleFunc("/me", user.Me)
	s.router.HandleFunc("/logout", user.Logout)
	s.router.HandleFunc("/chengelogin", user.ChangeLogin())
	s.router.HandleFunc("/chengepass", user.ChangePassword())
	s.router.HandleFunc("/changeAvatar", user.ChangeAvatar)
	s.router.HandleFunc("/user/avatar/{id:[0-9]+}", user.AvatarById)
	s.router.HandleFunc("/people/{id:[0-9]+}", user.GetById)
	//Film routes ...
	s.router.HandleFunc("/film/{id:[0-9]+}", film.FilmById)
	s.router.HandleFunc("/film/{genre:[A-z]+}", film.FilmsByGenre)
	s.router.HandleFunc("/person_film/{id:[0-9]+}", film.FilmsByPerson)
	//Rate routes ...
	s.router.HandleFunc("/rate", rating.Rate())
	s.router.HandleFunc("/review/add", rating.AddReview())
	s.router.HandleFunc("/review/{film_id:[0-9]+}", rating.ShowReviews)
	s.router.HandleFunc("/currentRating/{film_id:[0-9]+}", rating.GetCurrentUserRating())
	//Person routes ...
	s.router.HandleFunc("/person/{id:[0-9]+}", person.PersonById)
	s.router.HandleFunc("/{role:actor|director}/{film_id:[0-9]+}", person.PersonsByFilm)
	//Playlist routes ...
	s.router.HandleFunc("/playlist/create" , playlist.CreatePlaylist())
	s.router.HandleFunc("/playlist/add" , playlist.AddPlaylist())
	s.router.HandleFunc("/playlist/list" , playlist.ShowList)
	s.router.HandleFunc("/playlist/show" , playlist.ShowPlaylist)
	s.router.HandleFunc("/playlist/delete" , playlist.DeletePlaylist())
	s.router.HandleFunc("/playlist/remove" , playlist.RemovePlaylist())
	//Subscribe routes ...
	s.router.HandleFunc("/follow" , subscribe.Subscribe())
	s.router.HandleFunc("/unfollow" , subscribe.UnSubscribe())
	s.router.HandleFunc("/authors" , subscribe.ShowAuthors)
	s.router.HandleFunc("/news" , subscribe.ShowFeed)
	s.router.HandleFunc("/sub/check", subscribe.Check())

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

func (s *APIServer) InitHandler() (userHandler.UserHandler, filmHandler.FilmHandler, ratingHandler.RatingHandler, personHandler.PersonHandler, playlistHandler.PlaylistHandler, subscibeHandler.SubscribeHandler) {

	SessionRep := sessionRep.NewSessionRepository(s.store.Db)
	//user
	UserRep := userRep.NewUserRepository(s.store.Db)
	UserUC := userUC.NewUserUseCase(UserRep, SessionRep)
	UserHandler := userHandler.UserHandler{
		UseCase:   UserUC,
		Logger:    s.logger,
		Sanitazer: s.sanitazer,
	}
	//film
	FilmRep := filmRep.NewFilmRepository(s.store.Db)
	FilmUC := filmUC.NewFilmUseCase(FilmRep)
	FilmHandler := filmHandler.FilmHandler{
		UseCase: FilmUC,
		Logger:  s.logger,
	}
	//rating
	RatingRep := ratingRep.NewRatingRepository(s.store.Db)
	RatingUC := ratingUC.NewRatingUseCase(RatingRep, SessionRep)
	RatingHandler := ratingHandler.RatingHandler{
		UseCase:   RatingUC,
		Logger:    s.logger,
		Sanitazer: s.sanitazer,
	}
	//person
	PersonRep := personRep.NewPersonRepository(s.store.Db)
	PersonUC := personUC.NewPersonUseCase(PersonRep)
	PersonHandler := personHandler.PersonHandler{
		UseCase: PersonUC,
		Logger:  s.logger,
	}
	//playlist
	PlaylistRep := playlistRep.NewRPlaylistRepository(s.store.Db)
	PlaylistUC := playlistUC.NewPlaylistUseCase(PlaylistRep, FilmRep, SessionRep)
	PlaylistHandler := playlistHandler.PlaylistHandler{
		UseCase:   PlaylistUC,
		Logger:    s.logger,
		Sanitazer: s.sanitazer,
	}
	//subscribe
	SubscribeRep := subscibeRep.NewSubscribeRepository(s.store.Db)
	SubscribeUC := subscribeUC.NewSubscribeUseCase(SubscribeRep, SessionRep)
	SubscribeHandler := subscibeHandler.SubscribeHandler{
		UseCase:   SubscribeUC,
		Logger:    s.logger,
		Sanitazer: s.sanitazer,
	}


	return UserHandler, FilmHandler, RatingHandler, PersonHandler, PlaylistHandler, SubscribeHandler
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}
}
