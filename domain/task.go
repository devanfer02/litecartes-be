package domain

type Task struct {
    UID             string      `json:"uid"`
    QuestionList    []Question  `json:"questions_list"`
}