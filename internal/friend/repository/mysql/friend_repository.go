package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/devanfer02/litecartes/domain"
)

type mysqlFriendRepository struct {
    Conn *sql.DB
}

func NewMysqlFriendRepository(conn *sql.DB) domain.FriendRepository {
    return &mysqlFriendRepository{conn}
}

func(m *mysqlFriendRepository) fetch(
    ctx context.Context, 
    query string, 
    args ...interface{},
) ([]domain.User, error) {
    rows, err := m.Conn.QueryContext(ctx, query, args...)

    if err != nil {
        log.Printf("error querying to user table. ERR:%s\n", err.Error())
        return nil, domain.ErrServerError
    }

    defer func(){
        err := rows.Close()
        if err != nil {
            log.Printf("failed to close rows. ERR:%s\n", err.Error())
        }
    }()
    
    users := make([]domain.User, 0)

    for rows.Next() {
        var user domain.User
        var schoolId sql.NullInt64
        
        err := rows.Scan(
            &user.UID       ,
            &user.Username  ,
            &user.Email     ,
            &user.SubID     ,
            &schoolId  ,
            &user.TotalExp  ,
            &user.Gems      ,
            &user.Streaks   ,
            &user.LastActive,
            &user.Role      ,
            &user.CreatedAt ,
            &user.UpdatedAt ,
        )

        if err != nil {
            log.Printf("failed to scan row. ERR:%s\n", err.Error())
            return nil, err 
        }

        if schoolId.Valid {
            user.SchoolID = int64(schoolId.Int64)
        } else {
            user.SchoolID = 0
        }

        users = append(users, user)
    }

    return users, nil 
}

func(m *mysqlFriendRepository) FetchUsersFriend(
    ctx context.Context,
    userUID string, 
    column string,
) ([]domain.User, error) {
    query := fmt.Sprintf(`SELECT user.* 
        FROM user 
        JOIN friend ON user.uid = friend.%s WHERE friend.%s = ?`, column, column,
    )

    users, err := m.fetch(ctx, query, userUID)

    if err != nil {
        return nil, err 
    }

    return users, nil 
}

func(m *mysqlFriendRepository) InsertNewFollower(
    ctx context.Context,
    followedID,
    followerID string,
) error {
    query := `INSERT INTO friend (followed_uid, follower_uid) VALUES (?, ?)`

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("[REPOSITORY] failed to prepare statement. ERR:%s\n", err.Error())
        return err
    }

    _, err = stmt.ExecContext(ctx, followedID, followerID)

    if err != nil {
        log.Printf("[REPOSITORY] failed to execute statement. ERR:%s\n", err.Error())
        return err 
    }

    return nil 
}

func(m *mysqlFriendRepository) DeleteFriend(
    ctx context.Context,
    followedID,
    followerID string,
) error {
    query := `DELETE FROM friend WHERE followed_uid = ? AND follower_uid = ?`

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("[REPOSITORY] failed to prepare statement. ERR:%s\n", err.Error())
        return err
    }

    rows, err := stmt.ExecContext(ctx, followedID, followerID)

    if err != nil {
        log.Printf("[REPOSITORY] failed to execute statement. ERR:%s\n", err.Error())
        return err 
    }

    affected, _ := rows.RowsAffected()

    if affected != 1 {
        if affected == 0 {
            return domain.ErrBadRequest
        }

        return fmt.Errorf("weird behaviour. rows affected:%d\n", affected)
    }

    return nil 
}