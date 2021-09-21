# Build backend
FROM golang:1.17.1-alpine AS backend

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
RUN npm install

COPY web .
RUN ./node_modules/.bin/ng build --prod --source-map


# Deploy stage
FROM alpine

WORKDIR /app
COPY --from=backend /build/claerence ./claerence
COPY --from=frontend /build/public/ ./public/

EXPOSE 1401

ENTRYPOINT ["./claerence"]