FROM golang:1.24-alpine as builder
WORKDIR /app
COPY ./filter-service .
RUN go build -o /filter-service ./cmd

FROM alpine:latest
COPY ./filter-service/.env .env
COPY --from=builder /filter-service /filter-service
CMD [ "/filter-service" ]