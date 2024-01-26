FROM golang:1.21 AS bin-stage

SHELL ["/bin/bash", "-c"]

RUN mkdir -p /go/src/github.com/tventura-hermes/backend-golang-gin

WORKDIR /go/src/github.com/tventura-hermes/backend-golang-gin

COPY . .

RUN go mod init backend-golang-gin ;  go get -u github.com/gin-gonic/gin

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend-golang-gin

FROM gcr.io/distroless/base-debian11 AS release-stage

SHELL ["/bin/bash", "-c"]

WORKDIR /

COPY --from=bin-stage /backend-golang-gin /backend-golang-gin

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/backend-golang-gin"]
