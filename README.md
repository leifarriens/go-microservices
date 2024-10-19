# go-microservices

## Init

Rebuild the go.work file

```sh
go work init
go work use -r .
```

## Add swagger docs to a service

Import the generated docs in the main.go

```go
import (
  _ "github.com/leifarriens/go-microservices/services/{service}/docs"
)
```
