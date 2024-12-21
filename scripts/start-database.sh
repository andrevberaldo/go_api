#!/bin/bash

docker compose --profile default up -d
go run ./cmd/main.go