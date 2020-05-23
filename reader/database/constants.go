package database

const (
	queryFormat                           = "%s %s %s %s"
	selectConditionsQuery                 = "SELECT * FROM conditions"
	locationDateRangeWhereConditionsQuery = "WHERE location = $1 AND time >= $2 AND time <= $3"
	dateRangeOnlyWhereConditionsQuery     = "WHERE time >= $1 AND time <= $2"
	orderbyConditionsQuery                = "ORDER BY time DESC"
	limitConditionsQuery                  = "LIMIT %d;"
)
