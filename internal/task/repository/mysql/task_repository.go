package repository

import (
	"context"
	"database/sql"

	"github.com/devanfer02/litecartes/domain"
)

type mysqlTaskRepository struct {
	Conn *sql.DB
}

func NewMysqlTaskRepository(conn *sql.DB) *mysqlTaskRepository {
    return &mysqlTaskRepository{conn}
}

func(m *mysqlTaskRepository) Fetch(
    ctx context.Context, 
    cursor domain.Cursor,
)

