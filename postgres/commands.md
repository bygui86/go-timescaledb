
# PostgreSQL - Commands

## psql
```
psql -U {username}

# or

psql -U {username} -d {database}
```

## List databases
```
\l
```

## Use database
```
\c {database}
```

## List tables
```
\dt
```

## Create database
```sql
CREATE DATABASE timeseries;
```

## [OPTIONAL] Extend the database with TimescaleDB
```sql
CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;
```

## Create table
```sql
CREATE TABLE conditions (
    time        TIMESTAMPTZ       NOT NULL,
    location    TEXT              NOT NULL,
    temperature DOUBLE PRECISION  NULL
);
```

## Create hypertable
```sql
SELECT create_hypertable('conditions', 'time');
```

## Alter table
```sql
ALTER TABLE conditions
  ADD COLUMN humidity DOUBLE PRECISION NULL;
```

## Drop table
```sql
DROP TABLE conditions;
```

## Create index
```sql
-- For indexing columns with discrete (limited-cardinality) values (e.g., where you are most likely to use an "equals" or "not equals" comparator) we suggest using an index like this (using our hypertable conditions for the example):
CREATE INDEX ON conditions (location, time DESC);

-- For all other types of columns, i.e., columns with continuous values (e.g., where you are most likely to use a "less than" or "greater than" comparator) the index should be in the form:
CREATE INDEX ON conditions (time DESC, temperature);
```

## Select
```sql
SELECT * FROM conditions ORDER BY time DESC LIMIT %d;

SELECT * FROM conditions WHERE location = $1 AND time >= $2 AND time <= $3 ORDER BY time DESC LIMIT %d;
```

## Insert
```sql
INSERT INTO conditions(time, location, temperature) VALUES ($1, $2, $3);
```
