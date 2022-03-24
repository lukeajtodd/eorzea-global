# Usage

## Protoc

### Building

Replace service with the service we want to build in.

```bash
protoc --proto_path=. --go_out=. --micro_out=. \
		<service-name>/proto/consignment/consignment.proto
```

Alternatively we can run `./scripts/build-proto.sh aetheryte-service-consignment consignment`.

## Docker

### Building

```bash
docker build -t <service-name> .
```

Alternatively we can run `./scripts/build.sh aetheryte-service-consignment`.

### Running

```bash
docker run <service-name>
```

Alternatively we can run `./scripts/run.sh aetheryte-cli-consignment`.
