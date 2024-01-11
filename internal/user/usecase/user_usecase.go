package usecase

import (
	"context"
	"math/rand"
	"time"

	"github.com/devanfer02/litecartes/domain"
    "github.com/devanfer02/litecartes/internal/utils"
    
	"firebase.google.com/go/v4/auth"
    fberr "firebase.google.com/go/v4/errorutils"
	"github.com/go-playground/validator/v10"
)

type userUsecase struct {
	userRepo    domain.UserRepository 
    ctxTimeout  time.Duration
    userV10    *validator.Validate 
    auth        *auth.Client
}

func NewUserUsecase(
    repo domain.UserRepository, 
    timeout time.Duration, 
    fireAuth *auth.Client,
) domain.UserUsecase {
    return &userUsecase{
        userRepo: repo,
        ctxTimeout: timeout,
        userV10: validator.New(),
        auth: fireAuth, 
    }
}

func(u *userUsecase) Fetch(
    ctx context.Context, 
    req *domain.PaginationRequest,
) ([]domain.User, *domain.PaginationResponse, error) {

    cursor, err := utils.GetCursor(req)
    
    if err != nil {
        return nil, nil, err 
    }

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    users, pageResp, err := u.userRepo.Fetch(c, cursor)

    if err != nil {
        return nil, nil, err 
    }

    return users, pageResp, nil 
}

func(u *userUsecase) FetchByUsername(
    ctx context.Context,
    username string,
) (domain.User, error) {
    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    user, err := u.userRepo.FetchOneByArg(c, "username", username)

    if err != nil {
        return domain.User{}, nil 
    }

    return user, nil 
}


func(u *userUsecase) FetchByUID(
    ctx context.Context,
    uid string,
) (domain.User, error) {
    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    user, err := u.userRepo.FetchOneByArg(c, "uid", uid)

    if err != nil {
        return domain.User{}, nil 
    }

    return user, nil 
}

func(u *userUsecase) RegisterUser(
    ctx context.Context, 
    fireuid string,
) error {
    // fetch credential from fire auth
    fbUser, err := u.auth.GetUser(ctx, fireuid)

    if err != nil {
        if fberr.IsNotFound(err) {
            return domain.ErrNotFound
        }
        return err 
    }

    if fbUser.DisplayName == "" {
        fbUser.DisplayName = u.generateRandomUsername(8)
    }

    user := &domain.User{
        UID: fbUser.UID,
        Username: fbUser.DisplayName,
        Email: fbUser.Email,
    }

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    err = u.userRepo.InsertUser(c, user)

    if err != nil {
        return err 
    }

    return nil 
}

func(u *userUsecase) UpdateUser(
    ctx context.Context,
    user *domain.UserUpdate, 
) error {
    // only authorized user can update account

    if err := u.userV10.Struct(user); err != nil {
        return domain.ValidationFailed(err.Error())
    }

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    err := u.userRepo.UpdateUser(c, user)

    if err != nil {
        return err 
    }

    return nil 
}

func(u *userUsecase) DeleteUser(
    ctx context.Context,
    uid string, 
) error {
    // only authorized user can delete account

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    err := u.userRepo.DeleteUser(c, uid)

    if err != nil {
        return err 
    }

    err = u.auth.DeleteUser(ctx, uid)

    if err != nil {
        return err
    }

    return nil 
}

func(u *userUsecase) generateRandomUsername(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    rand.Seed(time.Now().UnixNano())

    username := make([]byte, length)
    for i := range username {
        username[i] = charset[rand.Intn(len(charset))]
    }

    return string(username)
}