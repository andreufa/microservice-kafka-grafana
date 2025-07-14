FROM golang:1.24-alpine as builder
WORKDIR /app
COPY ./processor-service .
RUN go build -o /processor-service ./cmd

FROM alpine:latest
COPY --from=builder /processor-service /processor-service
CMD [ "/processor-service" ]