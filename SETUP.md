# Setup

1. [Install Go](https://go.dev/doc/install) (1.17 or greater). `export GO111MODULE=on` should be set.
2. [Install protoc](https://grpc.io/docs/protoc-installation/)
3. [Install gRPC](https://grpc.io/docs/languages/go/quickstart/)
4. Install the following libs:

```bash
go install google.golang.org/grpc
go install github.com/golang/protobuf/protoc-gen-go@v1.3
go install go-micro.dev/v4
go install go-micro.dev/v4/cmd/protoc-gen-micro@v4
```
