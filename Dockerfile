FROM golang:1.24

RUN apt-get update

ARG APP_NAME=billing-engine-api
RUN mkdir /$APP_NAME
COPY . /$APP_NAME
WORKDIR /$APP_NAME

RUN go run database/migrate/migrate.go -t up
RUN go mod download
RUN go build -o billing_engine_api .

CMD ["./billing_engine_api"]
