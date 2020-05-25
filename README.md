
# Go TimescaleDB

Sample project to test TimescaleDB in Golang

## Run

1. start TimescaleDB
    ```shell script
    make start-timescaledb
    ```

2. start writer
    ```shell script
    make start-writer
    ```

3. start reader
    ```shell script
    make start-reader
    ```

---

## PostgreSQL

[Commands](postgres/commands.md)

[Backup & restore](postgres/backup-restore.md)

---

## TODOs

- [x] run timescaledb in container
- [x] writer application
- [x] reader application
- [x] logging
- [x] monitoring
- [x] tracing
- [x] container images
    - [x] writer
    - [x] reader
- [ ] run in kubernetes
    - [ ] timescaledb kube manifests - `WIP`
    - [x] writer kubernetes probes
    - [x] reader kubernetes probes
    - [x] writer kube manifests
    - [x] reader kube manifests

---

## Links

- https://docs.timescale.com/latest/getting-started/installation/docker/installation-docker
- https://docs.timescale.com/latest/getting-started/setup
- https://chartio.com/resources/tutorials/how-to-list-databases-and-tables-in-postgresql-using-psql/
- https://github.com/timescale/timescaledb-kubernetes
