FROM golang:1.19-alpine as builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY *.go ./

ENV GOPROXY direct
RUN go build -o /rps-backend

FROM alpine:latest as runner

WORKDIR /app
COPY --from=builder /rps-backend .

ENV GIN_MODE release
CMD [ "/app/rps-backend" ]