// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
)

type Querier interface {
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCategories(ctx context.Context, id int32) error
	GetCategories(ctx context.Context, arg GetCategoriesParams) ([]Category, error)
	GetCategoriesByUserIdAndType(ctx context.Context, arg GetCategoriesByUserIdAndTypeParams) ([]Category, error)
	GetCategoriesByUserIdAndTypeAndDescription(ctx context.Context, arg GetCategoriesByUserIdAndTypeAndDescriptionParams) ([]Category, error)
	GetCategoriesByUserIdAndTypeAndTitle(ctx context.Context, arg GetCategoriesByUserIdAndTypeAndTitleParams) ([]Category, error)
	GetCategory(ctx context.Context, id int32) (Category, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	UpdateCategories(ctx context.Context, arg UpdateCategoriesParams) (Category, error)
}

var _ Querier = (*Queries)(nil)
