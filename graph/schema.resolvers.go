package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"auth-system/graph/model"
	"auth-system/internal/models"
	"context"
	"strconv"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	user, err := r.AuthService.Register(input.Email, input.Password, input.FirstName, input.LastName)
	if err != nil {
		return nil, err
	}

	token, err := r.AuthService.Login(input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Token: token,
		User: &model.User{
			ID:        strconv.FormatUint(uint64(user.ID), 10),
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	token, err := r.AuthService.Login(input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	var user models.User
	// TODO: Get user from database using the token claims

	return &model.AuthResponse{
		Token: token,
		User: &model.User{
			ID:        strconv.FormatUint(uint64(user.ID), 10),
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user, err := r.AuthService.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        strconv.FormatUint(uint64(user.ID), 10),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type MutationResolver interface {
	Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error)
	Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error)
}
type QueryResolver interface {
	Me(ctx context.Context) (*model.User, error)
}
*/