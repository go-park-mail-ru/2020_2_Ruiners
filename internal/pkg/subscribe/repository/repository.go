package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
)

type SubscribeRepository struct {
	db *sql.DB
}

func NewSubscribeRepository(db *sql.DB) *SubscribeRepository {
	return &SubscribeRepository{
		db: db,
	}
}


func (s *SubscribeRepository) AddSubscribe(subscriberId int, authorId int) error {
	_, err := s.db.Exec("insert into subscribe(subscriber, author) VALUE(?, ?)", subscriberId, authorId)
	if err != nil {
		return err
	}
	return nil
}

func (s *SubscribeRepository) DeleteSubscribe(subscriberId int, authorId int) error {
	_, err := s.db.Exec("DELETE FROM subscribe WHERE subscriber = ? AND author = ?", subscriberId, authorId)
	if err != nil {
		return err
	}
	return nil
}

func (s *SubscribeRepository) GetAuthors(subscriberId int) (*models.PublicUsers, error) {
	user := models.PublicUser{}
	users := models.PublicUsers{}
	authorsQuery, err := s.db.Query("SELECT u.id, u.username, u.email FROM subscribe s JOIN users u ON(s.author = u.id) WHERE s.subscriber = ?", subscriberId)
	if err != nil {
		return nil, err
	}
	defer authorsQuery.Close()

	for authorsQuery.Next() {
		if authorsQuery.Scan(&user.Id, &user.Login, &user.Email) != nil {
			return nil, errors.New("db error")
		}
		users = append(users, user)
	}
	return &users, nil
}

func (s *SubscribeRepository) GetRatingFeed(subscriberId int) (*models.Feed, error) {
	sub := models.Subscribe{}
	sub.IsRating = true
	feed := models.Feed{}
	feedQuery, err := s.db.Query("SELECT r.rating, r.user_id, u.username, r.film_id, f.title, UNIX_TIMESTAMP(r.create_date)  FROM subscribe s JOIN users u ON (s.author = u.id) JOIN rating r ON (u.id = r.user_id) JOIN films f ON (r.film_id = f.id) where s.subscriber = ?", subscriberId)
	if err != nil {
		return nil, err
	}
	defer feedQuery.Close()

	for feedQuery.Next() {
		if feedQuery.Scan(&sub.Body, &sub.UserId, &sub.UserLogin, &sub.FilmId, &sub.FilmTitle, &sub.Date) != nil {
			return nil, errors.New("db error")
		}
		fmt.Println(sub)
		feed = append(feed, sub)
	}
	return &feed, nil
}

func (s *SubscribeRepository) GetReviewFeed(subscriberId int) (*models.Feed, error) {
	sub := models.Subscribe{}
	sub.IsReview = true
	feed := models.Feed{}
	feedQuery, err := s.db.Query("SELECT r.body, r.user_id, u.username, r.film_id, f.title, UNIX_TIMESTAMP(r.create_date)  FROM subscribe s JOIN users u ON (s.author = u.id) JOIN review r ON (u.id = r.user_id) JOIN films f ON (r.film_id = f.id) where s.subscriber = ?", subscriberId)
	if err != nil {
		return nil, err
	}
	defer feedQuery.Close()

	for feedQuery.Next() {
		if feedQuery.Scan(&sub.Body, &sub.UserId, &sub.UserLogin, &sub.FilmId, &sub.FilmTitle, &sub.Date) != nil {
			return nil, errors.New("db error")
		}
		fmt.Println(sub)
		feed = append(feed, sub)
	}
	return &feed, nil
}

