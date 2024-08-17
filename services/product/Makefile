.PHONY: build docs

build:
	go build -o ./tmp/main .

docs:
	swag fmt && swag init --dir ./,./model/,./handler/ --generalInfo main.go --requiredByDefault --outputTypes yaml,go