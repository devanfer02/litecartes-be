package domain

import (
    "time"
)

type Task struct {
    UID             string      `json:"uid"`
    CategoryID      string      `json:"category_id"`         
    QuestionList    []Question  `json:"questions_list"`   
}

type CompletedTask struct {
    UserUID         string      `json:"user_uid"`
    TaskUID         string      `json:"task_uid"`
    CompletedDate   time.Time   `json:"completed_date"`
}

/*
Task Left Join On CompletedTask Table WHERE CompletedTask.UserUID = authorized user
*/