package domain

import (
	"context"
	"time"
)

type Task struct {
    UID                     string      `json:"uid"`
    Level                   int64       `json:"level" validate:"required"`
    Sign                    string      `json:"sign" validate:"required"`
    LevelCategoryID         int64       `json:"level_category_id" validate:"required"`         
    Completed               int64       `json:"completed"`
    QuestionList            []Question  `json:"questions_list,omitempty"`   
    CreatedAt               time.Time   `json:"created_at"`
    UpdatedAt               time.Time   `json:"updated_at"`
}

type CompletedTask struct {
    ID              int64       `json:"id"`
    UserUID         string      `json:"user_uid" validate:"required"`
    TaskUID         string      `json:"task_uid" valdate:"required"`
    CompletedAt     time.Time   `json:"completed_at"`
}

type TaskUsecase interface {
    FetchAll(ctx context.Context, req *PaginationRequest,) ([]Task, *PaginationResponse, error) 
    FetchTasksByUserUID(ctx context.Context, pageReq *PaginationRequest, userUID string, categoryID string) ([]Task, *PaginationResponse, error)
    FetchTaskWithQuestions(ctx context.Context,taskUID string) (Task, error)
    InsertTask(ctx context.Context, task *Task) error
    InsertCompletedTask(ctx context.Context, completedTask *CompletedTask) error
    UpdateTask(ctx context.Context, task *Task) error
    DeleteTask(ctx context.Context, uid string) error
}

type TaskRepository interface {
    FetchAll(ctx context.Context, cursor Cursor) ([]Task, *PaginationResponse, error) 
    FetchTasksByUserUID(ctx context.Context, cursor Cursor, userUID string, categoryID string) ([]Task, *PaginationResponse, error)
    FetchTaskByUID(ctx context.Context, taskUID string) (Task, error)
    InsertTask(ctx context.Context, task *Task) error
    InsertCompletedTask(ctx context.Context, completedTask *CompletedTask) error
    UpdateTask(ctx context.Context, task *Task) error
    DeleteTask(ctx context.Context, uid string) error 
}