package post

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingPost = errors.New("failed to fetch post by id")
	ErrNotImplemented = errors.New("not implemented")
)

// Post - a representation of the post structure for our service
type Post struct {
	ID string
	Slug string
	Body string
	author string
}

// Store - this interface defines all the methods that our service needs to operate
type Store interface {
	GetPost(context.Context, string) (Post, error)
}

// Service - is the struct on which all our logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetPost(ctx context.Context, id string)(Post, error){
	fmt.Println("Retrieving a post")
	pst, err := s.Store.GetPost(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Post{}, ErrFetchingPost
	}

	return pst, nil
}

func (s *Service) UpdatePost(ctx context.Context, pst Post) error {
	return ErrNotImplemented
}

func (s *Service) DeletePost(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreatePost(ctx context.Context, pst Post) (Post, error) {
	return Post{}, ErrNotImplemented
}