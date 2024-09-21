package testservice

import (
	"context"
	"errors"
	"fmt"
	"vvnbd/internal/domain"
	"vvnbd/internal/domain/db"
)

const (
	login = "johndoe"
)

func (s *Service) GetHelloWorld(ctx context.Context) (result domain.HelloWorld, isFound bool) {
	helloWorld, err := s.repo.GetHelloWorld(ctx, login)
	if err != nil {
		var notFoundErr *db.NotFoundError
		if errors.As(err, &notFoundErr) {
			return domain.HelloWorld{}, false
		}

		fmt.Println("unable to get hello world Err: %w", err)
		return domain.HelloWorld{}, false
	}

	return helloWorld, true
}
