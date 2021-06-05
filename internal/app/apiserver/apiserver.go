package apiserver

import (
	"fmt"
	filmHandler "github.com/Arkadiyche/http-rest-api/internal/pkg/film/delivery/http"
	filmRep "github.com/Arkadiyche/http-rest-api/internal/pkg/film/repository"
	filmUC "github.com/Arkadiyche/http-rest-api/internal/pkg/film/usecase"
	client "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/client"
	client3 "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/rate/client"
	client2 "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/client"
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
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"log"
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
	//prometheus.MustRegister(middleware.HttpDuration, middleware.HttpHits)
	//User routes ...
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/api/signup", user.Signup)
	s.router.HandleFunc("/api/login", user.Login)
	s.router.HandleFunc("/api/me", user.Me)
	s.router.HandleFunc("/api/logout", user.Logout)
	s.router.HandleFunc("/api/chengelogin", user.ChangeLogin())
	s.router.HandleFunc("/api/chengepass", user.ChangePassword())
	s.router.HandleFunc("/api/changeAvatar", user.ChangeAvatar)
	s.router.HandleFunc("/api/user/avatar/{id:[0-9]+}", user.AvatarById)
	s.router.HandleFunc("/api/user/delete", user.DeleteUser)
	s.router.HandleFunc("/api/people/{id:[0-9]+}", user.GetById)
	s.router.HandleFunc("/api/users/search", user.Search).Queries("key", "{.*}")
	//Film routes ...
	s.router.HandleFunc("/api/film/{id:[0-9]+}", film.FilmById)
	s.router.HandleFunc("/api/film/{genre:[A-z]+}", film.FilmsByGenre)
	s.router.HandleFunc("/api/person_film/{id:[0-9]+}", film.FilmsByPerson)
	s.router.HandleFunc("/api/similar/{id:[0-9]+}", film.SimilarFilms)
	s.router.HandleFunc("/api/films/search", film.Search).Queries("key", "{.*}")
	//Rate routes ...
	s.router.HandleFunc("/api/rate", rating.Rate())
	s.router.HandleFunc("/api/review/add", rating.AddReview())
	s.router.HandleFunc("/api/review/{film_id:[0-9]+}", rating.ShowReviews)
	s.router.HandleFunc("/api/currentRating/{film_id:[0-9]+}", rating.GetCurrentUserRating())
	//Person routes ...
	s.router.HandleFunc("/api/person/{id:[0-9]+}", person.PersonById)
	s.router.HandleFunc("/api/{role:actor|director}/{film_id:[0-9]+}", person.PersonsByFilm)
	s.router.HandleFunc("/api/persons/search", person.Search).Queries("key", "{.*}")
	//Playlist routes ...
	s.router.HandleFunc("/api/playlist/create", playlist.CreatePlaylist())
	s.router.HandleFunc("/api/playlist/add", playlist.AddPlaylist())
	s.router.HandleFunc("/api/playlist/list", playlist.ShowList)
	s.router.HandleFunc("/api/playlist/show", playlist.ShowPlaylist)
	s.router.HandleFunc("/api/playlist/delete", playlist.DeletePlaylist())
	s.router.HandleFunc("/api/playlist/remove", playlist.RemovePlaylist())
	//Subscribe routes ...
	s.router.HandleFunc("/api/follow", subscribe.Subscribe())
	s.router.HandleFunc("/api/unfollow", subscribe.UnSubscribe())
	s.router.HandleFunc("/api/authors", subscribe.ShowAuthors)
	s.router.HandleFunc("/api/news", subscribe.ShowFeed)
	s.router.HandleFunc("/api/sub/check/{user_id:[0-9]+}", subscribe.Check())

	s.router.Handle("/api/metrics", promhttp.Handler())

	s.router.Use(middleware.PrometheusMiddleware)

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

	rpcAuth, err := client.NewAuthClient("localhost", ":8001")
	if err != nil {
		log.Fatal(err.Error())
	}
	rpcSession, err := client2.NewSessionClient("localhost", ":8002")
	if err != nil {
		log.Fatal(err.Error())
	}
	rpcRate, err := client3.NewRateClient("localhost", ":8003")
	if err != nil {
		log.Fatal(err.Error())
	}

	//SessionRep := sessionRep.NewSessionRepository(s.store.Db)
	//user
	UserRep := userRep.NewUserRepository(s.store.Db)
	UserUC := userUC.NewUserUseCase(UserRep, rpcSession)
	UserHandler := userHandler.UserHandler{
		RpcAuth:   rpcAuth,
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
	RatingUC := ratingUC.NewRatingUseCase(RatingRep, rpcSession)
	RatingHandler := ratingHandler.RatingHandler{
		RpcRate:   rpcRate,
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
	PlaylistUC := playlistUC.NewPlaylistUseCase(PlaylistRep, FilmRep, rpcSession)
	PlaylistHandler := playlistHandler.PlaylistHandler{
		UseCase:   PlaylistUC,
		Logger:    s.logger,
		Sanitazer: s.sanitazer,
	}
	//subscribe
	SubscribeRep := subscibeRep.NewSubscribeRepository(s.store.Db)
	SubscribeUC := subscribeUC.NewSubscribeUseCase(SubscribeRep, rpcSession)
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
