# Usage

## Protoc

### Building

Replace service with the service we want to build in.

```bash
protoc --proto_path=. --go_out=. --micro_out=. \
		<service-name>/proto/consignment/consignment.proto
```

## Docker

### Building

```bash
docker build -t <service-name> .
```

### Running (Port & env required)

```bash
docker run -p 50051:50051 --env-file ./<service-directory>/.env <service-name>
```

This can go through a different local port if we want it too. Changing `50051:50051` to `8080:50051` would open it on `8080` locally.
The .env file port would need updating alongside the trailing port.

### Running (CLI)

```bash
docker run <service-name>
```
