package postgres

import (
	"database/sql"
	"time"

	"fmt"

	appConfig "github.com/AntonyIS/notelify-logging-service/config"
	"github.com/AntonyIS/notelify-logging-service/internal/core/domain"
	_ "github.com/lib/pq"
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

	connectionAttemps := 1
	db, err := dbConnectionAttempts(dsn, connectionAttemps)

	if err != nil {
		return nil, err
	}

	err = dbPingAttempts(db, connectionAttemps)

	if err != nil {
		return nil, err
	}

	queryString := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			log_id VARCHAR(255) PRIMARY KEY UNIQUE,
			log_level VARCHAR(10) NOT NULL,
			message VARCHAR(255) NOT NULL,
			service VARCHAR(10) NOT NULL
	)
	`, tablename)

	_, err = db.Exec(queryString)
	if err != nil {
		return nil, err
	}

	return &postgresDBClient{db: db, tablename: tablename}, nil
}

func (psql *postgresDBClient) CreateLog(logEntry domain.LogMessage) error {
	query := fmt.Sprintf(`INSERT INTO %s (log_id,log_level,message,service) VALUES ($1,$2,$3,$4)`, psql.tablename)
	_, err := psql.db.Exec(
		query,
		logEntry.LogID,
		logEntry.LogLevel,
		logEntry.Message,
		logEntry.Service,
	)

	if err != nil {
		return err
	}
	return nil
}

func (psql *postgresDBClient) GetLogs() (*[]domain.LogMessage, error) {
	query := fmt.Sprintf(`SELECT log_id, log_level,message, service FROM %s `, psql.tablename)
	rows, err := psql.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	logs := []domain.LogMessage{}
	for rows.Next() {
		var logEntry domain.LogMessage
		err := rows.Scan(
			&logEntry.LogID,
			&logEntry.LogLevel,
			&logEntry.Message,
			&logEntry.Service,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, logEntry)
	}

	return &logs, nil
}

func (psql *postgresDBClient) GetServiceLogs(service string) (*[]domain.LogMessage, error) {
	query := fmt.Sprintf(`
		SELECT 
		log_id, 
		log_level, 
		service, 
		message
		FROM %s WHERE service = $1`, psql.tablename)

	rows, err := psql.db.Query(query, service)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	logs := []domain.LogMessage{}
	for rows.Next() {
		var logEntry domain.LogMessage
		err := rows.Scan(
			&logEntry.LogID,
			&logEntry.LogLevel,
			&logEntry.Message,
			&logEntry.Service,
		)

		if err != nil {
			return nil, err
		}
		logs = append(logs, logEntry)
	}
	return &logs, nil

}

func (psql *postgresDBClient) GetServiceLogsByLogLevel(service, log_level string) (*[]domain.LogMessage, error) {
	query := fmt.Sprintf(`
		SELECT 
		log_id, 
		log_level, 
		service, 
		message
		FROM %s WHERE service = $1 AND log_level =$2`, psql.tablename)

	rows, err := psql.db.Query(query, service, log_level)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	logs := []domain.LogMessage{}
	for rows.Next() {
		var logEntry domain.LogMessage
		err := rows.Scan(
			&logEntry.LogID,
			&logEntry.LogLevel,
			&logEntry.Message,
			&logEntry.Service,
		)

		if err != nil {
			return nil, err
		}
		logs = append(logs, logEntry)
	}
	return &logs, nil
}

func dbConnectionAttempts(dsn string, connectionAttemps int) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		if connectionAttemps <= 3 {
			fmt.Println("Sleeping for 5 seconds on count ", connectionAttemps)
			time.Sleep(5 * time.Second)
			connectionAttemps += 1
			dbConnectionAttempts(dsn, connectionAttemps)
		} else {
			return nil, err
		}
	}

	return db, nil
}

func dbPingAttempts(db *sql.DB, connectionAttemps int) error {
	err := db.Ping()
	if err != nil {
		if connectionAttemps <= 3 {
			fmt.Println("DB Ping attept :", connectionAttemps)
			time.Sleep(5 * time.Second)
			connectionAttemps += 1
			dbPingAttempts(db, connectionAttemps)
		} else {
			return err
		}
	}

	return nil
}
