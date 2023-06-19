FROM node:18.16.0-alpine3.18 as node_builder

WORKDIR /app/frontend

RUN apk add --no-cache openjdk17

COPY frontend/package.json frontend/package-lock.json ./

RUN npm ci

COPY docs ../docs
COPY frontend ./

RUN npm run build

FROM golang:1.20.5-alpine3.18 as go_builder

ENV CGO_ENABLED=1

WORKDIR /app

RUN apk add --no-cache git gcc musl-dev coreutils

COPY go.mod go.sum ./

RUN go mod download

COPY docs ./docs/
COPY pkg ./pkg/
COPY main.go ./

RUN go build -o /app/bin/ .

FROM alpine:3.18

ENV CGO_ENABLED=1

WORKDIR /app

COPY --from=node_builder /app/frontend/dist ./frontend-dist
COPY --from=go_builder /app/bin/ .

EXPOSE 3000

ENTRYPOINT [ "./pms" ]
