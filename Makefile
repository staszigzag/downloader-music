build:
	GOOS=linux CGO_ENABLED=0 go build -o ./.bin/app ./cmd/app/main.go

run:
	go run -race ./cmd/app -configPath=configs/config
run-docker: build
	docker-compose up

lint:
	golangci-lint run

build-image-multistage:
	docker build -t downloader-music  -f Dockerfile.multistage .

build-image:
	docker build -t downloader-music .

start-container:
	docker run --env-file .env -p 80:80 downloader-music