package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "host=localhost port=5433 user=auditor password=Password! dbname=wlm_audit sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Database connection failed: %v", err))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("Database is not reachable: %v", err))
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS screening_audit (
		id SERIAL PRIMARY KEY,
		request_id VARCHAR(100),
		system_name VARCHAR(100),
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		is_hit BOOLEAN,
		rule_name VARCHAR(100),
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(fmt.Sprintf("Table could not be created: %v", err))
	}

	fmt.Println("✅ Audit-Database ready")
	return db
}

func LogScreening(db *sql.DB, request ScreeningRequest, results []MatchResult) {
	insertQuery := `
	INSERT INTO screening_audit (request_id, system_name, first_name, last_name, is_hit, rule_name)
	VALUES ($1, $2, $3, $4, $5, $6)`

	for _, res := range results {
		_, err := db.Exec(insertQuery, request.RequestId, request.System, request.FirstName, request.LastName, res.IsHit, res.RuleName)
		if err != nil {
			fmt.Printf("⚠️ Error writing the Audit-Logs: %v\n", err)
		}
	}
}
