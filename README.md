# To build and run

```bash
docker-compose up --build
```

## Prerequisites

1. protoc for Windows, protobuf-compiler for Linux.
2. grpcio-tools, grpclib, protobuf (python packages).

### Run

```bash
go run cmd/sso/main.go
```

### Generate proto files

Run makefile:

```bash
make
```
