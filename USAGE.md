# Usage

## Protoc

### Building

Replace service with the service we want to build in.

```bash
protoc --proto_path=. --go_out=. --micro_out=. \
		<service>/proto/consignment/consignment.proto
```

## Docker

### Building

```bash
docker build -t shippy-service-consignment .
```

### Running

```bash
docker run -p 50051:50051 --env-file ./<service-directory>/.env <service-name>
```

This can go through a differnet local port if we want it too. Changing `50051:50051` to `8080:50051` would open it on `8080` locally.
