FROM golang:alpine as build
ENV CGO_ENABLED=1 \
	GOOS=linux

RUN apk add --no-cache build-base
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go run github.com/a-h/templ/cmd/templ@latest generate && \
	go build -o containerama cmd/containerama/main.go

FROM alpine:latest
ENV BIND_ADDR=:80
RUN apk add --no-cache curl && \
	curl -SL https://github.com/docker/compose/releases/download/v2.27.0/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose && \
	chmod +x /usr/local/bin/docker-compose
WORKDIR /app
COPY --from=build /app/containerama /app/assets ./

EXPOSE 80
CMD ["/app/containerama"]
