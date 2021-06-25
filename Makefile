build:
	mkdir -p dolt-data
	go get -u \
	  github.com/go-sql-driver/mysql \
	  github.com/getsentry/sentry-go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o kanban
	docker-compose build dolt kanban

.PHONY: build
