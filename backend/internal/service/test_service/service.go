package testservice

import (
	"context"
	"vvnbd/internal/domain"
)

type repo interface {
	GetHelloWorld(ctx context.Context, login string) (domain.HelloWorld, error)
	InsertHelloWorld(ctx context.Context, helloWorld domain.HelloWorld) error
}

type Service struct {
	repo repo
}

func NewService(ctx context.Context, repo repo) *Service {
	return &Service{
		repo: repo,
	}
}
