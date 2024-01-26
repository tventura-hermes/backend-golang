## How to use

#### Build and run Docker

```bash
$ docker build --tag backend-golang:gin .

$ docker run --publish 8080:8080  backend-golang:gin
```

#### test 
 
```bash
$ curl http://localhost:8080/ping
```