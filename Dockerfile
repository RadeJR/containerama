FROM golang:alpine as backend-build
ENV CGO_ENABLED=1 \
	GOOS=linux

RUN apk add --no-cache build-base
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/. .
RUN go build -o containerama cmd/main.go

FROM node:20-alpine as frontend-build
WORKDIR /app
COPY frontend/package*.json ./
RUN npm i
COPY frontend/ ./
RUN ls
RUN npm run build

FROM alpine:latest
ENV BIND_ADDR=:80
RUN apk add --no-cache curl && \
	curl -SL https://github.com/docker/compose/releases/download/v2.27.0/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose && \
	chmod +x /usr/local/bin/docker-compose
WORKDIR /app
COPY --from=backend-build /app/containerama ./
COPY --from=frontend-build /app/dist /app/public

EXPOSE 80
CMD ["/app/containerama"]
