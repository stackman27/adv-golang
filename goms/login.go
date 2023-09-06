package main

import (
	"context"
	"fmt"
	"time"
)

type LoginService struct {
	next Service // middleware to seperate business logic from login
}

func NewLoggingService(next Service) Service{
	return &LoginService{
		next: next,
	}
}

func (s *LoginService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%v  err=%s took=%v", fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetCatFact(ctx)

}