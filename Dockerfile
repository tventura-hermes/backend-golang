FROM golang AS bin-stage

RUN mkdir -p /go/src/github.com/tventura-hermes/backend-golang-gin

WORKDIR /go/src/github.com/tventura-hermes/backend-golang-gin

RUN go mod init backend-golang-gin ;  go get -u github.com/gin-gonic/gin

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend-golang-gin

FROM bin-stage AS test-stage

RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS release-stage

WORKDIR /

COPY --from=bin-stage /backend-golang-gin /backend-golang-gin

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/backend-golang-gin"]
