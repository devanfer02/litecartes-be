package usecase

import (
    "context"
    "time"

    "github.com/devanfer02/litecartes/domain"
)

type friendUsecase struct {
    friendRepo domain.FriendRepository
    ctxTimeout time.Duration
}

func NewFriendUsecase(friendRepo domain.FriendRepository, timeout time.Duration) domain.FriendUsecase {
    return &friendUsecase{friendRepo: friendRepo, ctxTimeout: timeout}
}

func(u *friendUsecase) FetchFollowers(ctx context.Context, userUID string) ([]domain.User, error) {
    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    users, err := u.friendRepo.FetchUsersFriend(c, userUID, "followed_id")

    return users, err 
}

func(u *friendUsecase) FetchFollowings(ctx context.Context, userUID string) ([]domain.User, error) {
    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    users, err := u.friendRepo.FetchUsersFriend(c, userUID, "follower_id")

    return users, err 
}

func(u *friendUsecase) InsertNewFollower(ctx context.Context, followedID, followerID string) error {
    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    err := u.friendRepo.InsertNewFollower(c, followedID, followerID)

    return err 
}

func(u *friendUsecase) DeleteFriend(ctx context.Context, followedID, followerID string) error {
    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    err := u.friendRepo.DeleteFriend(c, followedID, followerID)

    return err 
}


