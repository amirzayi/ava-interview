package service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/amirzayi/ava-interview/database/model"
)

type Service struct {
	db *model.Queries
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: model.New(db),
	}
}

func (s *Service) ListUsers(ctx context.Context) ([]model.User, error) {
	users, err := s.db.ListUsers(ctx)
	if err != nil {
		log.Print(err)
	}
	return users, err
}

func (s *Service) CreateUser(ctx context.Context, arg model.CreateUserParams) (model.User, error) {
	user, err := s.db.CreateUser(ctx, arg)
	if err != nil {
		log.Print(err)
	}
	return user, err
}

func (s *Service) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	user, err := s.db.GetUserByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Print(err)
	}
	return user, err
}

func (s *Service) DeleteUserByID(ctx context.Context, id int64) error {
	err := s.db.DeleteUserByID(ctx, id)
	if err != nil {
		log.Print(err)
	}
	return err
}
