# connecterr

*connecterr* is Go static analyzer for development with [connect-go](https://github.com/connectrpc/connect-go).

This analyser warns when an error is returned without being wrapped in connect.NewError.

```go
// NG (warn by analyzer)
return nil, err

// OK (no warning)
return nil, connect.NewError(connect.CodeInternal, err)
```

## Install / How To Use

It can be installed with the `go install` command.

```
go install github.com/ry023/connecterr/cmd/connecterr@v0.1.0
```

Diagnostics can be performed by specifying this analyser binary as `-vettool` when `go vet`.

```
go vet -vettool=$(which connecterr) ./...
```
