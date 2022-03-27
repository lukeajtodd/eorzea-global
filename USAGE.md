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
docker-compose build
```

### Running

```bash
docker-compose up
```

```bash
docker-compose up -d
```

```bash
docker-compose run <service>
```
