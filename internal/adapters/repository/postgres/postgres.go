package postgres

import (
	"database/sql"

	"fmt"
	"log"

	appConfig "github.com/AntonyIS/notelify-logging-svc/config"
	"github.com/AntonyIS/notelify-logging-svc/internal/core/domain"
)

type postgresDBClient struct {
	db        *sql.DB
	tablename string
}

func NewPostgresClient(appConfig appConfig.Config) (*postgresDBClient, error) {
	dbname := appConfig.POSTGRES_DB
	tablename := appConfig.LOGGING_TABLE
	user := appConfig.POSTGRES_USER
	password := appConfig.POSTGRES_PASSWORD
	port := appConfig.POSTGRES_PORT
	host := appConfig.POSTGRES_HOST

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	queryString := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			log_id VARCHAR(255) PRIMARY KEY UNIQUE,
			service VARCHAR(255) NOT NULL
			message VARCHAR(255) NOT NULL
	)
	`, tablename)

	_, err = db.Exec(queryString)
	if err != nil {
		return nil, err
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err

	}

	return &postgresDBClient{db: db, tablename: tablename}, nil
}

func (psql *postgresDBClient) Log(message domain.LogMessage) {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			log_id,
			message,
			service,
		)
		VALUES ($1,$2)`,
		psql.tablename)
	_, err := psql.db.Exec(
		query,
		message.Log_id,
		message.Message,
	)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Println(err.Error())
		}
	}
}
