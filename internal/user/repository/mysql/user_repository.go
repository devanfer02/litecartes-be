package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/devanfer02/litecartes/domain"
	"github.com/devanfer02/litecartes/internal/utils"
)

type mysqlUserRepository struct {
    Conn    *sql.DB
}

func NewMysqlUserRepository(conn *sql.DB) domain.UserRepository {
    return &mysqlUserRepository{conn}
}

func(m *mysqlUserRepository) fetch(
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
            &user.DisplayName,
            &user.Email     ,
            &user.SubID     ,
            &schoolId       ,
            &user.TotalExp  ,
            &user.Gems      ,
            &user.Streaks   ,
            &user.Level     , 
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

func (m *mysqlUserRepository) fetchPaged(ctx context.Context, cursor domain.Cursor) ([]domain.User, error){
    var query string 
    var users []domain.User
    var err error 
    
    if cursor.CreatedAt == "" {
        query = "SELECT * FROM user ORDER BY created_at LIMIT ?"
        users, err = m.fetch(ctx, query, cursor.LimitData)
    } else {
        query = fmt.Sprintf(
            "SELECT * FROM user WHERE created_at %s ? ORDER BY created_at LIMIT ?", 
            utils.GetPaginationOperator(cursor.PointNext),
        )
        users, err = m.fetch(ctx, query, cursor.CreatedAt, cursor.LimitData)
    }

    return users, err    
}

func(m *mysqlUserRepository) Fetch(
    ctx context.Context, 
    cursor domain.Cursor,
) ([]domain.User, *domain.PaginationResponse, error) {
        
    users, err := m.fetchPaged(ctx, cursor)

    if err != nil {
        return nil, nil, domain.ErrServerError
    }

    if len(users) == 0 {
        return []domain.User{}, nil, nil 
    }

    prevPage := utils.CreateCursor(users[0].CreatedAt, false, cursor.LimitData)
    nextPage := utils.CreateCursor(users[len(users)-1].CreatedAt, true, cursor.LimitData)

    pageResponse := utils.CreatePaginationResponse(prevPage, nextPage)

    return users, &pageResponse, nil 
}

func(m *mysqlUserRepository) FetchUsersLike(
    ctx context.Context,
    param, 
    arg string,
) ([]domain.User, error) {
    query := fmt.Sprintf("SELECT * FROM user WHERE %s LIKE ?", param)

    users, err := m.fetch(ctx, query, arg)

    if err != nil {
        return nil, err 
    }

    return users, nil
}

func(m *mysqlUserRepository) FetchOneByArg(
    ctx context.Context, 
    param,
    arg string,
) (domain.User, error) {
    query := fmt.Sprintf("SELECT * FROM user WHERE %s = ? LIMIT 1", param)

    users, err := m.fetch(ctx, query, arg)

    if err != nil {
        return domain.User{}, domain.ErrServerError 
    }

    if len(users) == 0 {
        return domain.User{}, domain.ErrNotFound
    }

    return users[0], nil 
}

func(m *mysqlUserRepository) InsertUser(
    ctx context.Context,
    user *domain.User,
) (error) {
    query := `INSERT INTO user (uid, username, display_name, email, level, last_active, created_at, updated_at) 
                VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("failed to prepare statement. ERR:%s\n", err.Error())
        return err 
    }

    currTime := utils.CurrentTime()

    _, err = stmt.ExecContext(
        ctx, 
        user.UID, 
        user.Username, 
        user.DisplayName,
        user.Email, 
        user.Level,
        currTime,
        currTime,
        currTime,
    )

    if err != nil {
        log.Printf("failed to execute statement. ERR:%s\n", err.Error())
        return err 
    }

    return nil 
}

func(m *mysqlUserRepository) UpdateUser(
    ctx context.Context, 
    user *domain.UserUpdate,
) (error) {
    query := "UPDATE user SET username = ?, display_name = ?, email = ?, subscription_id = ?, school_id = ?, total_exp = ?, gems = ?, streaks = ?, level = ?, last_active = ?, updated_at = ? WHERE uid = ?"

    if user.SchoolID != nil && *user.SchoolID == 0 {
        return domain.ErrBadRequest
    } else if user.SchoolID == nil {
        query = "UPDATE user SET username = ?, display_name = ?, email = ?, subscription_id = ?, total_exp = ?, gems = ?, streaks = ?, level = ?, last_active = ?, updated_at = ? WHERE uid = ?"
    }

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("failed to prepare statement. ERR:%s\n", err.Error())
        return err 
    }

    currTime := utils.CurrentTime()

    var rows sql.Result

    if user.SchoolID == nil {
        rows, err = stmt.ExecContext(
            ctx, 
            user.Username, 
            user.Email, 
            user.SubID, 
            user.TotalExp, 
            user.Gems, 
            user.Streaks, 
            user.Level,
            currTime, 
            currTime, 
            user.UID,
        )
    } else {
        rows, err = stmt.ExecContext(
            ctx, 
            user.Username, 
            user.DisplayName,
            user.Email, 
            user.SubID, 
            user.SchoolID,
            user.TotalExp, 
            user.Gems, 
            user.Streaks, 
            user.Level,
            currTime, 
            currTime, 
            user.UID,
        )
    }

    if err != nil {
        if err == sql.ErrNoRows {
            return domain.ErrNotFound
        }
        log.Printf("failed to exec statement. ERR:%s\n", err.Error())
        return err 
    }

    affected, _ := rows.RowsAffected()

    if affected != 1 {
        if affected == 0 {
            return domain.ErrNotFound
        } else {
            return fmt.Errorf("weird behaviour. rows affected:%d\n", affected)
        }
    }

    return nil 
}

func(m *mysqlUserRepository) DeleteUser(
    ctx context.Context,
    uid string,
) (error) {
    query := "DELETE FROM user WHERE uid = ?"

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("failed to prepare statement. ERR:%s\n", err.Error())
        return err 
    }

    result, err := stmt.ExecContext(ctx, uid)

    if err != nil {
        log.Printf("failed to execute statement. ERR:%s\n", err.Error())
        return err 
    }

    if affected, _ := result.RowsAffected(); affected != 1 {
        if affected == 0 {
            return domain.ErrNotFound
        } else {
            return fmt.Errorf("weird behaviour. rows affected:%d\n", affected)
        }
    }

    return nil 
}