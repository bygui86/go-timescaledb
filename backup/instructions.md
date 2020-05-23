# Backup & restore

## Backup single hypertable

### Schema
```
pg_dump -U postgres -s -d timeseries --table conditions -N _timescaledb_internal | grep -v _timescaledb_internal > timeseries_schema.sql
```

### Data
```
psql -U postgres -d timeseries -c "\COPY (SELECT * FROM conditions) TO timeseries_data.csv DELIMITER ',' CSV"
```

## Restore single hypertable

### Schema
```
psql -U postgres -d timeseries < timeseries_schema.sql

psql -U postgres -d timeseries -c "SELECT create_hypertable('conditions', 'time')"
```

### Data
```
psql -U postgres -d timeseries -c "\COPY conditions FROM timeseries_data.csv CSV"
```

## Links

- https://docs.timescale.com/latest/using-timescaledb/backup
