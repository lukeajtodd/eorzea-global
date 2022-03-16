# Usage

## Protoc

### Building

Replace service with the service we want to build in.

```bash
protoc --proto_path=. --go_out=. --micro_out=. \
		<service-name>/proto/consignment/consignment.proto
```

Alternatively we can run `./scripts/build-proto.sh aetheryte-service-consignment`.

## Docker

### Building

```bash
docker build -t <service-name> .
```

Alternatively we can run `./scripts/build.sh aetheryte-service-consignment`.

### Running (Full service)

```bash
docker run -p 50051:50051 --env-file ./<service-directory>/.env <service-name>
```

This can go through a different local port if we want it too. Changing `50051:50051` to `8080:50051` would open it on `8080` locally.
The .env file port would need updating alongside the trailing port.

Alternatively we can run `./scripts/run-full.sh aetheryte-service-consignment 50051`.

### Running (CLI)

```bash
docker run <service-name>
```

Alternatively we can run `./scripts/run.sh aetheryte-cli-consignment`.
