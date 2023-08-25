# base Go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o authenticationApp ./cmd/api

RUN chmod +x /app/authenticationApp

# build a tiny docker
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/authenticationApp /app

CMD [ "/app/authenticationApp" ]