package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/devanfer02/litecartes/domain"
	"github.com/devanfer02/litecartes/internal/utils"
)

type mysqlTaskRepository struct {
	Conn *sql.DB
}

func NewMysqlTaskRepository(conn *sql.DB) domain.TaskRepository {
    return &mysqlTaskRepository{conn}
}

func(m *mysqlTaskRepository) fetchRows(
    ctx context.Context,
    query string, 
    args ...interface{},
) (*sql.Rows, func(), error) {
    rows, err := m.Conn.QueryContext(ctx, query, args...)

    if err != nil {
        log.Printf("[REPOSITORY] error querying to task table. ERR:%s\n", err.Error())
        return nil, nil, domain.ErrServerError
    }

    closeRow :=  func(){
        err := rows.Close()
        if err != nil {
            log.Printf("[REPOSITORY] failed to close rows. ERR:%s\n", err.Error())
        }
    }

    return rows, closeRow, nil
}

func(m *mysqlTaskRepository) fetchTasks(
    ctx context.Context,
    query string, 
    args ...interface{},
) ([]domain.Task, error) {
    rows, close, err := m.fetchRows(ctx, query, args...)
    defer close()

    if err != nil {
        return nil, err 
    } 

    tasks := make([]domain.Task, 0)

    for rows.Next() {
        var task domain.Task 

        err := rows.Scan(
            &task.UID, 
            &task.Level, 
            &task.Sign,
            &task.LevelCategoryID,
            &task.CreatedAt,
            &task.UpdatedAt,
        )

        if err != nil {
            log.Printf("[REPOSITORY] failed to scan row in task repository. ERR:%s\n", err.Error())
            return nil, err 
        }

        tasks = append(tasks, task)
    }

    return tasks, nil 
}

func(m *mysqlTaskRepository) fetchTaskWithUserCompletion(
    ctx context.Context, 
    query string, 
    args ...interface{},
) ([]domain.Task, error) {
    rows, close, err := m.fetchRows(ctx, query, args...)
    defer close()

    if err != nil {
        return nil, err 
    } 

    tasks := make([]domain.Task, 0)

    for rows.Next() {
        var task domain.Task 

        err := rows.Scan(
            &task.UID, 
            &task.Level, 
            &task.Sign,
            &task.LevelCategoryID,
            &task.CreatedAt,
            &task.UpdatedAt,
            &task.Completed,
        )

        if err != nil {
            log.Printf("[REPOSITORY] failed to scan row in task repository. ERR:%s\n", err.Error())
            return nil, err 
        }

        tasks = append(tasks, task)
    }

    return tasks, nil 
}

func(m *mysqlTaskRepository) FetchAll(
    ctx context.Context,
    cursor domain.Cursor,
) ([]domain.Task, *domain.PaginationResponse, error) {
    var tasks []domain.Task 
    var err error 


    if cursor.CreatedAt == "" {
        query := `SELECT * FROM task ORDER BY created_at LIMIT ?`
        tasks, err = m.fetchTasks(ctx, query, cursor.LimitData)
    } else {
        query := fmt.Sprintf(
            `SELECT * FROM task WHERE created_at %s ? ORDER BY created_at LIMIT ?`,
            utils.GetPaginationOperator(cursor.PointNext),
        )
        tasks, err = m.fetchTasks(ctx, query, cursor.CreatedAt, cursor.LimitData)
    }


    if err != nil {
        return nil, nil, err 
    }

    if len(tasks) == 0 {
        return []domain.Task{}, nil, nil 
    }

    prevPage := utils.CreateCursor(tasks[0].CreatedAt, false, cursor.LimitData)
    nextPage := utils.CreateCursor(tasks[len(tasks)-1].CreatedAt, true, cursor.LimitData)

    pageResponse := utils.CreatePaginationResponse(prevPage, nextPage)

    return tasks, &pageResponse, nil 
}

func(m *mysqlTaskRepository) FetchTaskByUID(
    ctx context.Context, 
    taskUID string,
) (domain.Task, error) {
    query := `SELECT * FROM task WHERE uid = ?`

    tasks, err := m.fetchTasks(ctx, query, taskUID)

    if err != nil {
        return domain.Task{}, err 
    }

    if len(tasks) == 0 {
        return domain.Task{}, domain.ErrNotFound
    }

    return tasks[0], nil 
}

func(m *mysqlTaskRepository) FetchTasksByUserUID(
    ctx context.Context, 
    cursor domain.Cursor,
    userUID string, 
    categoryID string,
) ([]domain.Task, *domain.PaginationResponse, error) {
    query := `SELECT t.*, 
                CASE WHEN uc.task_uid IS NULL THEN 0 
                ELSE 1 END AS user_completed
            FROM task t 
            LEFT JOIN 
                (SELECT ct.task_uid FROM completed_task ct
                INNER JOIN user u ON ct.user_uid = u.uid
                WHERE u.uid = ?) uc
            ON  t.uid = uc.task_uid
            WHERE t.level_category_id = ?

    `

    tasks, err := m.fetchTaskWithUserCompletion(ctx, query, userUID, categoryID)

    if err != nil {
        return nil, nil, domain.ErrServerError
    }

    if len(tasks) == 0 {
        return []domain.Task{}, nil, nil
    }

    prevPage := utils.CreateCursor(tasks[0].CreatedAt, false, cursor.LimitData)
    nextPage := utils.CreateCursor(tasks[len(tasks) - 1].CreatedAt, true, cursor.LimitData)

    pageResponse := utils.CreatePaginationResponse(prevPage, nextPage)

    return tasks, &pageResponse, nil 
}

func(m *mysqlTaskRepository) InsertTask(
    ctx context.Context,
    task *domain.Task,
) (error) {
    query := `INSERT INTO task 
        (uid, level, sign, level_category_id, created_at, updated_at) 
        VALUES 
        (?, ?, ?, ?, ?, ?)
    `

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("[REPOSITORY] failed to prepare statement in task_repository. ERR:%s\n", err.Error())
        return err 
    }

    currTime := utils.CurrentTime()

    _, err = stmt.ExecContext(ctx, task.UID, task.Level, task.Sign, task.LevelCategoryID, currTime, currTime)

    if err != nil {
        log.Printf("[REPOSITORY] failed to execute statement in task_repositroy. ERR:%s\n", err.Error())
        return err 
    }

    return nil 
}

func(m *mysqlTaskRepository) InsertCompletedTask(
    ctx context.Context,
    completedTask *domain.CompletedTask,
) (error) {

    query := `INSERT 
        INTO completed_task (user_uid, task_uid, completed_at)
        VALUES (?, ?, ?)
    `

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("[REPOSITORY] failed to prepare statement in task_repository. ERR:%s\n", err.Error())
        return err
    }

    currTime := utils.CurrentTime()

    _, err = stmt.ExecContext(ctx, completedTask.UserUID, completedTask.TaskUID, currTime)

    if err != nil {
        log.Printf("[REPOSITORY] failed to execute statement in task_repository. ERR:%s\n", err.Error())
        return err
    }

    return nil
}

func(m *mysqlTaskRepository) UpdateTask(
    ctx context.Context,
    task *domain.Task,
) error {
    query := `UPDATE task SET level = ?, sign = ?, level_category_id = ?, updated_at = ? WHERE uid = ?`

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("[REPOSITORY] failed to prepare statement in task_repository. ERR:%s\n", err.Error())
        return err 
    }

    currTime := utils.CurrentTime()

    rows, err := stmt.ExecContext(ctx, task.Level, task.Sign, task.LevelCategoryID, currTime, task.UID)

    if err != nil {
        log.Printf("[REPOSITORY] failed to execute statement in task_repositroy. ERR:%s\n", err.Error())
        return err 
    }

    affected, _ := rows.RowsAffected()

    if affected != 1 {
        if affected == 0 {
            return domain.ErrNotFound
        }

        return fmt.Errorf("weird behaviour. rows affected:%d\n", affected)
    }

    return nil 
}

func(m *mysqlTaskRepository) DeleteTask(
    ctx context.Context,
    uid string,
) error {
    query := `DELETE FROM task WHERE uid = ?`

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("[REPOSITORY] failed to prepare statement in task_repository. ERR:%s\n", err.Error())
        return err 
    }

    rows, err := stmt.ExecContext(ctx, uid)

    if err != nil {
        log.Printf("[REPOSITORY] failed to execute statement in task_repositroy. ERR:%s\n", err.Error())
        return err 
    }

    affected, _ := rows.RowsAffected()

    if affected != 1 {
        if affected == 0 {
            return domain.ErrNotFound
        }

        return fmt.Errorf("weird behaviour. rows affected:%d\n", affected)
    }

    return nil 
}