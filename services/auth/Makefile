.PHONY: build docs

build:
	go build -o ./tmp/main ./cmd

docs:
	swag fmt && swag init --dir ./,./handler/ --generalInfo cmd/main.go --requiredByDefault --outputTypes yaml,go