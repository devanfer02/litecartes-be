package usecase

import (
	"context"
	"time"

	"github.com/devanfer02/litecartes/domain"
	"github.com/devanfer02/litecartes/internal/utils"
	"github.com/go-playground/validator/v10"
)

type taskUsecase struct {
	taskRepo domain.TaskRepository
    queRepo domain.QuestionRepository
    ctxTimeout time.Duration
    taskV10 *validator.Validate
}

func NewTaskUsecase(tRepo domain.TaskRepository, qRepo domain.QuestionRepository, timeout time.Duration) domain.TaskUsecase {
    return &taskUsecase{taskRepo: tRepo, queRepo: qRepo, ctxTimeout: timeout, taskV10: validator.New()}
}

func(u *taskUsecase) FetchAll(
    ctx context.Context,
    req *domain.PaginationRequest,
) ([]domain.Task, *domain.PaginationResponse, error) {

    cursor, err := utils.GetCursor(req)    

    if err != nil {
        return nil, nil, err 
    }

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    tasks, pageResp, err := u.taskRepo.FetchAll(c, cursor)

    if err != nil {
        return nil, nil, err 
    }

    return tasks, pageResp, nil 
}

func(u *taskUsecase) FetchTasksByUserUID(
    ctx context.Context,
    req *domain.PaginationRequest,
    userUID string, 
    categoryID string,
) ([]domain.Task, *domain.PaginationResponse, error) {

    cursor, err := utils.GetCursor(req)

    if err != nil {
        return nil, nil, err
    }

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    tasks, pageResp, err := u.taskRepo.FetchTasksByUserUID(c, cursor, userUID, categoryID)

    if err != nil {
        return nil, nil, err 
    }    

    return tasks, pageResp, nil 
}

func(u *taskUsecase) FetchTaskWithQuestions(
    ctx context.Context,
    taskUID string,
) (domain.Task, error) {
    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    task, err := u.taskRepo.FetchTaskByUID(ctx, taskUID)

    if err != nil {
        return domain.Task{}, err 
    }

    questions, err := u.queRepo.FetchAllByTaskUID(c, taskUID)

    if err != nil {
        return domain.Task{}, err 
    }

    task.QuestionList = questions 

    return task, nil 
}

func(u *taskUsecase) InsertTask(
    ctx context.Context,
    task *domain.Task,
) error {
    
    if err := u.taskV10.Struct(task); err != nil {
        return domain.ValidationFailed(err.Error())
    }

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    task.UID = utils.CreateUUID() + "-TSK"
    err := u.taskRepo.InsertTask(c, task)

    if err != nil {
        return err 
    }

    return nil 
}

func(u *taskUsecase) InsertCompletedTask(
    ctx context.Context,
    completedTask *domain.CompletedTask,
) error {

    if err := u.taskV10.Struct(completedTask); err != nil {
        return domain.ValidationFailed(err.Error())
    }

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    err := u.taskRepo.InsertCompletedTask(c, completedTask)

    if err != nil {
        return err 
    }

    return nil 
}

func(u *taskUsecase) UpdateTask(
    ctx context.Context,
    task *domain.Task,
) error {

    if err := u.taskV10.Struct(task); err != nil {
        return domain.ValidationFailed(err.Error())
    }

    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    err := u.taskRepo.UpdateTask(c, task)

    if err != nil {
        return err 
    }

    return nil 
}

func(u *taskUsecase) DeleteTask(
    ctx context.Context,
    uid string,
) error {
    c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    err := u.taskRepo.DeleteTask(c, uid)

    if err != nil {
        return err
    }

    return nil 
}