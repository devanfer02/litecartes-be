package domain

import (
    "context"
    "time"
)

type Question struct {
    UID         string      `json:"uid"`
    CategoryID  string      `json:"category_id" validate:"required"`
    TaskUID     string      `json:"task_uid"`
    Title       string      `json:"title" validate:"required"`
    Literacy    string      `json:"literacy" validate:"required"`
    Question    string      `json:"question" validate:"required"`
    Answer      string      `json:"answer" validate:"required"`
    Options     string      `json:"options,omitempty"`
    OptionList  []string    `json:"option_list"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}

type QuestionRepository interface {
    Fetch(ctx context.Context, cursor Cursor) ([]Question, *PaginationResponse, error)
    FetchOneByArg(ctx context.Context, param, arg string) (Question, error)
    FetchAllByTaskUID(ctx context.Context,taskUID  string) ([]Question, error) 
    InsertQuestion(ctx context.Context, question *Question) error
    UpdateQuestion(ctx context.Context, question *Question) error
    DeleteQuestion(ctx context.Context, uid string) error
}

type QuestionUsecase interface {
    Fetch(ctx context.Context, req *PaginationRequest) ([]Question, *PaginationResponse, error)
    FetchByUID(ctx context.Context, uid string) (Question, error)
    FetchAllByTaskUID(ctx context.Context,taskUID  string) ([]Question, error) 
    Insert(ctx context.Context, question *Question) error 
    Update(ctx context.Context, question *Question) error
    Delete(ctx context.Context, uid string) error
}