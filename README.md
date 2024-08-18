# go-microservices

## Init

Rebuild the go.work file

```sh
go work init
go work use -r .
```

## Add swagger docs to a service

Install swag & echo-swagger

```sh
go get -d github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/echo-swagger
```

Add docs build to the services makefile

```makefile
docs:
	swag fmt && swag init --dir ./,./model/,./handler/ --generalInfo main.go --requiredByDefault --outputTypes yaml,go
```

Import the generated docs in the main.go

```go
import (
  _ "github.com/leifarriens/go-microservices/services/product/docs"
)
```
