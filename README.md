# Downloader Music

## Stack:
- Go 1.16
- Docker
## Requirements
- Create .env file in root directory and add following values:
```
TELEGRAM_TOKEN= <telegram token secret>
DB_PASSWORD= <qwerty>
```


### Optionally
```s
SUDO_CHAT_ID= <sudo chat id for send info>
```
### Run project
```
make run
```
### Build image Docker and run container
```
make build-image-multistage && make start-container
```

### Check code with linter
```
make lint
```

### Project Description

