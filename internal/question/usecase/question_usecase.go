package usecase

import (
	"context"
	"time"
    "strings"
    "fmt"

	"github.com/devanfer02/litecartes/domain"
	"github.com/devanfer02/litecartes/internal/utils"
	"github.com/go-playground/validator/v10"
)

type questionUsecase struct {
	queRepo   domain.QuestionRepository
    ctxTimeout time.Duration
    queV10 *validator.Validate
}

func NewQuestionUsecase(qRepo domain.QuestionRepository, timeout time.Duration) domain.QuestionUsecase {
    return &questionUsecase{queRepo: qRepo, ctxTimeout: timeout, queV10: validator.New()}
}

func(u *questionUsecase) Fetch(
    ctx context.Context, 
    req *domain.PaginationRequest,
) ([]domain.Question, *domain.PaginationResponse, error) {

    cursor, err := utils.GetCursor(req)
    
    if err != nil {
        return nil, nil, err 
    }

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    questions, pageResp, err := u.queRepo.Fetch(c, cursor)

    if err != nil {
        return nil, nil, err 
    }

    return questions, pageResp, nil 
}

func(u *questionUsecase) FetchByUID(
    ctx context.Context,
    uid string,
) (domain.Question, error) {
    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    question, err := u.queRepo.FetchOneByArg(c, "uid", uid)

    if err != nil {
        return domain.Question{}, err
    }

    return question, nil 
}

func(u *questionUsecase) FetchAllByTaskUID(
    ctx context.Context,
    taskUID  string, 
) ([]domain.Question, error) {
    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    questions, err := u.queRepo.FetchAllByTaskUID(c, taskUID)

    if err != nil {
        return nil, err 
    }

    return questions, nil 
}

func(u *questionUsecase) Insert(
    ctx context.Context, 
    question *domain.Question,
) error {

    if err := u.queV10.Struct(question); err != nil {
        return domain.ValidationFailed(err.Error())
    }

    splitted := strings.Split(question.Options, "|")

    length := len(splitted)

    if length != 4 {
        return domain.ValidationFailed(fmt.Sprintf("len of splitting: %d", length))
    }

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    question.UID = utils.CreateUUID() + "-QST"
    err := u.queRepo.InsertQuestion(c, question)

    if err != nil {
        return err 
    }

    return nil 
}

func(u *questionUsecase) Update(
    ctx context.Context,
    question *domain.Question, 
) error {

    if err := u.queV10.Struct(question); err != nil {
        return domain.ValidationFailed(err.Error())
    }

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    err := u.queRepo.UpdateQuestion(c, question)

    if err != nil {
        return err 
    }

    return nil 
}

func(u *questionUsecase) Delete(
    ctx context.Context,
    uid string, 
) error {
    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    err := u.queRepo.DeleteQuestion(c, uid)

    if err != nil {
        return err 
    }

    return nil 
}