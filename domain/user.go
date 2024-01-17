package domain

import (
    "time"
	"context"
)

type User struct {
    UID              string         `json:"uid"`
    Username         string         `json:"username"`
    DisplayName      string         `json:"display_name"`
    Email            string         `json:"email"`
    SubID            int64          `json:"subscription_id"`
    SchoolID         int64          `json:"school_id"`     
    TotalExp         int64          `json:"total_exp"`
    Gems             int64          `json:"gems"`
    Streaks          int64          `json:"streaks"`
    Level            int64          `json:"level"`
    LastActive       string         `json:"last_active"`
    Role             string         `json:"role"`
    CreatedAt        time.Time      `json:"created_at"`
    UpdatedAt        time.Time      `json:"updated_at"`
}

type UserUpdate struct {
    UID              string         `json:"uid"`
    Username         string         `json:"username" validate:"required"`
    DisplayName      string         `json:"display_name" validate:"required"`
    Email            string         `json:"email" validate:"required"`
    SubID            *int64         `json:"subscription_id" validate:"required"`
    SchoolID         *int64         `json:"school_id"`     
    TotalExp         *int64         `json:"total_exp" validate:"required"`
    Gems             *int64         `json:"gems" validate:"required"`
    Streaks          *int64         `json:"streaks" validate:"required"`
    Level            int64          `json:"level" validate:"required"`
}

type UserRepository interface {
    Fetch(ctx context.Context, cursor Cursor) ([]User, *PaginationResponse, error)
    FetchOneByArg(ctx context.Context, param, arg string) (User, error)
    FetchUsersLike(ctx context.Context, param, arg string) ([]User, error) 
    InsertUser(ctx context.Context, user *User) error
    UpdateUser(ctx context.Context, user *UserUpdate) error
    DeleteUser(ctx context.Context, uid string) error
}

type UserUsecase interface {
	Fetch(ctx context.Context, req *PaginationRequest) ([]User, *PaginationResponse, error)
	FetchByUsername(ctx context.Context, username string) (User, error)
    FetchByUID(ctx context.Context, uid string) (User, error)
    FetchUsersByUsername(ctx context.Context,username string) ([]User, error)
    RegisterUser(ctx context.Context, uid string) error
	UpdateUser(ctx context.Context, user *UserUpdate) error 
	DeleteUser(ctx context.Context, uid string) error 
}