## How to use

#### Start services

```bash
$ docker-compose up -d --build
```

#### Postgres Container Access

```bash
$ docker exec -it postgres-db /bin/bash

$ psql -U postgres -d backend_golang_gin
```

#### Stop services

```bash
$ docker-compose down
```