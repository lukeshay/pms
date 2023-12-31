FROM node:18.16.0 as node_builder

ENV CI=true

WORKDIR /app

RUN apt update && apt-get install -y openjdk-17-jre openjdk-17-jdk
RUN npm i -g npm

ENV PATH="/root/.bun/bin:${PATH}"

COPY solid/package.json ./solid/
COPY packages/api/package.json ./packages/api/
COPY package.json package-lock.json ./

RUN npm ci

COPY docs ./docs/
COPY solid ./solid/
COPY packages/api ./packages/api/

RUN npm run build

FROM golang:1.20.5-alpine3.18 as go_installer

ENV CGO_ENABLED=1

WORKDIR /app

RUN apk add --no-cache git gcc musl-dev coreutils

COPY go.mod go.sum ./

RUN go mod download

COPY docs ./docs/
COPY pkg ./pkg/
COPY cmd ./cmd/

FROM go_installer as app_builder

ENV CGO_ENABLED=1

WORKDIR /app

RUN go build -buildvcs=false -ldflags "-s -w -extldflags '-static'" -tags osusergo,netgo -o /app/bin/ ./cmd/app

FROM go_installer as migrate_builder

ENV CGO_ENABLED=1

WORKDIR /app

RUN go build -buildvcs=false -ldflags "-s -w -extldflags '-static'" -tags osusergo,netgo -o /app/bin/ ./cmd/migrate

FROM alpine:3.18

RUN apk add --no-cache bash fuse3 sqlite ca-certificates curl

ENV CGO_ENABLED=1

WORKDIR /app

ADD ./litefs.yml /etc/litefs.yml

COPY db/migrations /app/migrations
COPY --from=flyio/litefs:0.4 /usr/local/bin/litefs /usr/local/bin/litefs
COPY --from=node_builder /app/frontend/dist /app/frontend-dist
COPY --from=app_builder /app/bin/app /usr/local/bin/app
COPY --from=migrate_builder /app/bin/migrate /usr/local/bin/migrate

EXPOSE 3000

ENTRYPOINT ["litefs", "mount"]
