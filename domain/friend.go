package domain

import (
    "context"
)

type FriendUsecase interface {
    FetchFollowers(ctx context.Context, userUID string) ([]User, error)
    FetchFollowings(ctx context.Context, userUID string) ([]User, error)
    InsertNewFollower(ctx context.Context, followedID, followerID string) error
    DeleteFriend(ctx context.Context, followedID, followerID string) error
}

type FriendRepository interface {
    FetchUsersFriend(ctx context.Context, userUID string,  column string) ([]User, error)
    InsertNewFollower(ctx context.Context, followedID, followerID string) error
    DeleteFriend(ctx context.Context, followedID, followerID string) error
}