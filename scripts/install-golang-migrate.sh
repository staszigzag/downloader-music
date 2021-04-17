#!/bin/bash
sudo apt update && sudo apt upgrade
go install -tags ' postgres ' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate -version
exit 0
