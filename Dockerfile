FROM golang:alpine AS builder
WORKDIR /app
COPY . ./
RUN go mod download && go build -o todo-rest-api cmd/main.go

FROM alpine:3.17.0
WORKDIR /app
RUN addgroup -S user && adduser -S user -G user

COPY --from=builder /app/todo-rest-api /app/todo-rest-api
COPY .env  /app/
COPY configs  /app/configs

RUN chown -R user:user /app
USER user
CMD ["/app/todo-rest-api"]