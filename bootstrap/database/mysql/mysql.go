package mysql 

import (
	"fmt"
	"log"
	"os"
	"database/sql"

	"github.com/devanfer02/litecartes/bootstrap/env"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlConn() *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.ProcEnv.DBUser, 
		env.ProcEnv.DBPassword,
		env.ProcEnv.DBHost, 
		env.ProcEnv.DBPort,
		env.ProcEnv.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("[MYSQL] Failed to open database. ERR: %s\n", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("[MYSQL] Could not ping database. ERR: %s\n", err.Error())
	}

    // gotta be in correct order to form foreign key constraint
    migrate(
        db,
        "bootstrap/database/mysql/migrations/create_question_category_table.sql", 
        "bootstrap/database/mysql/migrations/create_question_table.sql",
        "bootstrap/database/mysql/migrations/create_subscription_table.sql",
        "bootstrap/database/mysql/migrations/create_school_table.sql",
        "bootstrap/database/mysql/migrations/create_user_table.sql",
        "bootstrap/database/mysql/migrations/create_level_category_table.sql",
        "bootstrap/database/mysql/migrations/create_task_table.sql",
        "bootstrap/database/mysql/migrations/create_completed_task_table.sql",
    )

	return db 
}

func migrate(db *sql.DB, migrationspath ...string) {
	for _, filename := range migrationspath {
		filecontent, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("[MYSQL] Failed to read migration file {%s}. ERR: %s\n", filename, err.Error())
		}

		_, err = db.Exec(string(filecontent))
		if err != nil {
			log.Fatalf("[MYSQL] Failed to execute migration file {%s}. ERR: %s\n", filename, err.Error())
		}

		log.Printf("[MYSQL] Migration file {%s} success\n", filename)
	}
}

func GenerateSeeders(db *sql.DB) {
    seeders(
        db, 
        "bootstrap/database/mysql/seeders/create_question_category_seeders.sql",
        "bootstrap/database/mysql/seeders/create_level_category_seeders.sql",
        "bootstrap/database/mysql/seeders/create_subscription_seeders.sql",
    )
    log.Printf("Seeders Generated!\n")
}

func seeders(db *sql.DB, seederspath ...string) {
    for _, filename := range seederspath {
		filecontent, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("[MYSQL] Failed to read seeders file {%s}. ERR: %s\n", filename, err.Error())
		}

		_, err = db.Exec(string(filecontent))
		if err != nil {
			log.Fatalf("[MYSQL] Failed to execute seeders file {%s}. ERR: %s\n", filename, err.Error())
		}

		log.Printf("[MYSQL] Seeders file {%s} success\n", filename)
	}
}