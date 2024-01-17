package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
    "strings"

	"github.com/devanfer02/litecartes/domain"
	"github.com/devanfer02/litecartes/internal/utils"
)

type mysqlQuestionRepository struct {
    Conn    *sql.DB
}

func NewMysqlQuestionRepository(conn *sql.DB) domain.QuestionRepository {
    return &mysqlQuestionRepository{conn}
}

func(m *mysqlQuestionRepository) fetch(
    ctx context.Context, 
    query string, 
    args ...interface{},
) ([]domain.Question, error) {
    rows, err := m.Conn.QueryContext(ctx, query, args...)

    if err != nil {
        log.Printf("[REPOSITORY] error querying to question table. ERR:%s\n", err.Error())
        return nil, domain.ErrServerError
    }

    defer func(){
        err := rows.Close()
        if err != nil {
            log.Printf("[REPOSITORY] failed to close rows. ERR:%s\n", err.Error())
        }
    }()
    
    questions := make([]domain.Question, 0)

    for rows.Next() {
        var question domain.Question
        var taskUID sql.NullString
        
        err := rows.Scan(
            &question.UID,
            &question.CategoryID,
            &taskUID,
            &question.Title,
            &question.Literacy,
            &question.Question,
            &question.Answer,
            &question.Options,
            &question.CreatedAt,
            &question.UpdatedAt,
        )

        if err != nil {
            log.Printf("[REPOSITORY] failed to scan row in question_repository. ERR:%s\n", err.Error())
            return nil, err 
        }

        if taskUID.Valid {
            question.TaskUID = taskUID.String
        } 

        question.OptionList = strings.Split(question.Options, "|")
        question.Options = ""

        questions = append(questions, question)
    }

    return questions, nil 
}

func (m *mysqlQuestionRepository) fetchPaged(ctx context.Context, cursor domain.Cursor) ([]domain.Question, error){
    var query string 
    var questions []domain.Question
    var err error 

    if cursor.CreatedAt == "" {
        query = `SELECT 
            question.uid, question_category.category_name, question.task_uid, question.title, 
            question.literacy, question.question, question.answer, 
            question.options, question.created_at, question.updated_at
            FROM question JOIN question_category ON question.category_id = question_category.uid ORDER BY question.created_at LIMIT ?`
        questions, err = m.fetch(ctx, query, cursor.LimitData)
    } else {
        query = fmt.Sprintf(
            `SELECT 
            question.uid, question_category.category_name, question.task_uid, question.title, 
            question.literacy, question.question, question.answer, 
            question.options, question.created_at, question.updated_at
            FROM question JOIN question_category ON question.category_id = question_category.uid 
            WHERE question.created_at %s ? ORDER BY question.created_at LIMIT ?`, 
            utils.GetPaginationOperator(cursor.PointNext),
        )
        questions, err = m.fetch(ctx, query, cursor.CreatedAt, cursor.LimitData)
    }

    return questions, err    
}

func(m *mysqlQuestionRepository) Fetch(
    ctx context.Context, 
    cursor domain.Cursor,
) ([]domain.Question, *domain.PaginationResponse, error) {
        
    questions, err := m.fetchPaged(ctx, cursor)

    if err != nil {
        return nil, nil, domain.ErrServerError
    }

    if len(questions) == 0 {
        return questions, nil, nil 
    }

    prevPage := utils.CreateCursor(questions[0].CreatedAt, false, cursor.LimitData)
    nextPage := utils.CreateCursor(questions[len(questions)-1].CreatedAt, true, cursor.LimitData)

    pageResponse := utils.CreatePaginationResponse(prevPage, nextPage)



    return questions, &pageResponse, nil 
}

func(m *mysqlQuestionRepository) FetchOneByArg(
    ctx context.Context, 
    param,
    arg string,
) (domain.Question, error) {
    query := fmt.Sprintf(`SELECT 
        question.uid, question_category.category_name, question.task_uid, 
        question.title, question.literacy, question.question, question.answer, 
        question.options, question.created_at, question.updated_at 
        FROM question JOIN question_category ON question.category_id = question_category.uid 
        WHERE question.%s = ? LIMIT 1`, param,
    )

    questions, err := m.fetch(ctx, query, arg)

    if err != nil {
        return domain.Question{}, domain.ErrServerError 
    }

    if len(questions) == 0 {
        return domain.Question{}, domain.ErrNotFound
    }

    return questions[0], nil 
}

func(m *mysqlQuestionRepository) FetchAllByTaskUID(
    ctx context.Context,
    taskUID  string, 
) ([]domain.Question, error) {
    query := `SELECT 
        *
        FROM question
        WHERE task_uid = ?
    `

    questions, err := m.fetch(ctx, query, taskUID)

    if err != nil {
        return nil, err 
    }

    return questions, nil 
}

func(m *mysqlQuestionRepository) InsertQuestion(
    ctx context.Context,
    question *domain.Question,
) (error) {
    query := `INSERT INTO question (uid, category_id, title, literacy, question, answer, options, created_at, updated_at) 
                VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("[REPOSITORY] failed to prepare statement. ERR:%s\n", err.Error())
        return err 
    }

    currTime := utils.CurrentTime()

    _, err = stmt.ExecContext(
        ctx, 
        question.UID, 
        question.CategoryID, 
        question.Title,
        question.Literacy, 
        question.Question, 
        question.Answer, 
        question.Options,
        currTime, 
        currTime,
    )

    if err != nil {
        log.Printf("[REPOSITORY] failed to execute statement. ERR:%s\n", err.Error())
        return err 
    }

    return nil 
}

func(m *mysqlQuestionRepository) UpdateQuestion(
    ctx context.Context, 
    question *domain.Question,
) (error) {
    query := "UPDATE question SET category_id = ?, task_uid = ?, title = ?, literacy = ?, question = ?, answer = ?, options = ?, updated_at = ? WHERE uid = ?"

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("[REPOSITORY] failed to prepare statement. ERR:%s\n", err.Error())
        return err 
    }

    currTime := utils.CurrentTime()
    rows, err := stmt.ExecContext(
        ctx, 
        question.CategoryID, 
        question.TaskUID, 
        question.Title,
        question.Literacy, 
        question.Question, 
        question.Answer,
        question.Options,
        currTime, 
        question.UID,
    )

    if err != nil {
        log.Printf("[REPOSITORY] failed to exec statement. ERR:%s\n", err.Error())
        return err 
    }

    affected, _ := rows.RowsAffected()

    if affected != 1 {
        if affected == 0 {
            return domain.ErrNotFound
        }
        return fmt.Errorf("weird behaviour. rows affected: %v\n", affected)
    }

    return nil 
}

func(m *mysqlQuestionRepository) DeleteQuestion(
    ctx context.Context,
    uid string,
) (error) {
    query := "DELETE FROM question WHERE uid = ?"

    stmt, err := m.Conn.PrepareContext(ctx, query)

    if err != nil {
        log.Printf("[REPOSITORY] failed to prepare statement. ERR:%s\n", err.Error())
        return err 
    }

    rows, err := stmt.ExecContext(ctx, uid)

    affected, _ := rows.RowsAffected()

    if affected != 1 {
        if affected == 0 {
            return domain.ErrNotFound
        }
        return fmt.Errorf("weird behaviour. rows affected: %v\n", affected)
    }

    return nil 
}