## How to use

#### Build and run Docker

```bash
$ docker-compose build 

$ docker-compose up
```

#### Postgres Container Access

```bash
$ docker exec -it postgres-db /bin/bash

$ psql -U postgres -d backend_golang_gin
```