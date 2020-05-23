package database

const (
	createDbQuery = `CREATE DATABASE %s;`

	createTableQuery = `CREATE TABLE IF NOT EXISTS conditions (
	time        TIMESTAMPTZ       NOT NULL,
	location    TEXT              NOT NULL,
	temperature DOUBLE PRECISION  NULL
);`

	createHypertableQuery = `SELECT create_hypertable('conditions', 'time');`

	insertConditionQuery = "INSERT INTO conditions(time, location, temperature) VALUES ($1, $2, $3);"
)
