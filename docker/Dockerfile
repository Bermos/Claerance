# Build backend
FROM golang:1.17.1-alpine AS backend

RUN apk add g++

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd cmd
COPY internal internal
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o claerence cmd/main.go


# Build frontend
FROM node:alpine AS frontend

WORKDIR /build

COPY web/package.json .
COPY web/package-lock.json .
RUN npm install

COPY web .
RUN ./node_modules/.bin/ng build --configuration production --source-map


# Deploy stage
FROM alpine

WORKDIR /app
COPY conf/ ./conf/
COPY --from=backend /build/claerence ./claerence
COPY --from=frontend /build/public/ ./web/public/

EXPOSE 1401
VOLUME /app/conf/

ENTRYPOINT ["./claerence"]
