# Compact gRPC server and client implementation (Go & Python)

## To build and run

Using docker compose:

```bash
docker-compose up --build
```

## Prerequisites

1. protoc for Windows, protobuf-compiler for Linux.
2. grpcio-tools, grpclib, protobuf (python packages).
3. https://github.com/nipunn1313/mypy-protobuf (`pip install mypy-protobuf`)

### Generate proto files

Run makefile:

```bash
make
```

### Export requirements.txt

Run uv export:

```bash
cd python
uv export --no-dev --output-file requirements.txt
```

### Notes & Todos

-   For newly proto generated python files, inside `_pb2_grpc.py` files, fix first import with absolute package path.
-   No need to manage proto generated python files, as they are generated in the project directory, even though it would be a good idea to publish and use it as third party package.
-   Would be nice to add multi-stage docker build.
-   Would be nice to use uv for python docker build.
-   Would be nice to add a proper di in python app.
-   Would be nice to correctly set mypy to ignore proto generated python files to allow the use of `--strict` flag.
-   Would be nice to use proper configuration techniques for both apps for different environments.

<i>Even though it is just a short demo of how adequately use go grpc server and python grpc client.</i>
