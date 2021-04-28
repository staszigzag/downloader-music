build:
	GOOS=linux CGO_ENABLED=0 go build -o ./.bin/app ./cmd/app/main.go

run:
	go run -race ./cmd/app -configPath=configs/config

run-docker: build
	docker-compose up

test:
	go test -v ./pkg/... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

lint:
	golangci-lint run

build-image-multistage:
	docker build -t downloader-music  -f Dockerfile.multistage .

build-image:
	docker build -t downloader-music .

start-container:
	docker run --env-file .env -p 80:80 downloader-music

create-init-migrate:
	migrate create -ext sql -dir ./schema -seq init_schema

run-postgres-docker:
	docker run --name=downloader-music-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -v  ${PWD}/postgres-data:/var/lib/postgresql/data -d --rm postgres

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down